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

# Run database migrations

Before starting the application you should run any pending database migrations.

For this we use [atlas](https://atlasgo.io/), which can be installed via homebrew on your MAC:

```
brew install ariga/tap/atlas
```

Or manually

```
curl -sSf https://atlasgo.sh | sh
```

Atlas is a database migration tool, that is used both for generating the necessary
sql migration files (during development), and for managing the state of the running database (in production).

Copy the example file `atlas.hcl.example` to `atlas.hcl` and update where necessary

Now run `atlas migrate apply --env local` to apply any pending migrations

If you change anything to your entgo schema, run this atlas command:

```
atlas migrate diff <migration-name> --env local
```

This will generate new sql files in folder `ent/migrate/migrations`
You can update these by hand, but be sure to run `atlas hash --env local` afterwards
to update the checksum file `atlas.sum`

Due to some [inconveniences](https://github.com/ariga/atlas/issues/2158) we can only use atlas during development, in order to generate migration files.
These are stored in `ent/migrate/migrations`

In production we use [tern](https://github.com/jackc/tern). Make sure
that directories `ent/migrate/migrations` (for atlas) and `etc/migrations` (for tern) are kept in sync. And note that tern uses a different naming for sql
files (prefix is a padded number instead of a timestamp)

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

# Update flow organizations

## incoming message via NATS consumer

Any nats message on subject `${PEOPLE_NATS_STREAM_NAME}.organization` is processed
by the command `people-service inbox listen organization`. Expected is a CERIF XML
blob. The message translates this XML into a list of attributes (with start and end date), and assigns them to a new/existing organization record with the following rules:

* if there is a record that matches the gismo identifier, then it is loaded. Otherwise a new record is created.

* if the incoming message has as action `DELETE`, then the record is deleted. Other actions are treated as an upsert.

* Existing attributes `name_dut`, `name_dut`, `parent_id`, `type` and `identifier` are cleared before assigning attributes. So every update from GISMO should send all attributes. No merging happens.

* gismo identifier is added to list of external identifiers (column `identifier`) ALWAYS

* `parent_id` is added if current date is between start and end date. To assign this, the parent organization is searched by its gismo identifier and assigned as such. When not found, it is created with only the gismo identifier (dummy record). Later updates should fill this parent record with more information

* `name_dut` is added if current date is between start and end date. 

* `name_eng` is added if current date is between start and end date. 

* `type` is added ALWAYS

* `ugent_memorialis_id` is added ALWAYS to `identifier->'ugent_memorialis_id'`

* `code` is added ALWAYS to `identifier->'ugent_id'`

* `biblio_code` is added ALWAYS to `identifier->'biblio_id'`

Problems:

* There may be more than one attribute `parent_id`. Above code solves this by only using those attributes that are still within range. But for deprecated organizations (like our older departments) that means that they do not have a parent organization. e.g. LW06

## call to openapi `/api/v1/add-organization`

Any call to the the openapi handler `/api/v1/add-organization` overwrites
the existing record. Note that only updates that contain an `id` are considered
updates, and those without as new records (for who a new id is minted).

Any call to this handler should happen rarely.

The route via the nats consumer is preferred.

# Update flow people

## incoming message via NATS consumer

Any nats message on subject `${PEOPLE_NATS_STREAM_NAME}.person` is processed
by the command `people-service inbox listen person`. Expected is a CERIF XML
blob. The message translates this XML into a list of attributes (with start and end date), and assigns them to a new/existing person record with the following rules:

* if there is a record that matches the gismo identifier or any of the provided historic_ugent_id's, then it is loaded. Otherwise a new record is created.

* every message requires one or more historic_ugent_id's. Without this we cannot student records that were inserted previously via the student importer (via ldap) and that have no gismo identifier yet.

* Existing attributes `email`, `given_name`, `family_name`, `name`, `preferred_given_name`, `preferred_family_name`, `organization`, `honorific_prefix` and `identifier` are cleared before assigning attributes. So every update from GISMO should send all attributes. No merging happens. For every `organization` mentioned, it makes sure that a organization record with that gismo id exists in the table for organizations, even if it means making a dummy record. That person is linked to the list of provided organizations.

* The person record is enriched with attributes from the ugent ldap (if any). If no ldap entry can be found, the attribute `active` is set to `false`. The following attributes are set from ldap:
  * `identifier->'ugent_username'`
  * `identifier->'ugent_barcode'`
  * `name`
  * `object_class`
  * `expiration_date`. 

Problems:

* a call to openapi handler `/api/v1/set-person-orcid` may also set orcid, and can be cleared by the next incoming message from gismo

## nightly cron job that upserts student records

A nightly cron job reads student record from the ugent ldap

and inserts/updates person records. Matching on existing records

is done by matching on identifier `historic_ugent_id`.

The following attributes are overwritten:

* `identifier->'ugent_username'`
* `identifier->'ugent_id'`
* `identifier->'historic_ugent_id'`
* `identifier->'ugent_barcode'`
* `given_name`
* `family_name`
* `name`
* `birth_date`
* `email`
* `job_category`
* `honorific_prefix`
* `object_class`
* `expiration_date`
* `organization`. 

Note that if no organization be found based on `identifier->'ugent'` then no (dummy) organization record is made for it. In that case the attribute is ignored.

## nightly cron job that deactivates expired person records

All person records with column `expiration_date` set to a non zero value,

and whose expiration date has passed, are deactivated, by setting column

`active` to `false`. The expiration date is cleared also.

Expiration dates are taken from the ugent ldap entry.

## call to openapi `/api/v1/add-person`