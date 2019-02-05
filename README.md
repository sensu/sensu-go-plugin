# Sensu Go CHANGEME Plugin
TravisCI: [![TravisCI Build Status](https://travis-ci.org/CHANGEME/sensu-CHANGEME.svg?branch=master)](https://travis-ci.org/CHANGEME/sensu-CHANGEME)

TODO: Description.

## Instructions for Asset Collections

Before you create an asset, determine if it is a collection. How does one determine if an asset is collection?

1. What is an asset collection?

An asset collection is a large group of executables used for monitoring a service.

2. What are some examples of an asset collection?

* NOT a asset collection - [Pushbullet Plugin][2]
* IS a asset collection - [Sensu AWS Plugins][3]

3. Why mention asset collections?

Due to the nature of asset collections, it can be difficult to write documentation that clearly states how to use the collection. We want to ensure that asset documentation is clear and that it effectively guides users on using a collection.

4. What should I do if I have a collection?

Instead of using the examples below, we ask that you please use the [asset collection README example][4]

## Installation

Download the latest version of the sensu-CHANGEME from [releases][1],
or create an executable script from this source.

From the local path of the sensu-CHANGEME repository:

```
go build -o /usr/local/bin/sensu-CHANGEME main.go
```

## Configuration

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
[2]: https://github.com/rgeniesse/sensu-pushbullet-handler
[3]: https://github.com/sensu/sensu-aws
[4]: README-ASSET-COLLECTION.md