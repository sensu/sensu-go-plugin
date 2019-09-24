# Sensu Go CHANGEME Plugin
[![Bonsai Asset Badge](https://img.shields.io/badge/CHANGEME-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/CHANGEME/CHANGEME) [![TravisCI Build Status](https://travis-ci.org/CHANGEME/sensu-CHANGEME.svg?branch=master)](https://travis-ci.org/CHANGEME/sensu-CHANGEME)

TODO: Description.

## Configuration

### Asset Registration

Assets are the best way to make use of this plugin. If you're not using an asset, please consider doing so! If you're using sensuctl 5.13 or later, you can use the following command to add the asset: 

`sensuctl asset add CHANGEME/sensu-CHANGEME:VERSION`

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/CHANGEME/sensu-CHANGEME).

Example Sensu Go definition:

```json
{
    "api_version": "core/v2",
    "type": "CHANGEME",
    "metadata": {
        "namespace": "default",
        "name": "CHANGEME"
    },
    "spec": {
        "...": "..."
    }
}
```

## Installation from source and contributing

The preferred way of installing and deploying this plugin is to use it as an [asset]. If you would like to compile and install the plugin from source or contribute to it, download the latest version of the sensu-CHANGEME from [releases][1]
or create an executable script from this source.

From the local path of the sensu-CHANGEME repository:

```
go build -o /usr/local/bin/sensu-CHANGEME main.go
```

## Usage Examples

Help:

```
The Sensu Go CHANGEME for x

Usage:
  sensu-CHANGEME [flags]

Flags:
  -f, --foo string   example
  -h, --help         help for sensu-CHANGEME
```

## Contributing

See https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/CHANGEME/sensu-CHANGEME/releases
[2]: #asset-registration
