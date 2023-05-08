# people
People service

# Build process

```
go build
```

# Prepare environment variables

Declare the following environment variables,

or store them in file `.env` in the root of your folder (important: exclude `export` statement)
(see .env.example):

* `PEOPLE_PRODUCTION`:

  type: `bool`

  default: `false`

* `PEOPLE_API_USERNAME`

  type: `string`

  default: `people`

  description: username for basic authentication. Used in authentication for GPRC api

* `PEOPLE_API_PASSWORD`

  type: `string`

  default: `people`

  description: password for basic authentication. Used in authentication for GPRC api

* `PEOPLE_API_PORT`

  type: `int`

  default: `3999`

  description: port address for GRPC api

* `PEOPLE_API_HOST`

  type : `string`

  default: `127.0.0.1`

  description: host address for GRPC api

* `PEOPLE_API_TLS_ENABLED`

  type: `bool`

  default: `false`

  description: start GRPC server with TLS support. Requires options `api.tls_server_cert` and `api.tl_server_key`.

* `PEOPLE_API_TLS_SERVER_CERT`

  type: `string`

  description: file location of server certificate

* `PEOPLE_API_TLS_SERVER_KEY`

  type: `string`

  description: file location of server private key

* `PEOPLE_API_CLIENT_INSECURE`

  type: `bool`

  default: `false`

  description: set to false when your GRPC is running without TLS support

* `PEOPLE_API_CLIENT_CACERT`

  type: `string`

  description: file location to root certificate. Only needed when GRPC is running with TLS support

* `PEOPLE_API_PROXY_HOST`

  type: `string`

  description: host address for api proxy

  default: `localhost`

* `PEOPLE_API_PROXY_PORT`

  type: `int`

  description: proxy address for api proxy

  default: `4001`

* `PEOPLE_DB_URL`

  type: `string`

  default: `postgres://people:people@localhost:5432/authority?sslmode=disable`

  description: postgres database connection url

* `PEOPLE_NATS_URL`

  type: `string`

  default: `nats://127.0.0.1:4222`

  description: NATS connection url

# Start GRPC api

```
$ ./people api start
```

# Start GRPC gateway


```
$ ./people api proxy
```

Starts a JSON api at address `localhost:4001` with the following routes:

* `POST /api.v1.People/GetAllPerson`

No request body is expected

* `POST /api.v1.People/GetPerson`

Expected is a JSON request body. In this case JSON with attribute "id"

# Start NATS consumer

```
$ ./people inbox listen
```

* creates NATS stream `PEOPLE` with subjects `person.update`. If already present, does not try to change it. Expections are:

  * messages must be persisted to disk

  * messages are created by [soap-bridge](https://github.com/ugent-library/soap-bridge) and look as [following](https://github.com/ugent-library/soap-bridge/blob/main/main.go#L59):

    ```
    {
        "id": "<person identifier>",
        "language": "<iso language code>",
        "attributes": [
            {
                "name": "<attribute name>",
                "value": "<attribute value>",
                "start_date": "<value valid from start date>",
                "end_date": "<value valid till end date>"
            },
            ...
        ]
    }
    ```

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

# Commands

```
$ ./people person get --id <person-id>
```

```
$ ./people person all
```

```
$ ./people person suggest --query nic
```

```
$ ./people organization get --id <org-id>
```

```
$ ./people organization all
```

```
$ ./people organization suggest --query nic
```
