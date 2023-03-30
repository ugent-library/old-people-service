# people
People service

# Build process

```
go build
```

# Prepare configuration

`cp people.toml.example people.toml`

## Options

**production**

type: `bool`

default: `false`

**api.username**

type: `string`

default: `people`

description: username for basic authentication. Used in authentication for GPRC api

**api.password**

type: `string`

default: `people`

description: password for basic authentication. Used in authentication for GPRC api

**api.port**

type: `int`

default: `3999`

description: port address for GRPC api

**api.host**

type : `string`

default: `127.0.0.1`

description: host address for GRPC api

**api.tls_enabled**

type: `bool`

default: `false`

description: start GRPC server with TLS support. Requires options `api.tls_server_cert` and `api.tl_server_key`.

**api.tls_server_cert**

type: `string`

description: file location of server certificate

**api.tls_server_key**

type: `string`

description: file location of server private key

**api_client.insecure**

type: `bool`

default: `false`

description: set to false when your GRPC is running without TLS support

**api_client.cacert**

type: `string`

description: file location to root certificate. Only needed when GRPC is running with TLS support

**db.url**

type: `string`

default: `postgres://biblio:biblio@localhost:5432/authority?sslmode=disable`

description: postgres database connection url

**nats.url**

type: `string`

default: `nats://127.0.0.1:4222`

description: NATS connection url

# Start GRPC api

```
$ ./people api start -c people.toml
```

# Start NATS consumer

```
$ ./people inbox listen -c people.toml
```

* creates NATS stream `PEOPLE` with subjects `person.update`. If already present, does not try to change it. Expections are:

  *  messages must be persisted to disk

* creates NATS push consumer `inbox` on stream `PEOPLE`. If already present, does not try to change it. Expectations are:

  * DeliverSubject: `inboxDeliverSubject`. This option makes it a push based consumer.

  * Durable: `inbox`

  * AckPolicy: explicit. So no automatic acknowledgments.

  * MaxAckPending: 1. When higher than 1, pending messages may be sent out of order by the server.

  * AckWait: 10 seconds. This means that messages are resent after 10 seconds when no acknowledgment was received

* binds to consumer `inbox` on stream `PEOPLE`

* listens to subject `person.update`

* processes JSON messages in order

* messages with malformed JSON are discarded. Acknowledgment is sent to ensure progression.

* messages with unexpected JSON structure are published to `inbox.rejected`. Acknowledgment is sent to ensure progression.

* messages that result in invalid person records are republished to subject `inbox.person.rejected`. Acknowledgment is sent to ensure progression.

* on successfull update of the person record, the updated person record is republished on subject `person.updated`. Messages like this contain the full person record in JSON as payload. Acknowledgment is sent when record is stored successfully.
