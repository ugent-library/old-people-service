package cmd

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

var natsStreamConfig = nats.StreamConfig{
	Name: "PEOPLE",
	Subjects: []string{
		"person.update",
		"person.delete",
		"organization.update",
		"organization.delete",
	},
	Storage: nats.FileStorage,
}

var natsPersonConsumerConfig = nats.ConsumerConfig{
	AckPolicy: nats.AckExplicitPolicy,
	Durable:   "inboxPerson",
	//DeliverSubject makes it a push based consumer
	//this must be different per consumer
	//reason: messages are republished by consumer using this subject
	//make sure you have subscribe permission to that subject too
	DeliverSubject: "inboxPersonDeliverSubject", // makes it is a push based consumer
	AckWait:        time.Second * 10,
	/*
		when more than 1, messages can be delivered out of order
		when they need to be redelivered
	*/
	MaxAckPending: 1,
	FilterSubject: "person.*",
}

var natsOrganizationConsumerConfig = nats.ConsumerConfig{
	AckPolicy: nats.AckExplicitPolicy,
	Durable:   "inboxOrganization",
	//DeliverSubject makes it a push based consumer
	//this must be different per consumer
	//reason: messages are republished by consumer using this subject
	//make sure you have subscribe permission to that subject too
	DeliverSubject: "inboxOrganizationDeliverSubject", // makes it is a push based consumer
	AckWait:        time.Second * 10,
	/*
		when more than 1, messages can be delivered out of order
		when they need to be redelivered
	*/
	MaxAckPending: 1,
	FilterSubject: "organization.*",
}

var inboxPersonStreamSubjects = []string{"person.update", "person.delete"}

var inboxOrganizationStreamSubjects = []string{"organization.update", "organization.delete"}

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

func ensureAck(msg *nats.Msg) {
	if err := msg.Ack(); err != nil {
		logger.Fatal(fmt.Errorf("unable to acknowledge nats message: %w", err))
	}
}

func natsConnect(config ConfigNats) (*nats.Conn, error) {
	options := make([]nats.Option, 0)

	/*
		IMPORTANT: error "nkeys not supported by the server" if there are no users
		configured with nkey
	*/
	if config.Nkey != "" && config.NkeySeed != "" {
		user, err := nkeys.FromSeed([]byte(config.NkeySeed))
		if err != nil {
			return nil, err
		}
		options = append(options, nats.Nkey(config.Nkey, func(nonce []byte) ([]byte, error) {
			return user.Sign(nonce)
		}))
	}

	return nats.Connect(config.Url, options...)
}
