# people-service
Person service

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

* `PEOPLE_API_KEY`

  type: `string`

  default: `people`

  description: api key. Used in authentication header `X-Api-Key` for server (openapi)

* `PEOPLE_API_PORT`

  type: `int`

  default: `3999`

  description: port address for server

* `PEOPLE_API_HOST`

  type : `string`

  default: `127.0.0.1`

  description: host address for server

* `PEOPLE_DB_URL`

  type: `string`

  default: `postgres://people:people@localhost:5432/authority?sslmode=disable`

  description: postgres database connection url

* `PEOPLE_DB_AES_KEY`

  type: `string`

  required: `true`

  description: [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard) key. This is now used to encrypt attribute `orcid_token`. Note that an AES key must be 128 bits long (or 16 characters). Generate one with command `openssl rand -hex 16`

* `PEOPLE_NATS_URL`

  type: `string`

  default: `nats://127.0.0.1:4222`

  description: NATS connection url

* `PEOPLE_NATS_NKEY`

  type: `string`

  description: [Ed25519](https://ed25519.cr.yp.to/) public key

  In order to generate a key pair do the following:

  ```
  git clone https://github.com/nats-io/nkeys
  cd nkeys/nk
  go build
  ./nk -gen user -pubout
  ```

  This command outputs two lines, the first containing the seed or private
  (starting with `S`), followed by the public key

  Add the public key to the nats user configuration (`nats.conf`):

  ```
  authorization {
    default_permissions = {}
    PEOPLE_SERVICE = {
      subscribe = [
        "gismo.person",
        "gismo.organization",
        "inboxOrganizationDeliverSubject",
        "inboxPersonDeliverSubject",
        "_INBOX.>",
        "$JS.API.STREAM.CREATE.gismo",
        "$JS.API.STREAM.UPDATE.gismo",
        "$JS.API.STREAM.INFO.gismo",
        "$JS.API.STREAM.DELETE.gismo",
        "$JS.API.CONSUMER.DURABLE.CREATE.gismo.>",
        "$JS.API.CONSUMER.DELETE.gismo.>",
        "$JS.API.CONSUMER.INFO.gismo.>"
      ]
    }
    users = [
      { nkey: "<public-key>", permissions: $PEOPLE_SERVICE }
    ]
  }
  ```

* `PEOPLE_NATS_NKEY_SEED`

  type: `string`

  description: [Ed25519](https://ed25519.cr.yp.to/) private key

  Keep the private key close to your application.

  Nats does not need to know about this.

* `PEOPLE_NATS_STREAM_NAME`

  type: `string`

  required: `true`

  description: name of stream in NATS. Stream will be created automatically. Also serves as prefix (with '.' after it) for every subject created on it (including delivery subjects).

  e.g. `gismoDev` creates subjects `gismoDev.person`, `gismoDev.organization` and delivery subjects `gismoDev.inboxPersonDeliverSubject` and `gismoDev.inboxOrganizationDeliverSubject`

* `PEOPLE_LDAP_URL`

  type: `string`

  description: ldap connection url. e.g. `ldaps://ldaps.ugent.be:636`

  required: `true`

* `PEOPLE_LDAP_USERNAME`

  type: `string`

  description: ldap username

  required: `true`

  Note: internally we bind to scope `ou=people,dc=ugent,dc=be`, so make sure these
  credentials are valid for that scope.

* `PEOPLE_LDAP_PASSWORD`

  type: `string`

  description: ldap password

  required: `true`

# Start the api server (openapi)

```
$ ./people-service server
```

# Start NATS consumer organization

```
$ ./people-service inbox listen organization
```

* creates NATS stream with name as configured by environment variable `PEOPLE_NATS_STREAM_NAME` with subjects `{streamName}.organization` and `{streamName}.person`. The default stream name is `gismo`. If already present, does not try to change it. Expections are:

  * messages must be persisted to disk

  * messages are created by [soap-bridge](https://github.com/ugent-library/soap-bridge) and are SOAP Cerif XML messages as produced by GISMO. We have to convert those XML messages into this `models.Message` format:

    ```
    {
        "id": "<gismo organization identifier>",
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

* creates NATS push consumer `inboxOrganization` on stream. If already present, does not try to change it. Expectations are:

  * DeliverSubject: `inboxOrganizationDeliverSubject`. This option makes it a push based consumer.

  * Durable: `inboxOrganization`

  * AckPolicy: explicit. So no automatic acknowledgments.

  * MaxAckPending: 1. When higher than 1, pending messages may be sent out of order by the server.

  * AckWait: 1 minute. This means that messages are resent after 1 minute when no acknowledgment was received. Putting this too low puts pressure on the processing.

* binds to consumer `inboxOrganization` on stream `{streamName}`

* listens to subject `{streamName}.organization`

* processes SOAP XML messages in order

* messages with malformed XML are discarded. Acknowledgment is sent to ensure progression.

* messages with unexpected XML structure are discarded, and acknowledged to ensure progression.

* XML messages are converted to `models.Message` format.

* `models.Message` results into insert or update of a organization record

* Notes:

  * if GISMO states that the organization is child of a parent organization, and that parent organization does not exist yet, then that parent organization is made automatically with only a gismo_id as known attribute


# Start NATS consumer person

```
$ ./people-service inbox listen person
```

* creates NATS stream `{streamName}` with subjects `{streamName}.person`. If already present, does not try to change it. Expections are:

  * messages must be persisted to disk

  * messages are created by [soap-bridge](https://github.com/ugent-library/soap-bridge) and are SOAP Cerif XML messages as produced by GISMO. We have to convert those XML messages into this `models.Message` format:

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

* creates NATS push consumer `inboxPerson` on stream `{streamName}`. If already present, does not try to change it. Expectations are:

  * DeliverSubject: `inboxPersonDeliverSubject`. This option makes it a push based consumer.

  * Durable: `inboxPerson`

  * AckPolicy: explicit. So no automatic acknowledgments.

  * MaxAckPending: 1. When higher than 1, pending messages may be sent out of order by the server.

  * AckWait: 1 minute. This means that messages are resent after 1 minute when no acknowledgment was received. Putting this too low puts pressure on the processing.

* binds to consumer `inboxPerson` on stream `{streamName}`

* listens to subject `{streamName}.person`

* processes SOAP XML messages in order

* messages with malformed XML are discarded. Acknowledgment is sent to ensure progression.

* messages with unexpected XML structure are discarded, and acknowledged to ensure progression.

* XML messages are converted to `models.Message` format.

* `models.Message` results into insert or update of a person record

* Notes:

  * If GISMO states that a person belongs to an organization that does not exist
    yet in the database, then an organization is made automatically with that organization's gismo_id as only known attribute.
  * Every GISMO person message needs to have an ugent_id. Without an ugent_id one cannot
    link with the ugent ldap.

# run in docker

Build base docker image `people-service`:

```
$ docker build -t ugentlib/people-service .
$ docker push ugentlib/people-service
```

If image `people-service` is already docker github,

then you may skip that step.

Start set of services using `docker compose`:

```
$ docker compose up
```

Docker compose uses that image `people-service`
