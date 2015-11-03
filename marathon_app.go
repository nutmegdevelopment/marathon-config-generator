package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

//  "gopkg.in/yaml.v2"

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

/*
  "id": "file-interface",
  "cpus": 0.1,
  "mem": 128,
  "instances": 1,
  "args": ["/src/file-interface", "-log.file=/dev/stdout", "-log.level=info", "-config.file=/src/dev/config.yml"],
  "env": {
    "AWS_ACCESS_KEY_ID": "TESTKEY",
    "AWS_SECRET_ACCESS_KEY": "TESTSECRET"
  },
  "container": {
    "type": "DOCKER",
    "docker": {
      "image": "zk.nmtest.int:10006/file-interface:VERSION",
      "network": "BRIDGE",
      "portMappings": [
        { "containerPort": 80 }
      ]
    }
  },
  "healthChecks": [
    {
      "protocol": "HTTP",
      "path": "/health",
      "gracePeriodSeconds": 1500,
      "intervalSeconds": 15,
      "timeoutSeconds": 5,
      "maxConsecutiveFailures": 10
    }
  ],
  "labels": {
    "environment": "development",
    "HAPROXY_HTTP": "true",
    "HTTP_PORT_IDX_0_NAME": "file-interface"
  }
*/

/*
"portMappings": [
    {
        "containerPort": 8080,
        "hostPort": 0,
        "servicePort": 9000,
        "protocol": "tcp"
    },
    {
       ...
    }
],
*/

// MarathonAppPortMappings represents the port mappings section of the yaml file.
type MarathonAppPortMappings struct {
	PortMapping MarathonAppPortMapping `yaml:"portMapping"`
}

// MarathonAppPortMapping represents the port mapping sections of the yaml file.
type MarathonAppPortMapping struct {
	ContainerPort int    `yaml:"containerPort" json:"containerPort"`
	HostPort      int    `yaml:"hostPort"`
	ServicePort   int    `yaml:"servicePort"`
	Protocol      string `yaml:"protocol"`
}

// MarshalJSON handles the correct output format for port mappings.
// Handles any number of keys existing.
func (t MarathonAppPortMappings) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	if t.PortMapping.ContainerPort != 0 {
		AppendJSON(&buffer, fmt.Sprintf(`"containerPort":%d`, t.PortMapping.ContainerPort))
	}
	if t.PortMapping.HostPort != 0 {
		AppendJSON(&buffer, fmt.Sprintf(`"hostPort":%d`, t.PortMapping.HostPort))
	}
	if t.PortMapping.ServicePort != 0 {
		AppendJSON(&buffer, fmt.Sprintf(`"servicePort":%d`, t.PortMapping.ServicePort))
	}
	if t.PortMapping.Protocol != "" {
		AppendJSON(&buffer, fmt.Sprintf(`"protocol":"%s"`, t.PortMapping.Protocol))
	}

	buffer.WriteString("}")
	return buffer.Bytes(), nil
	//jsonTemplate := `{"containerPort":%d,"hostPort":%d,"servicePort":%d,"protocol":"%s"}`
	//return []byte(fmt.Sprintf(jsonTemplate, t.PortMapping.ContainerPort, t.PortMapping.HostPort, t.PortMapping.ServicePort, t.PortMapping.Protocol)), nil
}

// MarathonAppVolumes represents the volumes section of the yaml file.
type MarathonAppVolumes struct {
	Volume MarathonAppVolume `yaml:"volume"`
}

// MarathonAppVolume represents the volume sections of the yaml file.
type MarathonAppVolume struct {
	ContainerPath string `yaml:"containerPath" json:"containerPath"`
	HostPath      string `yaml:"hostPath" json:"hostPath"`
	Mode          string `yaml:"mode" json:"mode"`
}

// MarshalJSON handles the correct output format for volumes.
func (t MarathonAppVolumes) MarshalJSON() ([]byte, error) {
	jsonTemplate := `{"containerPath":"%s","hostPath":"%s","mode":"%s"}`
	return []byte(fmt.Sprintf(jsonTemplate, t.Volume.ContainerPath, t.Volume.HostPath, t.Volume.Mode)), nil
}

// MarathonAppHealthChecks represents the healthChecks section of the yaml file.
type MarathonAppHealthChecks struct {
	HealthCheck MarathonAppHealthCheck `yaml:"healthCheck"`
}

// MarathonAppHealthCheck represents the healthCheck sections of the yaml file.
type MarathonAppHealthCheck struct {
	PortIndex              int    `yaml:"portIndex" json:"portIndex"`
	Protocol               string `yaml:"protocol" json:"protocol"`
	Path                   string `yaml:"path" json:"path"`
	GracePeriodSeconds     int    `yaml:"gracePeriodSeconds" json:"gracePeriodSeconds"`
	IntervalSeconds        int    `yaml:"intervalSeconds" json:"intervalSeconds"`
	TimeoutSeconds         int    `yaml:"timeoutSeconds" json:"timeoutSeconds"`
	MaxConsecutiveFailures int    `yaml:"maxConsecutiveFailures" json:"maxConsecutiveFailures"`
}

// MarshalJSON handles the correct output format for healthChecks.
func (t MarathonAppHealthChecks) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	if t.HealthCheck.Protocol != "" {
		AppendJSON(&buffer, fmt.Sprintf(`"protocol":"%s"`, t.HealthCheck.Protocol))
	}
	if t.HealthCheck.Path != "" {
		AppendJSON(&buffer, fmt.Sprintf(`"path":"%s"`, t.HealthCheck.Path))
	}
	if t.HealthCheck.GracePeriodSeconds != 0 {
		AppendJSON(&buffer, fmt.Sprintf(`"gracePeriodSeconds":%d`, t.HealthCheck.GracePeriodSeconds))
	}
	if t.HealthCheck.IntervalSeconds != 0 {
		AppendJSON(&buffer, fmt.Sprintf(`"intervalSeconds":%d`, t.HealthCheck.IntervalSeconds))
	}
	if t.HealthCheck.TimeoutSeconds != 0 {
		AppendJSON(&buffer, fmt.Sprintf(`"timeoutSeconds":%d`, t.HealthCheck.TimeoutSeconds))
	}
	if t.HealthCheck.MaxConsecutiveFailures != 0 {
		AppendJSON(&buffer, fmt.Sprintf(`"maxConsecutiveFailures":%d`, t.HealthCheck.MaxConsecutiveFailures))
	}
	if t.HealthCheck.PortIndex != 0 {
		AppendJSON(&buffer, fmt.Sprintf(`"portIndex":%d`, t.HealthCheck.PortIndex))
	}

	buffer.WriteString("}")
	return buffer.Bytes(), nil
	//jsonTemplate := `{"protocol":"%s","path":"%s","gracePeriodSeconds":%d,"intervalSeconds":%d,"timeoutSeconds":%d,"maxConsecutiveFailures":%d,"portIndex":%d}`
	//return []byte(fmt.Sprintf(jsonTemplate, t.HealthCheck.Protocol, t.HealthCheck.Path, t.HealthCheck.GracePeriodSeconds, t.HealthCheck.IntervalSeconds, t.HealthCheck.TimeoutSeconds, t.HealthCheck.MaxConsecutiveFailures, t.HealthCheck.PortIndex)), nil
}

// AppendJSON is used to append a string to a string with a JSON separator ","
// between items.
func AppendJSON(b *bytes.Buffer, s string) {
	if b.Len() > 1 {
		b.WriteString(",")
	}
	b.WriteString(s)
}

// MarathonApp represents a Marathon application configuration.
// It inputs from yaml.
// It outputs to JSON.
type MarathonApp struct {
	ID        string   `yaml:"id" json:"id"`
	CPUs      float32  `yaml:"cpus" json:"cpus"`
	Memory    int      `yaml:"mem" json:"mem"`
	Instances int      `yaml:"instances" json:"instances"`
	Args      []string `yaml:"args" json:"args"`
	Container struct {
		ContainerType string `yaml:"type" json:"type"`
		Docker        struct {
			Image        string                    `yaml:"image" json:"image"`
			Network      string                    `yaml:"network" json:"network"`
			PortMappings []MarathonAppPortMappings `yaml:"portMappings" json:"portMappings"`
		} `yaml:"docker" json:"docker"`
		Volumes []MarathonAppVolumes `yaml:"volumes" json:"volumes"`
	} `yaml:"container" json:"container"`
	HealthChecks    []MarathonAppHealthChecks `yaml:"healthChecks" json:"healthChecks"`
	Labels          map[string]string         `yaml:"labels" json:"labels"`
	Ports           []int                     `yaml:"ports" json:"ports"`
	UpgradeStrategy struct {
		MinimumHealthCapacity float32 `yaml:"minimumHealthCapacity" json:"minimumHealthCapacity"`
		MaximumOverCapacity   float32 `yaml:"maximumOverCapacity" json:"maximumOverCapacity"`
	} `yaml:"upgradeStrategy" json:"upgradeStrategy"`
}

// ParseYAML takes a YAML string and unmarshalls it against itself.
// This can be applied multiple times with different YAML file to, for example,
// load a base configuration and then load a subset of the YAML to override prod
// configuration.
func (t *MarathonApp) ParseYAML(yamlString string) {
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

/*
type MarathonApp struct {
	ID        string   `yaml:"id" json:"id"`
	CPUs      float32  `yaml:"cpus" json:"cpus"`
	Memory    int      `yaml:"mem" json:"mem"`
	Instances int      `yaml:"instances" json:"instances"`
	Args      []string `yaml:"args" json:"args"`
	Container struct {
		ContainerType string `yaml:"type" json:"type"`
		Docker        struct {
			Image        string `yaml:"image" json:"image"`
			Network      string `yaml:"network" json:"network"`
			PortMappings []map[string]int
		} `yaml:"docker" json:"docker"`
		Volumes []MarathonAppVolumes `yaml:"volumes" json:"volumes"`
		// Volumes []struct {
		// 	Volume struct {
		// 		ContainerPath string `yaml:"containerPath" json:"containerPath"`
		// 		HostPath      string `yaml:"hostPath" json:"hostPath"`
		// 		Mode          string `yaml:"mode" json:"mode"`
		// 	} `yaml:"volume" json:"volume"`
		// } `yaml:"volumes" json:"volumes"`
	} `yaml:"container" json:"container"`
	HealthChecks []MarathonAppHealthChecks `yaml:"healthChecks" json:"healthChecks"`
	// HealthChecks []struct {
	// 	HealthCheck struct {
	// 		Protocol               string `yaml:"protocol" json:"protocol"`
	// 		Path                   string `yaml:"path" json:"path"`
	// 		GracePeriodSeconds     int    `yaml:"gracePeriodSeconds" json:"gracePeriodSeconds"`
	// 		IntervalSeconds        int    `yaml:"intervalSeconds" json:"intervalSeconds"`
	// 		TimeoutSeconds         int    `yaml:"timeoutSeconds" json:"timeoutSeconds"`
	// 		MaxConsecutiveFailures int    `yaml:"maxConsecutiveFailures" json:"maxConsecutiveFailures"`
	// 	} `yaml:"healthCheck" json:"healthCheck"`
	// } `yaml:"healthChecks" json:"healthChecks"`
	Ports           []int `yaml:"ports" json:"ports"`
	UpgradeStrategy struct {
		MinimumHealthCapacity float32 `yaml:"minimumHealthCapacity" json:"minimumHealthCapacity"`
		MaximumOverCapacity   float32 `yaml:"maximumOverCapacity" json:"maximumOverCapacity"`
	} `yaml:"upgradeStrategy" json:"upgradeStrategy"`
}
*/
