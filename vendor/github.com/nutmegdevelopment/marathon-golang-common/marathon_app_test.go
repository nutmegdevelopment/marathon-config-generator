package models

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	baseFile    = "test-data/base-file.yml"
	overlayFile = "test-data/overlay-file.yml"
	outputFile  = "test-data/output.json"

	// JSON string outputs.
	MarathonAppUpgradeStrategyTestJSON             = `{"maximumOverCapacity":0,"minimumHealthCapacity":1}`
	MarathonAppPortDefinitionTestJSON              = `{"labels":{"label-name":"label-value"},"name":"port-name","port":8080,"protocol":"HTTP"}`
	MarathonAppHealthCheckTestJSON                 = `{"command":{"value":"i am the command value"},"gracePeriodSeconds":15,"intervalSeconds":10,"maxConsecutiveFailures":3,"path":"/health","port":8080,"portIndex":1,"protocol":"TCP","timeoutSeconds":360}`
	MarathonAppIPAddressTestJSON                   = `{"groups":["group0","groups1"],"labels":{"label-name":"label-value"},"networkName":"network-name"}`
	MarathonAppFetchFalseTestJSON                  = `{"uri":"http://b3ta.com","extract":false,"executable":false,"cache":false}`
	MarathonAppFetchTrueTestJSON                   = `{"uri":"http://b3ta.com","extract":true,"executable":true,"cache":true}`
	MarathonAppContainerVolumeTestJSON             = `{"containerPath":"container-path","hostPath":"host-path","mode":"mode"}`
	MarathonAppContainerDockerPortMappingsTestJSON = `{"containerPort":8080,"hostPort":31001,"protocol":"HTTP","servicePort":15001}`
	MarathonAppContainerDockerTestJSON             = `{"image":"registry.nutmeg.co.uk:8443/...","network":"BRIDGE","portMappings":[{"containerPort":8080,"hostPort":31001,"protocol":"HTTP","servicePort":15001},{}],"privileged":false,"parameters":[{"key":"key1","value":"value1"},{"key":"key2","value":"value2"}]}`
	MarathonAppContainerTestJSON                   = `{"type":"DOCKER","forcePullImage":true}`
)

func TestMinimumOutput(t *testing.T) {
	m := MarathonApp{}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	assert.Equal(t, "{}", string(b))
}

func TestMarathonAppUpgradeStrategy(t *testing.T) {
	m := MarathonAppUpgradeStrategy{}
	min := 1.0
	max := 0.0
	m.MaximumOverCapacity = &max
	m.MinimumHealthCapacity = &min
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppUpgradeStrategyTestJSON, string(b))
}

func TestMarathonAppPortDefinitionEmpty(t *testing.T) {
	m := MarathonAppPortDefinition{}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppPortDefinitionFull(t *testing.T) {
	m := MarathonAppPortDefinition{}
	labels := make(map[string]string)
	labels["label-name"] = "label-value"
	m.Labels = labels
	m.Name = "port-name"
	m.Port = 8080
	m.Protocol = "HTTP"
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppPortDefinitionTestJSON, string(b))
}

func TestMarathonAppHealthCheckEmpty(t *testing.T) {
	m := MarathonAppHealthCheck{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppHealthCheckFull(t *testing.T) {
	m := MarathonAppHealthCheck{}
	command := make(map[string]string)
	command["value"] = "i am the command value"
	m.Command = command
	m.GracePeriodSeconds = 15
	m.IntervalSeconds = 10
	m.MaxConsecutiveFailures = 3
	m.Path = "/health"
	m.Port = 8080
	m.PortIndex = 1
	m.Protocol = "TCP"
	m.TimeoutSeconds = 360

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppHealthCheckTestJSON, string(b))
}

func TestMarathonAppIPAddressEmpty(t *testing.T) {
	m := MarathonAppIPAddress{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppIPAddressFull(t *testing.T) {
	m := MarathonAppIPAddress{}
	groups := make([]string, 2)
	groups[0] = "group0"
	groups[1] = "groups1"
	m.Groups = groups
	labels := make(map[string]string)
	labels["label-name"] = "label-value"
	m.Labels = labels
	m.NetworkName = "network-name"

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppIPAddressTestJSON, string(b))
}

func TestMarathonAppFetchEmpty(t *testing.T) {
	m := MarathonAppFetch{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppFetchAllFalse(t *testing.T) {
	m := MarathonAppFetch{}
	bo := false
	m.Cache = &bo
	m.Executable = &bo
	m.Extract = &bo
	m.URI = "http://b3ta.com"

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppFetchFalseTestJSON, string(b))
}

func TestMarathonAppFetchAllTrue(t *testing.T) {
	m := MarathonAppFetch{}
	bo := true
	m.Cache = &bo
	m.Executable = &bo
	m.Extract = &bo
	m.URI = "http://b3ta.com"

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppFetchTrueTestJSON, string(b))
}

func TestMarathonAppContainerVolumeEmpty(t *testing.T) {
	m := MarathonAppContainerVolume{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppContainerVolumeFull(t *testing.T) {
	m := MarathonAppContainerVolume{}
	m.ContainerPath = "container-path"
	m.HostPath = "host-path"
	m.Mode = "mode"

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppContainerVolumeTestJSON, string(b))
}

func TestMarathonAppContainerDockerParametersEmpty(t *testing.T) {
	m := MarathonAppContainerDockerParameters{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppContainerDockerParametersFull(t *testing.T) {
	m := MarathonAppContainerDockerParameters{}
	m.Key = "key"
	m.Value = "value"

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{"key":"key","value":"value"}`, string(b))
}

func TestMarathonAppContainerDockerPortMappingsEmpty(t *testing.T) {
	m := MarathonAppContainerDockerPortMappings{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppContainerDockerPortMappingsFull(t *testing.T) {
	m := MarathonAppContainerDockerPortMappings{}
	m.ContainerPort = 8080
	hostPort := 31001
	m.HostPort = &hostPort
	m.Protocol = "HTTP"
	m.ServicePort = 15001

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppContainerDockerPortMappingsTestJSON, string(b))
}

func TestMarathonAppContainerDockerEmpty(t *testing.T) {
	m := MarathonAppContainerDocker{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppContainerDockerFull(t *testing.T) {
	m := MarathonAppContainerDocker{}
	m.Image = "registry.nutmeg.co.uk:8443/..."
	m.Network = "BRIDGE"
	macdp := make([]MarathonAppContainerDockerParameters, 2)
	macdp[0].Key = "key1"
	macdp[0].Value = "value1"
	macdp[1].Key = "key2"
	macdp[1].Value = "value2"
	m.Parameters = macdp
	macdpm := make([]MarathonAppContainerDockerPortMappings, 2)
	macdpm[0].ContainerPort = 8080
	hostPort := 31001
	macdpm[0].HostPort = &hostPort
	macdpm[0].Protocol = "HTTP"
	macdpm[0].ServicePort = 15001
	m.PortMappings = macdpm
	priv := false
	m.Privileged = &priv

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppContainerDockerTestJSON, string(b))
}

func TestMarathonAppContainerEmpty(t *testing.T) {
	m := MarathonAppContainer{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}

func TestMarathonAppContainerFull(t *testing.T) {
	m := MarathonAppContainer{}
	m.ContainerType = "DOCKER"
	m.ForcePullImage = true

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, MarathonAppContainerTestJSON, string(b))
}

func TestMarathonAppEmpty(t *testing.T) {
	m := MarathonApp{}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	assert.Equal(t, `{}`, string(b))
}
