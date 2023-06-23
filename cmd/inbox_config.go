package cmd

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
	"github.com/ugent-library/person-service/models"
	"github.com/ugent-library/person-service/subscribers"
)

var natsStreamConfig = nats.StreamConfig{
	Name: "GISMO",
	Subjects: []string{
		"gismo.person",
		"gismo.organization",
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
	FilterSubject: "gismo.person",
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
	FilterSubject: "gismo.organization",
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

func ensureAck(msg *nats.Msg) {
	if err := msg.Ack(); err != nil {
		logger.Fatal(fmt.Errorf("unable to acknowledge nats message: %w", err))
	}
}

func natsConnect(config ConfigNats) (*nats.Conn, error) {
	options := nats.Options{
		Url:                  config.Url,
		MaxReconnect:         10,
		RetryOnFailedConnect: true,
		ReconnectWait:        10,
		Timeout:              30 * time.Second,
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

	return options.Connect()
}

func buildSubscribers(js nats.JetStreamContext, services *models.Services) ([]subscribers.Subcriber, error) {
	if err := initConsumer(js, natsStreamConfig.Name, &natsPersonConsumerConfig); err != nil {
		return nil, fmt.Errorf("unable to create nats consumer %s: %w", natsPersonConsumerConfig.Durable, err)
	}
	if err := initConsumer(js, natsStreamConfig.Name, &natsOrganizationConsumerConfig); err != nil {
		return nil, fmt.Errorf("unable to create nats consumer %s: %w", natsOrganizationConsumerConfig.Durable, err)
	}
	return []subscribers.Subcriber{
		subscribers.NewGismoOrganizationSubscriber(natsOrganizationConsumerConfig.FilterSubject, services.OrganizationService, nats.Bind(natsStreamConfig.Name, natsOrganizationConsumerConfig.Durable)),
		subscribers.NewGismoPersonSubscriber(natsPersonConsumerConfig.FilterSubject, services.PersonService, nats.Bind(natsStreamConfig.Name, natsPersonConsumerConfig.Durable)),
	}, nil
}
