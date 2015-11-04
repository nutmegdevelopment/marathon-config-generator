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
            "command": { "value": "curl -f -X GET http://$HOST:$PORT0/health" },
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
	Command      string   `yaml:"cmd" json:"cmd,omitempty"`
	RequirePorts bool     `yaml:"requirePorts,omitempty"`
	Executor     string   `yaml:"executor,omitempty"`
	ID           string   `yaml:"id" json:"id"`
	CPUs         float64  `yaml:"cpus" json:"cpus"`
	Memory       int      `yaml:"mem" json:"mem"`
	Instances    int      `yaml:"instances" json:"instances"`
	Args         []string `yaml:"args" json:"args"`
	Container    struct {
		ContainerType string `yaml:"type" json:"type"`
		Docker        struct {
			Image        string `yaml:"image" json:"image"`
			Network      string `yaml:"network" json:"network,omitempty"`
			PortMappings []struct {
				ContainerPort int    `yaml:"containerPort" json:"containerPort,omitempty"`
				HostPort      int    `yaml:"hostPort" json:"hostPort,omitempty"`
				ServicePort   int    `yaml:"servicePort" json:"servicePort,omitempty"`
				Protocol      string `yaml:"protocol" json:"protocol,omitempty"`
			} `yaml:"portMappings" json:"portMappings"`
		} `yaml:"docker" json:"docker"`
		Volumes []struct {
			ContainerPath string `yaml:"containerPath" json:"containerPath"`
			HostPath      string `yaml:"hostPath" json:"hostPath"`
			Mode          string `yaml:"mode" json:"mode"`
		} `yaml:"volumes" json:"volumes"`
	} `yaml:"container" json:"container"`
	HealthChecks []struct {
		PortIndex              int    `yaml:"portIndex" json:"portIndex,omitempty"`
		Protocol               string `yaml:"protocol" json:"protocol,omitempty"`
		Path                   string `yaml:"path" json:"path,omitempty"`
		GracePeriodSeconds     int    `yaml:"gracePeriodSeconds" json:"gracePeriodSeconds,omitempty"`
		IntervalSeconds        int    `yaml:"intervalSeconds" json:"intervalSeconds,omitempty"`
		TimeoutSeconds         int    `yaml:"timeoutSeconds" json:"timeoutSeconds,omitempty"`
		MaxConsecutiveFailures int    `yaml:"maxConsecutiveFailures" json:"maxConsecutiveFailures,omitempty"`
	} `yaml:"healthChecks" json:"healthChecks,omitempty"`
	Labels                map[string]string `yaml:"labels" json:"labels,omitempty"`
	Ports                 []int             `yaml:"ports" json:"ports,omitempty"`
	BackoffSeconds        int               `yaml:"backoffSeconds" json:"backOffSeconds,omitempty"`
	BackoffFactor         float64           `yaml:"backoffFactor" json:"backoffFactor,omitempty"`
	MaxLaunchDelaySeconds int               `yaml:"maxLaunchDelaySeconds" json:"maxLaunchDelaySeconds,omitempty"`
	UpgradeStrategy       struct {
		MinimumHealthCapacity float64 `yaml:"minimumHealthCapacity" json:"minimumHealthCapacity,omitempty"`
		MaximumOverCapacity   float64 `yaml:"maximumOverCapacity" json:"maximumOverCapacity,omitempty"`
	} `yaml:"upgradeStrategy" json:"upgradeStrategy,omitempty"`
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
	jsonString, err := json.Marshal(t)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err.Error())
	}
	return string(jsonString[:])
}
