package cli

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
	"github.com/ugent-library/people-service/gismo"
	"github.com/ugent-library/people-service/subscribers"
)

var natsStreamConfig = nats.StreamConfig{
	Name: "GISMO",
	Subjects: []string{
		"gismo.person",
		"gismo.organization",
	},
	Storage: nats.FileStorage,
}

/*
cf. https://docs.nats.io/nats-concepts/jetstream/consumers
*/
var natsPersonConsumerConfig = nats.ConsumerConfig{
	AckPolicy: nats.AckExplicitPolicy,
	Durable:   "inboxPerson",
	//DeliverSubject makes it a push based consumer
	//this must be different per consumer
	//reason: messages are republished by consumer using this subject
	//make sure you have subscribe permission to that subject too
	DeliverSubject: "inboxPersonDeliverSubject", // makes it is a push based consumer
	AckWait:        time.Minute,                 // resend if ack was not received within this time
	/*
		when more than 1, messages can be delivered out of order
		when they need to be redelivered
	*/
	MaxAckPending: 1,
	FilterSubject: "gismo.person",
	DeliverPolicy: nats.DeliverAllPolicy,
}

var natsOrganizationConsumerConfig = nats.ConsumerConfig{
	AckPolicy: nats.AckExplicitPolicy,
	Durable:   "inboxOrganization",
	//DeliverSubject makes it a push based consumer
	//this must be different per consumer
	//reason: messages are republished by consumer using this subject
	//make sure you have subscribe permission to that subject too
	DeliverSubject: "inboxOrganizationDeliverSubject", // makes it is a push based consumer
	AckWait:        time.Minute,                       // resend if ack was not received within this time
	/*
		when more than 1, messages can be delivered out of order
		when they need to be redelivered
	*/
	MaxAckPending: 1,
	FilterSubject: "gismo.organization",
	DeliverPolicy: nats.DeliverAllPolicy,
}

func initInboxStream(js nats.JetStreamContext) error {
	stream, _ := js.StreamInfo(natsStreamConfig.Name)

	if stream == nil {
		_, err := js.AddStream(&natsStreamConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func initConsumer(js nats.JetStreamContext, streamName string, consumerConfig *nats.ConsumerConfig) error {
	if err := initInboxStream(js); err != nil {
		return err
	}

	consumerInfo, _ := js.ConsumerInfo(
		streamName, consumerConfig.Durable,
	)

	if consumerInfo == nil {
		_, err := js.AddConsumer(streamName, consumerConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func natsConnect(config ConfigNats) (*nats.Conn, error) {
	options := nats.Options{
		Url:                  config.Url,
		MaxReconnect:         100, // try reconnect n times, and then give up
		RetryOnFailedConnect: true,
		ReconnectWait:        10 * time.Second,
		Timeout:              10 * time.Second, // connection timeout
		AllowReconnect:       true,
	}

	/*
		IMPORTANT: error "nkeys not supported by the server" if there are no users
		configured with nkey
	*/
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
		logger.Errorf("Client connection to NATS closed, and was unable to reconnect (num reconnections: %d): %s", c.Reconnects, err)
	}
	options.ReconnectedCB = func(c *nats.Conn) {
		logger.Infof("Client connection to NATS restored")
	}
	options.ClosedCB = func(c *nats.Conn) {
		logger.Infof("Client connection to NATS closed")
	}

	return options.Connect()
}

func newOrganizationSubscriber(js nats.JetStreamContext) (subscribers.Subcriber, error) {
	if err := initConsumer(js, natsStreamConfig.Name, &natsOrganizationConsumerConfig); err != nil {
		return nil, fmt.Errorf("unable to create nats consumer %s: %w", natsOrganizationConsumerConfig.Durable, err)
	}
	orgSConfig := subscribers.GismoOrganizationConfig{}
	gismoProcessor, err := newGismoProcessor()
	if err != nil {
		return nil, err
	}
	orgSConfig.GismoProcessor = gismoProcessor
	orgSConfig.Subject = natsOrganizationConsumerConfig.FilterSubject
	orgSConfig.SubOpts = []nats.SubOpt{nats.Bind(natsStreamConfig.Name, natsOrganizationConsumerConfig.Durable)}
	return subscribers.NewGismoOrganizationSubscriber(orgSConfig), nil
}

func newPersonSubscriber(js nats.JetStreamContext) (subscribers.Subcriber, error) {
	if err := initConsumer(js, natsStreamConfig.Name, &natsPersonConsumerConfig); err != nil {
		return nil, fmt.Errorf("unable to create nats consumer %s: %w", natsPersonConsumerConfig.Durable, err)
	}
	personSConfig := subscribers.GismoPersonConfig{}
	gismoProcessor, err := newGismoProcessor()
	if err != nil {
		return nil, err
	}
	personSConfig.GismoProcessor = gismoProcessor
	personSConfig.Subject = natsPersonConsumerConfig.FilterSubject
	personSConfig.SubOpts = []nats.SubOpt{nats.Bind(natsStreamConfig.Name, natsPersonConsumerConfig.Durable)}

	return subscribers.NewGismoPersonSubscriber(personSConfig), nil
}

func newGismoProcessor() (*gismo.Processor, error) {
	repo, err := newRepository()
	if err != nil {
		return nil, err
	}
	ugentLdapClient := newUgentLdapClient()
	return gismo.NewProcessor(repo, ugentLdapClient), nil
}
