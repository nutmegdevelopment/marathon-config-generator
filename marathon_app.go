package main

import (
	"encoding/json"
	"log"

	"gopkg.in/yaml.v2"
)

/* Full JSON example containing ALL fields. From: https://mesosphere.github.io/marathon/docs/rest-api.html#post-v2-apps
{
    "id": "/product/service/myApp",
    "cmd": "env && sleep 300",
    "args": ["/bin/sh", "-c", "env && sleep 300"]
    "cpus": 1.5,
    "mem": 256.0,
    "ports": [
        8080,
        9000
    ],
    "portDefinitions": [
        { "port": 8080, "protocol": "tcp", "name": "http", labels: { "VIP_0": "10.0.0.1:80" } },
        { "port": 9000, "protocol": "tcp", "name": "admin" }
    ],
    "requirePorts": false,
    "instances": 3,
    "executor": "",
    "container": {
        "type": "DOCKER",
        "docker": {
            "image": "group/image",
            "network": "BRIDGE",
            "portMappings": [
                {
                    "containerPort": 8080,
                    "hostPort": 0,
                    "servicePort": 9000,
                    "protocol": "tcp"
                },
                {
                    "containerPort": 161,
                    "hostPort": 0,
                    "protocol": "udp"
                }
            ],
            "privileged": false,
            "parameters": [
                { "key": "a-docker-option", "value": "xxx" },
                { "key": "b-docker-option", "value": "yyy" }
            ]
        },
        "volumes": [
            {
                "containerPath": "/etc/a",
                "hostPath": "/var/data/a",
                "mode": "RO"
            },
            {
                "containerPath": "/etc/b",
                "hostPath": "/var/data/b",
                "mode": "RW"
            }
        ]
    },
    "env": {
        "LD_LIBRARY_PATH": "/usr/local/lib/myLib"
    },
    "constraints": [
        ["attribute", "OPERATOR", "value"]
    ],
    "acceptedResourceRoles": [
        "role1", "*"
    ],
    "labels": {
        "environment": "staging"
    },
    "uris": [
        "https://raw.github.com/mesosphere/marathon/master/README.md"
    ],
    "dependencies": ["/product/db/mongo", "/product/db", "../../db"],
    "healthChecks": [
        {
            "protocol": "HTTP",
            "path": "/health",
            "gracePeriodSeconds": 3,
            "intervalSeconds": 10,
            "portIndex": 0,
            "timeoutSeconds": 10,
            "maxConsecutiveFailures": 3
        },
        {
            "protocol": "TCP",
            "gracePeriodSeconds": 3,
            "intervalSeconds": 5,
            "portIndex": 1,
            "timeoutSeconds": 5,
            "maxConsecutiveFailures": 3
        },
        {
            "protocol": "COMMAND",
            "command": {
              "value": "curl -f -X GET http://$HOST:$PORT0/health"
            },
            "maxConsecutiveFailures": 3
        }
    ],
    "backoffSeconds": 1,
    "backoffFactor": 1.15,
    "maxLaunchDelaySeconds": 3600,
    "upgradeStrategy": {
        "minimumHealthCapacity": 0.5,
        "maximumOverCapacity": 0.2
    }
}
*/

// MarathonApp represents a Marathon application configuration.
// It inputs from yaml.
// It outputs to JSON.
type MarathonApp struct {
	AcceptedResourceRoles []string   `yaml:"acceptedResourceRoles" json:"acceptedResourceRoles,omitempty"`
	Args                  []string   `yaml:"args" json:"args,omitempty"`
	BackoffFactor         float64    `yaml:"backoffFactor" json:"backoffFactor,omitempty"`
	BackoffSeconds        int        `yaml:"backoffSeconds" json:"backoffSeconds,omitempty"`
	Command               string     `yaml:"cmd" json:"cmd,omitempty"`
	Constraints           [][]string `yaml:"constraints" json:"constraints,omitempty"`
	Container             struct {
		ContainerType string `yaml:"type" json:"type"`
		Docker        struct {
			Image        string `yaml:"image" json:"image"`
			Network      string `yaml:"network" json:"network,omitempty"`
			PortMappings []struct {
				ContainerPort int    `yaml:"containerPort" json:"containerPort,omitempty"`
				HostPort      int    `yaml:"hostPort" json:"hostPort"`
				Protocol      string `yaml:"protocol" json:"protocol,omitempty"`
				ServicePort   int    `yaml:"servicePort" json:"servicePort,omitempty"`
			} `yaml:"portMappings" json:"portMappings"`
			Privileged bool `yaml:"privileged" json:"privileged"`
			Parameters []struct {
				Key   string `yaml:"key" json:"key,omitempty"`
				Value string `yaml:"value" json:"value,omitempty"`
			} `yaml:"parameters" json:"parameters"`
		} `yaml:"docker" json:"docker"`
		Volumes []struct {
			ContainerPath string `yaml:"containerPath" json:"containerPath"`
			HostPath      string `yaml:"hostPath" json:"hostPath"`
			Mode          string `yaml:"mode" json:"mode"`
		} `yaml:"volumes" json:"volumes"`
	} `yaml:"container" json:"container"`
	CPUs         float64           `yaml:"cpus" json:"cpus,omitempty"`
	Dependencies []string          `yaml:"dependencies" json:"dependencies,omitempty"`
	Environment  map[string]string `yaml:"env" json:"env,omitempty"`
	Executor     string            `yaml:"executor" json:"executor,omitempty"`
	HealthChecks []struct {
		Command                map[string]string `yaml:"command" json:"command,omitempty"`
		GracePeriodSeconds     int               `yaml:"gracePeriodSeconds" json:"gracePeriodSeconds,omitempty"`
		IntervalSeconds        int               `yaml:"intervalSeconds" json:"intervalSeconds,omitempty"`
		MaxConsecutiveFailures int               `yaml:"maxConsecutiveFailures" json:"maxConsecutiveFailures,omitempty"`
		Path                   string            `yaml:"path" json:"path,omitempty"`
		PortIndex              int               `yaml:"portIndex" json:"portIndex,omitempty"`
		Protocol               string            `yaml:"protocol" json:"protocol,omitempty"`
		TimeoutSeconds         int               `yaml:"timeoutSeconds" json:"timeoutSeconds,omitempty"`
		Port                   int               `yaml:"port" json:"port,omitempty"`
	} `yaml:"healthChecks" json:"healthChecks,omitempty"`
	ID                    string            `yaml:"id" json:"id"`
	Instances             int               `yaml:"instances" json:"instances,omitempty"`
	Labels                map[string]string `yaml:"labels" json:"labels,omitempty"`
	MaxLaunchDelaySeconds int               `yaml:"maxLaunchDelaySeconds" json:"maxLaunchDelaySeconds,omitempty"`
	Memory                int               `yaml:"mem" json:"mem,omitempty"`
	Ports                 []int             `yaml:"ports" json:"ports,omitempty"`
	PortDefinitions       []struct {
		Port     string `yaml:"port" json:"port,omitempty"`
		Protocol string `yaml:"protocol" json:"protocol,omitempty"`
		Name     string `yaml:"name" json:"name,omitempty"`
	} `yaml:"portDefinitions" json:"portDefinitions,omitempty"`
	RequirePorts    bool `yaml:"requirePorts,omitempty" json:"requirePorts"`
	UpgradeStrategy struct {
		MaximumOverCapacity   *float64 `yaml:"maximumOverCapacity" json:"maximumOverCapacity,omitempty"`
		MinimumHealthCapacity *float64 `yaml:"minimumHealthCapacity" json:"minimumHealthCapacity,omitempty"`
	} `yaml:"upgradeStrategy" json:"upgradeStrategy,omitempty"`
	URIs []string `yaml:"uris" json:"uris,omitempty"`
}

// LoadYAML takes a YAML string and unmarshalls it against itself.
// This can be applied multiple times with different YAML file to, for example,
// load a base configuration and then load a subset of the YAML to override prod
// configuration.
func (t *MarathonApp) LoadYAML(yamlString string) {
	err := yaml.Unmarshal([]byte(yamlString), t)
	if err != nil {
		log.Fatalf("Error parsing YAML: %s", err.Error())
	}
}

// ToJSON returns a JSON string representation of itself.
func (t *MarathonApp) ToJSON() string {
	jsonString, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err.Error())
	}
	return string(jsonString[:])
}
