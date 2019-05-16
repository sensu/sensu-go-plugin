# Sensu CHANGEME Plugin

[![Bonsai Asset Badge](https://img.shields.io/badge/sensu--CHANGEME--plugin-Download%20on%20Bonsai-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/CHANGEME/CHANGEME-email-plugin) [![TravisCI Build Status](https://travis-ci.org/sensu/sensu-email-handler.svg?branch=master)](https://travis-ci.org/sensu/sensu-email-handler)

TODO: Description

<!-- For a repository containing one Sensu plugin, use the standard plugin README template: https://github.com/sensu/sensu-go-plugin/blob/master/README.md -->

- [Plugin 1](#plugin-1)
- [Plugin 2](#plugin-2)
- [Installation](#installation)
- [Contributing](#contributing)

## Service checks
<!-- Include this section if the plugins included in this collection perform service checks -->

[`plugin 1`](#plugin-1)
- **description**: Monitors the service `/healthz` endpoint to confirm the service is healthy.
- **output**: status information

[`plugin 2`](#plugin-2)
- **description**: The check description should describe the service health checking capability provided by the plugin.
- **output**: status information

## Plugin 1

TODO: Description

### Metrics
<!-- Include this section if the plugin collects metrics. -->

`metric_name`
- **description**: The metric description should clearly explain what data is
  being measured, and provide some information regarding what it should mean to the user.
- **unit**: unit specification

`example.error_rate`
- **description**: The percentage of error responses, calculated as "count of
  error responses / total count of responses".
- **unit**: percentile

### Usage

```
The Sensu Go CHANGEME for x

Usage:
  sensu-CHANGEME [flags]

Flags:
  -f, --foo string   example
  -h, --help         help for sensu-CHANGEME
```

#### Dependencies
<!-- List dependencies required by the plugin, or write "none" -->


### Configuration

See the [Sensu docs](https://docs.sensu.io/sensu-go/latest/reference/) for more information about creating resources in Sensu.

#### Example Sensu check definition
<!-- Include an example mutator definition for check and handler plugins. -->

```yml
type: CheckConfig
api_version: core/v2
metadata:
  name: example
  namespace: default
spec:
  command: sensu-CHANGEME
  env_vars: null
  handlers:
    - 
  high_flap_threshold: 0
  interval: 60
  low_flap_threshold: 0
  output_metric_format: ""
  output_metric_handlers:
    - 
  proxy_entity_name: ""
  publish: true
  round_robin: false
  runtime_assets:
    - 
  stdin: false
  subscriptions:
  - 
```

#### Example Sensu handler definition
<!-- Include an example mutator definition for handler, mutator, and filter plugins. -->

```yml
type: Handler
api_version: core/v2
metadata:
  name: example
  namespace: default
spec:
  type: pipe
  command: sensu-CHANGEME
  env_vars:
    - 
  filters:
    - 
  runtime_assets:
    - 
```

#### Example Sensu filter definition
<!-- Include an example mutator definition only for filter plugins. -->

```yml
type: EventFilter
api_version: core/v2
metadata:
  name: example
  namespace: default
spec:
  action: allow
  expressions:
  - 
  runtime_assets:
  - 
```

#### Example Sensu mutator definition
<!-- Include an example mutator definition only for mutator plugins. -->

```yml
type: Mutator
api_version: core/v2
metadata:
  name: example
  namespace: default
spec:
  command: sensu-CHANGME
  env_vars:
    - 
  runtime_assets:
    - 
```

## Plugin 2

TODO: Description

### Metrics
<!-- Include this section if the plugin collects metrics. -->

`metric_name`
- **description**: The metric description should clearly explain what data is
  being measured, and provide some information regarding what it should mean to the user.
- **unit**: unit specification

`example.error_rate`
- **description**: The percentage of error responses, calculated as "count of
  error responses / total count of responses".
- **unit**: percentile

### Usage

```
The Sensu Go CHANGEME for x

Usage:
  sensu-CHANGEME [flags]

Flags:
  -f, --foo string   example
  -h, --help         help for sensu-CHANGEME
```

#### Dependencies
<!-- List dependencies required by the plugin, or write "none" -->


### Configuration

See the [Sensu docs](https://docs.sensu.io/sensu-go/latest/reference/) for more information about creating resources in Sensu.

#### Example Sensu check definition
<!-- Include an example mutator definition for check and handler plugins. -->

```yml
type: CheckConfig
api_version: core/v2
metadata:
  name: example
  namespace: default
spec:
  command: sensu-CHANGEME
  env_vars: null
  handlers:
    - 
  high_flap_threshold: 0
  interval: 60
  low_flap_threshold: 0
  output_metric_format: ""
  output_metric_handlers:
    - 
  proxy_entity_name: ""
  publish: true
  round_robin: false
  runtime_assets:
    - 
  stdin: false
  subscriptions:
  - 
```

#### Example Sensu handler definition
<!-- Include an example mutator definition for handler, mutator, and filter plugins. -->

```yml
type: Handler
api_version: core/v2
metadata:
  name: example
  namespace: default
spec:
  type: pipe
  command: sensu-CHANGEME
  env_vars:
    - 
  filters:
    - 
  runtime_assets:
    - 
  timeout: 0
```

#### Example Sensu filter definition
<!-- Include an example mutator definition only for filter plugins. -->

```yml
type: EventFilter
api_version: core/v2
metadata:
  name: example
  namespace: default
spec:
  action: allow
  expressions:
  - 
  runtime_assets:
  - 
```

#### Example Sensu mutator definition
<!-- Include an example mutator definition only for mutator plugins. -->

```yml
type: Mutator
api_version: core/v2
metadata:
  name: example
  namespace: default
spec:
  command: sensu-CHANGME
  env_vars:
    - 
  runtime_assets:
    - 
  timeout: 0
```

## Installation

Download the latest version of the sensu-CHANGEME from [Bonsai](http://bonsai.sensu.io/assets/CHANGEME/CHANGEME),
or create an executable script from this source.

From the local path of the sensu-CHANGEME repository:

```
go build -o /usr/local/bin/sensu-CHANGEME main.go
```

## Contributing

See the [Sensu Go contributing guide](https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md).
