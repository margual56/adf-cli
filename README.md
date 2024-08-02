# adf-cli

## Installation

```bash
go install github.com/margual56/adf-cli
```

## Usage
```bash
> adf-cli -h
With this CLI you can manage triggers from a factory in Azure Data Factory. More features will be added in the future.

Usage:
  adf-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  trigger

Flags:
      --factoryName string         The factory name.
  -h, --help                       help for adf-cli
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.

Use "adf-cli [command] --help" for more information about a command.
```

## Commands

### Trigger

```bash
> adf-cli trigger -h
Usage:
  adf-cli trigger [flags]
  adf-cli trigger [command]

Available Commands:
  get         Display the properties of a trigger by name.
  list        List all triggers in a factory.
  start       Start a trigger in a data factory.
  stop        Stop a trigger in a data factory.

Flags:
  -h, --help   help for trigger

Global Flags:
      --factoryName string         The factory name.
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.

Use "adf-cli trigger [command] --help" for more information about a command.
```

#### List

```bash
> adf-cli trigger list -h
List all triggers in a factory.

Usage:
  adf-cli trigger list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --factoryName string         The factory name.
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.
```

#### Get

```bash
> adf-cli trigger get -h
Display the properties of a trigger by name.

Usage:
  adf-cli trigger get <triggerName> [flags]

Flags:
  -h, --help   help for get

Global Flags:
      --factoryName string         The factory name.
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.
```

#### Start

```bash
> adf-cli trigger start -h
Start a trigger in a data factory.

Usage:
  adf-cli trigger start <triggerName> [flags]

Flags:
  -h, --help   help for start

Global Flags:
      --factoryName string         The factory name.
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.
```

#### Stop

```bash
> adf-cli trigger stop -h
Stop a trigger in a data factory.

Usage:
  adf-cli trigger stop <triggerName> [flags]

Flags:
  -h, --help   help for stop

Global Flags:
      --factoryName string         The factory name.
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.
```

### Parameters

```bash
> adf-cli parameter -h
Usage:
  adf-cli param [command] [flags]
  adf-cli param [command]

Available Commands:
  list        List all global parameters from a factory
  update      Update a global parameter in a data factory

Flags:
  -h, --help   help for param

Global Flags:
      --factoryName string         The factory name.
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.

Use "adf-cli param [command] --help" for more information about a command.
```

#### List

```bash
> adf-cli parameter list -h
List all global parameters from a factory

Usage:
  adf-cli param list <globalParameterName> [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --factoryName string         The factory name.
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.
```

#### Update

```bash
> adf-cli parameter update -h
Create or update a global parameter in a data factory

Usage:
  adf-cli param update <globalParameterName> [flags]

Flags:
      --group string   The group the parameter belongs to (default "default")
  -h, --help           help for update
      --type string    The type of the parameter (default "string")
      --value string   The value of the parameter

Global Flags:
      --factoryName string         The factory name.
      --resourceGroupName string   The resource group name.
      --subscriptionId string      The subscription identifier.
```


## Configuration

The CLI uses the Azure SDK for Go to interact with the Azure Data Factory service. The SDK requires the following environment variables to be set:

- AZURE_CLIENT_ID
- AZURE_CLIENT_SECRET
- AZURE_TENANT_ID


You can also set the flags factoryName, resourceGroupName, and subscriptionId as environment variables:
```bash
FACTORY_NAME=your-factory-name
RESOURCE_GROUP_NAME=your-resource-group-name
SUBSCRIPTION_ID=your-subscription-id
```

## Roadmap

- [ ] Add operations for triggers.
  - [x] Get a trigger
  - [x] List Triggers By Factory
  - [x] Start a trigger
  - [x] Stop a trigger
  - [ ] Create or Update
  - [ ] Delete
  - [ ] Get Event Subscription Status
  - [ ] Query By Factory
  - [ ] Subscribe To Events
  - [ ] Unsubscribe From Events
  - [ ] Cancel a single trigger run by runId
  - [ ] Query Trigger Runs By Factory
  - [ ] Rerun a single trigger run by runId
- [ ] Add operations for Global Parameters
  - [x] Create or Update
  - [ ] Delete
  - [x] Get
  - [x] List By Factory


## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
