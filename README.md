# Conduit Connector for <!-- readmegen:name -->Google-Drive<!-- /readmegen:name -->

[Conduit](https://conduit.io) connector for <!-- readmegen:name -->Google-Drive<!-- /readmegen:name -->.

<!-- readmegen:description -->
The Google Drive connector is one of [Conduit](https://github.com/ConduitIO/conduit)'s plugins. It provides a **destination connector** for writing records into a specified Google Drive folder.

## Destination

The Google Drive Destination Connector connects to a Google Drive account using a service account's credentials and uploads incoming records as files into a configured folder.

### Configuration

The connector requires the following configuration parameters for authentication and folder access:

- `drive.projectId`: The Google Cloud project ID.
- `drive.privateKeyId`: The private key ID of the service account.
- `drive.privateKey`: The private key associated with the service account.
- `drive.clientEmail`: The client email of the service account.
- `drive.clientId`: The client ID of the service account.
- `drive.clientCertUrl`: The client certificate URL.
- `drive.folderId`: The ID of the target Google Drive folder where files will be uploaded.

All fields are required. If the credentials are invalid or the specified folder is inaccessible, the connector will fail to initialize.

### Authentication

This connector uses a **Google service account** for authentication. Ensure that the service account has write access to the target Drive folder by either:

- Sharing the folder directly with the service account email, or  
- Using domain-wide delegation (if operating within a Google Workspace organization)<!-- /readmegen:description -->


### Example Configuration

<!-- readmegen:destination.parameters.yaml -->
```yaml
version: 2.2
pipelines:
  - id: example
    status: running
    connectors:
      - id: example
        plugin: "google-drive"
        settings:
          # The URL to the X.509 certificate for the service account, used to
          # verify its identity.
          # Type: string
          # Required: yes
          drive.clientCertUrl: ""
          # The email address of the service account (e.g.
          # my-service-account@project.iam.gserviceaccount.com).
          # Type: string
          # Required: yes
          drive.clientEmail: ""
          # The OAuth2 client ID associated with the service account.
          # Type: string
          # Required: yes
          drive.clientId: ""
          # The ID of the Google Drive folder where records will be uploaded.
          # This can be found in the folder's URL:
          # https://drive.google.com/drive/folders/<folderId>
          # Type: string
          # Required: yes
          drive.folderId: ""
          # The private key (PEM-encoded) used to sign service account requests.
          # Type: string
          # Required: yes
          drive.privateKey: ""
          # The ID of the private key used to authenticate the service account.
          # Type: string
          # Required: yes
          drive.privateKeyId: ""
          # The Google Cloud project ID associated with the service account.
          # Type: string
          # Required: yes
          drive.projectId: ""
          # Maximum delay before an incomplete batch is written to the
          # destination.
          # Type: duration
          # Required: no
          sdk.batch.delay: "0"
          # Maximum size of batch before it gets written to the destination.
          # Type: int
          # Required: no
          sdk.batch.size: "0"
          # Allow bursts of at most X records (0 or less means that bursts are
          # not limited). Only takes effect if a rate limit per second is set.
          # Note that if `sdk.batch.size` is bigger than `sdk.rate.burst`, the
          # effective batch size will be equal to `sdk.rate.burst`.
          # Type: int
          # Required: no
          sdk.rate.burst: "0"
          # Maximum number of records written per second (0 means no rate
          # limit).
          # Type: float
          # Required: no
          sdk.rate.perSecond: "0"
          # The format of the output record. See the Conduit documentation for a
          # full list of supported formats
          # (https://conduit.io/docs/using/connectors/configuration-parameters/output-format).
          # Type: string
          # Required: no
          sdk.record.format: "opencdc/json"
          # Options to configure the chosen output record format. Options are
          # normally key=value pairs separated with comma (e.g.
          # opt1=val2,opt2=val2), except for the `template` record format, where
          # options are a Go template.
          # Type: string
          # Required: no
          sdk.record.format.options: ""
          # Whether to extract and decode the record key with a schema.
          # Type: bool
          # Required: no
          sdk.schema.extract.key.enabled: "true"
          # Whether to extract and decode the record payload with a schema.
          # Type: bool
          # Required: no
          sdk.schema.extract.payload.enabled: "true"
```
<!-- /readmegen:destination.parameters.yaml -->

## Development

- To install the required tools, run `make install-tools`.
- To generate code (mocks, re-generate `connector.yaml`, update the README,
  etc.), run `make generate`.

## How to build?

Run `make build` to build the connector.

## Testing

Run `make test` to run all the unit tests. Run `make test-integration` to run
the integration tests.

The Docker compose file at `test/docker-compose.yml` can be used to run the
required resource locally.

## How to release?

The release is done in two steps:

- Bump the version in [connector.yaml](/connector.yaml). This can be done
  with [bump_version.sh](/scripts/bump_version.sh) script, e.g.
  `scripts/bump_version.sh 2.3.4` (`2.3.4` is the new version and needs to be a
  valid semantic version). This will also automatically create a PR for the
  change.
- Tag the connector, which will kick off a release. This can be done
  with [tag.sh](/scripts/tag.sh).

## Known Issues & Limitations

- Known issue A
- Limitation A

## Planned work

- [ ] Item A
- [ ] Item B
