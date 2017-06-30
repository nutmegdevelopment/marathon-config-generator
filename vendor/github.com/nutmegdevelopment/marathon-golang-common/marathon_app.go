package models

import (
	"encoding/json"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// MarathonApp represents the bulk of the response from /v2/apps/<APP_ID>
// This is the JSON that is posted to the destination Marathon.
type MarathonApp struct {
	AcceptedResourceRoles      []string                    `yaml:"acceptedResourceRoles" json:"acceptedResourceRoles,omitempty"`
	Args                       []string                    `yaml:"args" json:"args,omitempty"`
	BackoffFactor              float64                     `yaml:"backoffFactor" json:"backoffFactor,omitempty"`
	BackoffSeconds             int                         `yaml:"backoffSeconds" json:"backoffSeconds,omitempty"`
	Command                    string                      `yaml:"cmd" json:"cmd,omitempty"`
	Constraints                [][]string                  `yaml:"constraints" json:"constraints,omitempty"`
	Container                  *MarathonAppContainer       `yaml:"container" json:"container,omitempty"`
	CPUs                       float64                     `yaml:"cpus" json:"cpus,omitempty"`
	Dependencies               []string                    `yaml:"dependencies" json:"dependencies,omitempty"`
	Environment                map[string]string           `yaml:"env" json:"env,omitempty"`
	Executor                   string                      `yaml:"executor" json:"executor,omitempty"`
	Fetch                      []MarathonAppFetch          `yaml:"fetch" json:"fetch,omitempty"`
	HealthChecks               []MarathonAppHealthCheck    `yaml:"healthChecks" json:"healthChecks,omitempty"`
	ID                         string                      `yaml:"id" json:"id,omitempty"`
	IPAddress                  *MarathonAppIPAddress       `json:"ipAddress,omitempty"`
	OriginalID                 string                      `json:"-"`
	Instances                  int                         `yaml:"instances" json:"instances,omitempty"`
	Labels                     map[string]string           `yaml:"labels" json:"labels,omitempty"`
	MaxLaunchDelaySeconds      int                         `yaml:"maxLaunchDelaySeconds" json:"maxLaunchDelaySeconds,omitempty"`
	Memory                     float64                     `yaml:"mem" json:"mem,omitempty"`
	Ports                      []int                       `yaml:"ports" json:"ports,omitempty"`
	PortDefinitions            []MarathonAppPortDefinition `yaml:"portDefinitions" json:"portDefinitions,omitempty"`
	RequirePorts               bool                        `yaml:"requirePorts,omitempty" json:"requirePorts,omitempty"`
	TaskKillGracePeriodSeconds int                         `yaml:"taskKillGracePeriodSeconds" json:"taskKillGracePeriodSeconds,omitempty"`
	UpgradeStrategy            *MarathonAppUpgradeStrategy `yaml:"upgradeStrategy" json:"upgradeStrategy,omitempty"`
	URIs                       []string                    `yaml:"uris" json:"uris,omitempty"`
}

type MarathonAppContainer struct {
	ContainerType  string                       `yaml:"type" json:"type,omitempty"`
	Docker         *MarathonAppContainerDocker  `yaml:"docker" json:"docker,omitempty"`
	ForcePullImage bool                         `yaml:"forcePullImage" json:"forcePullImage,omitempty"`
	Volumes        []MarathonAppContainerVolume `yaml:"volumes" json:"volumes,omitempty"`
}

type MarathonAppContainerDocker struct {
	Image        string                                   `yaml:"image" json:"image,omitempty"`
	Network      string                                   `yaml:"network" json:"network,omitempty"`
	PortMappings []MarathonAppContainerDockerPortMappings `yaml:"portMappings" json:"portMappings,omitempty"`
	Privileged   *bool                                    `yaml:"privileged" json:"privileged,omitempty"`
	Parameters   []MarathonAppContainerDockerParameters   `yaml:"parameters" json:"parameters,omitempty"`
}

type MarathonAppContainerDockerPortMappings struct {
	ContainerPort int    `yaml:"containerPort" json:"containerPort,omitempty"`
	HostPort      *int   `yaml:"hostPort" json:"hostPort,omitempty"`
	Protocol      string `yaml:"protocol" json:"protocol,omitempty"`
	ServicePort   int    `yaml:"servicePort" json:"servicePort,omitempty"`
}

type MarathonAppContainerDockerParameters struct {
	Key   string `yaml:"key" json:"key,omitempty"`
	Value string `yaml:"value" json:"value,omitempty"`
}

type MarathonAppContainerVolume struct {
	ContainerPath string `yaml:"containerPath" json:"containerPath,omitempty"`
	HostPath      string `yaml:"hostPath" json:"hostPath,omitempty"`
	Mode          string `yaml:"mode" json:"mode,omitempty"`
}

type MarathonAppFetch struct {
	URI        string `yaml:"uri" json:"uri,omitempty"`
	Extract    *bool  `yaml:"extract" json:"extract,omitempty"`
	Executable *bool  `yaml:"executable" json:"executable,omitempty"`
	Cache      *bool  `yaml:"cache" json:"cache,omitempty"`
}

type MarathonAppHealthCheck struct {
	Command                map[string]string `yaml:"command" json:"command,omitempty"`
	GracePeriodSeconds     int               `yaml:"gracePeriodSeconds" json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        int               `yaml:"intervalSeconds" json:"intervalSeconds,omitempty"`
	MaxConsecutiveFailures int               `yaml:"maxConsecutiveFailures" json:"maxConsecutiveFailures,omitempty"`
	Path                   string            `yaml:"path" json:"path,omitempty"`
	Port                   int               `yaml:"port" json:"port,omitempty"`
	PortIndex              int               `yaml:"portIndex" json:"portIndex,omitempty"`
	Protocol               string            `yaml:"protocol" json:"protocol,omitempty"`
	TimeoutSeconds         int               `yaml:"timeoutSeconds" json:"timeoutSeconds,omitempty"`
}

type MarathonAppIPAddress struct {
	Groups      []string          `yaml:"groups" json:"groups,omitempty"`
	Labels      map[string]string `yaml:"labels" json:"labels,omitempty"`
	NetworkName string            `yaml:"networkName" json:"networkName,omitempty"`
}

type MarathonAppUpgradeStrategy struct {
	MaximumOverCapacity   *float64 `yaml:"maximumOverCapacity" json:"maximumOverCapacity,omitempty"`
	MinimumHealthCapacity *float64 `yaml:"minimumHealthCapacity" json:"minimumHealthCapacity,omitempty"`
}

type MarathonAppPortDefinition struct {
	Labels   map[string]string `yaml:"labels" json:"labels,omitempty"`
	Name     string            `yaml:"name" json:"name,omitempty"`
	Port     int               `yaml:"port" json:"port,omitempty"`
	Protocol string            `yaml:"protocol" json:"protocol,omitempty"`
}

// ToJSON returns a JSON string representation of itself.
func (t *MarathonApp) ToJSON() []byte {
	jsonString, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err.Error())
	}
	return jsonString
}

// LoadYAML loads a YAML file and marshals it into a MarathonApp type.
func (t *MarathonApp) LoadYAML(yamlString string) {
	err := yaml.Unmarshal([]byte(yamlString), t)
	if err != nil {
		log.Fatalf("Error parsing YAML: %s", err.Error())
	}
}
