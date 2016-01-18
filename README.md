#Marathon Config Generator

Generates a JSON string to represent a marathon app (and soon a group) from a
YAML definition file.

Allows multiple yaml files to be provided and applied (overlayed) in order so
that you have the ability to do the following:

1, Apply a full base configuration.
2, Apply an applications specific configuration over the top.
3, Apply an applications environment specific configuration over the top.

So if in prod you need 4 instances but the application specific config specifies
2 and the full base config specifies 1 then in your applications environment
specific config you set instances to 4.

A full example YAML file which shows all the options available through the
Marathon API can be found in the test-data/base-file.yml file.

## CLI

marathon-config-generator -h

```
Usage of /Users/james/Go/bin/marathon-config-generator:
  -config-file value
    	[] of config files. (default [])
  -var value
    	[] of replacement variables in the form of: key=value - multiple -var flags can be used, one per key/value pair. (default map[])
  -verbose
    	verbose output.
```

Example usage:

```
marathon-config-generator -var=GROUP_NAME=prod -var=CONSUL_IP=1.2.3.4 -config-file=test-data/base-file.yml -config-file=test-data/overlay-file.yml
```
