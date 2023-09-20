package jetstreamclient

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
	"go.uber.org/zap"
)

type Config struct {
	NatsUrl    string
	StreamName string
	Logger     *zap.SugaredLogger
	Nkey       string
	NkeySeed   string
}

type Client struct {
	natsConn            *nats.Conn
	streamName          string
	personSubject       string
	organizationSubject string
	logger              *zap.SugaredLogger
}

func New(config *Config) (*Client, error) {
	nc, err := natsConnect(config)
	if err != nil {
		return nil, err
	}

	cl := &Client{
		natsConn:            nc,
		streamName:          config.StreamName,
		personSubject:       config.StreamName + ".person",
		organizationSubject: config.StreamName + ".organization",
		logger:              config.Logger,
	}

	if err := cl.initStream(); err != nil {
		return nil, err
	}

	return cl, nil
}

func (c *Client) initStream() error {
	js, err := c.natsConn.JetStream()
	if err != nil {
		return err
	}

	stream, _ := js.StreamInfo(c.streamName)

	if stream == nil {
		_, err := js.AddStream(&nats.StreamConfig{
			Name: c.streamName,
			Subjects: []string{
				c.personSubject,
				c.organizationSubject,
			},
			Storage: nats.FileStorage,
		})
		if err != nil {
			return err
		}
	}

	consumerConfigs := []nats.ConsumerConfig{
		{
			AckPolicy:      nats.AckExplicitPolicy,
			Durable:        "inboxPerson",
			DeliverSubject: c.streamName + ".inboxPersonDeliverSubject",
			AckWait:        time.Minute,
			MaxAckPending:  1,
			FilterSubject:  c.personSubject,
			DeliverPolicy:  nats.DeliverAllPolicy,
		},
		{
			AckPolicy:      nats.AckExplicitPolicy,
			Durable:        "inboxOrganization",
			DeliverSubject: c.streamName + ".inboxOrganizationDeliverSubject",
			AckWait:        time.Minute,
			MaxAckPending:  1,
			FilterSubject:  c.organizationSubject,
			DeliverPolicy:  nats.DeliverAllPolicy,
		},
	}

	for _, consumerConfig := range consumerConfigs {
		consumerInfo, _ := js.ConsumerInfo(c.streamName, consumerConfig.Durable)

		if consumerInfo == nil {
			_, err := js.AddConsumer(c.streamName, &consumerConfig)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Client) PersonSubject() string {
	return c.personSubject
}

func (c *Client) OrganizationSubject() string {
	return c.organizationSubject
}

func (c *Client) SubscribePerson(cb func(msg *nats.Msg)) (*nats.Subscription, error) {
	js, err := c.natsConn.JetStream()
	if err != nil {
		return nil, err
	}
	return js.Subscribe(
		c.personSubject,
		cb,
		nats.AckExplicit(),
		nats.MaxAckPending(1),
		nats.ManualAck(),
		nats.Bind(c.streamName, "inboxPerson"),
	)
}

func (c *Client) SubscribeOrganization(cb func(msg *nats.Msg)) (*nats.Subscription, error) {
	js, err := c.natsConn.JetStream()
	if err != nil {
		return nil, err
	}
	return js.Subscribe(
		c.organizationSubject,
		cb,
		nats.AckExplicit(),
		nats.MaxAckPending(1),
		nats.ManualAck(),
		nats.Bind(c.streamName, "inboxOrganization"),
	)
}

func natsConnect(config *Config) (*nats.Conn, error) {
	options := nats.Options{
		Url:                  config.NatsUrl,
		MaxReconnect:         100, // try reconnect n times, and then give up
		RetryOnFailedConnect: true,
		ReconnectWait:        10 * time.Second,
		Timeout:              10 * time.Second, // connection timeout
		AllowReconnect:       true,
	}

	if config.Nkey != "" && config.NkeySeed != "" {
		user, err := nkeys.FromSeed([]byte(config.NkeySeed))
		if err != nil {
			return nil, err
		}
		options.Nkey = config.Nkey
		options.SignatureCB = func(nonce []byte) ([]byte, error) {
			return user.Sign(nonce)
		}
	}

	options.DisconnectedErrCB = func(c *nats.Conn, err error) {
		config.Logger.Errorf("Client connection to NATS closed, and was unable to reconnect (num reconnections: %d): %s", c.Reconnects, err)
	}
	options.ReconnectedCB = func(c *nats.Conn) {
		config.Logger.Infof("Client connection to NATS restored")
	}
	options.ClosedCB = func(c *nats.Conn) {
		config.Logger.Infof("Client connection to NATS closed")
	}

	return options.Connect()
}
