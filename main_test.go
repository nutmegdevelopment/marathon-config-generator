package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/nutmegdevelopment/marathon-golang-common"
)

var (
	baseFile    = "test-data/base-file.yml"
	overlayFile = "test-data/overlay-file.yml"
	outputFile  = "test-data/output.json"
)

func TestLoadingYAML(t *testing.T) {
	m := models.MarathonApp{}
	m.LoadYAML(readFile(baseFile))

	// Check fields contain what they should.
	// Verbose, but safer.
	assert.Equal(t, "sleep 300", m.Command)
	assert.Equal(t, "/product/service/myApp", m.ID)
	assert.Equal(t, len(m.Args), 3)
	assert.Equal(t, "sleep 300", m.Args[2])
	assert.Equal(t, 1.5, m.CPUs)
	assert.Equal(t, 256.0, m.Memory)
	assert.Equal(t, len(m.Ports), 2)
	assert.Equal(t, 9000, m.Ports[1])
	assert.Equal(t, false, m.RequirePorts)
	assert.Equal(t, 3, m.Instances)
	assert.Equal(t, "", m.Executor)
	assert.Equal(t, "DOCKER", m.Container.ContainerType)
	assert.Equal(t, "group/image", m.Container.Docker.Image)
	assert.Equal(t, "BRIDGE", m.Container.Docker.Network)
	assert.Equal(t, false, *m.Container.Docker.Privileged)
	assert.Equal(t, 2, len(m.Container.Docker.PortMappings))
	assert.Equal(t, 8080, m.Container.Docker.PortMappings[0].ContainerPort)
	assert.Equal(t, 0, *m.Container.Docker.PortMappings[0].HostPort)
	assert.Equal(t, 9000, m.Container.Docker.PortMappings[0].ServicePort)
	assert.Equal(t, "tcp", m.Container.Docker.PortMappings[0].Protocol)
	assert.Equal(t, false, *m.Container.Docker.Privileged)
	assert.Equal(t, 2, len(m.Container.Docker.Parameters))
	assert.Equal(t, "a-docker-option", m.Container.Docker.Parameters[0].Key)
	assert.Equal(t, "xxx", m.Container.Docker.Parameters[0].Value)
	assert.Equal(t, "b-docker-option", m.Container.Docker.Parameters[1].Key)
	assert.Equal(t, "yyy", m.Container.Docker.Parameters[1].Value)
	assert.Equal(t, 2, len(m.Container.Volumes))
	assert.Equal(t, "/etc/a", m.Container.Volumes[0].ContainerPath)
	assert.Equal(t, "/var/data/a", m.Container.Volumes[0].HostPath)
	assert.Equal(t, "RO", m.Container.Volumes[0].Mode)
	assert.Equal(t, "/etc/b", m.Container.Volumes[1].ContainerPath)
	assert.Equal(t, "/var/data/b", m.Container.Volumes[1].HostPath)
	assert.Equal(t, "RW", m.Container.Volumes[1].Mode)
	assert.Equal(t, "/usr/local/lib/myLib", m.Environment["LD_LIBRARY_PATH"])
	assert.Equal(t, "role1", m.AcceptedResourceRoles[0])
	assert.Equal(t, "*", m.AcceptedResourceRoles[1])
	assert.Equal(t, "staging", m.Labels["environment"])
	assert.Equal(t, "https://raw.github.com/mesosphere/marathon/master/README.md", m.URIs[0])
	assert.Equal(t, 3, len(m.Dependencies))
	assert.Equal(t, "/product/db/mongo", m.Dependencies[0])
	assert.Equal(t, "/product/db", m.Dependencies[1])
	assert.Equal(t, "../../db", m.Dependencies[2])
	assert.Equal(t, 3, len(m.HealthChecks))
	assert.Equal(t, "HTTP", m.HealthChecks[0].Protocol)
	assert.Equal(t, "/health", m.HealthChecks[0].Path)
	assert.Equal(t, 3, m.HealthChecks[0].GracePeriodSeconds)
	assert.Equal(t, 10, m.HealthChecks[0].IntervalSeconds)
	assert.Equal(t, 0, m.HealthChecks[0].PortIndex)
	assert.Equal(t, 10, m.HealthChecks[0].TimeoutSeconds)
	assert.Equal(t, 3, m.HealthChecks[0].MaxConsecutiveFailures)
	assert.Equal(t, "TCP", m.HealthChecks[1].Protocol)
	assert.Equal(t, 3, m.HealthChecks[1].GracePeriodSeconds)
	assert.Equal(t, 5, m.HealthChecks[1].IntervalSeconds)
	assert.Equal(t, 1, m.HealthChecks[1].PortIndex)
	assert.Equal(t, 5, m.HealthChecks[1].TimeoutSeconds)
	assert.Equal(t, 3, m.HealthChecks[1].MaxConsecutiveFailures)
	assert.Equal(t, "COMMAND", m.HealthChecks[2].Protocol)
	assert.Equal(t, "curl -f -X GET http://$HOST:$PORT0/health", m.HealthChecks[2].Command["value"])
	assert.Equal(t, 3, m.HealthChecks[2].MaxConsecutiveFailures)
	assert.Equal(t, 1, m.BackoffSeconds)
	assert.Equal(t, 1.15, m.BackoffFactor)
	assert.Equal(t, 3600, m.MaxLaunchDelaySeconds)
	assert.Equal(t, 0.2, *m.UpgradeStrategy.MaximumOverCapacity)
	assert.Equal(t, 0.5, *m.UpgradeStrategy.MinimumHealthCapacity)
}

func TestOverlayingYAML(t *testing.T) {
	m := models.MarathonApp{}
	m.LoadYAML(readFile(baseFile))
	m.LoadYAML(readFile(overlayFile))

	// Check fields contain what they should.
	// Verbose, but safer.
	assert.Equal(t, "sleep 300", m.Command)
	assert.Equal(t, "/product/service/myApp", m.ID)
	assert.Equal(t, len(m.Args), 3)
	assert.Equal(t, "sleep 300", m.Args[2])
	assert.Equal(t, 4.0, m.CPUs)
	assert.Equal(t, 666.0, m.Memory)
	assert.Equal(t, len(m.Ports), 2)
	assert.Equal(t, 9000, m.Ports[1])
	assert.Equal(t, false, m.RequirePorts)
	assert.Equal(t, 4, m.Instances)
	assert.Equal(t, "", m.Executor)
	assert.Equal(t, "DOCKER", m.Container.ContainerType)
	assert.Equal(t, "group/image", m.Container.Docker.Image)
	assert.Equal(t, "BRIDGE", m.Container.Docker.Network)
	assert.Equal(t, false, *m.Container.Docker.Privileged)
	assert.Equal(t, 2, len(m.Container.Docker.PortMappings))
	assert.Equal(t, 8080, m.Container.Docker.PortMappings[0].ContainerPort)
	assert.Equal(t, 0, *m.Container.Docker.PortMappings[0].HostPort)
	assert.Equal(t, 9000, m.Container.Docker.PortMappings[0].ServicePort)
	assert.Equal(t, "tcp", m.Container.Docker.PortMappings[0].Protocol)
	assert.Equal(t, false, *m.Container.Docker.Privileged)
	assert.Equal(t, 2, len(m.Container.Docker.Parameters))
	assert.Equal(t, "a-docker-option", m.Container.Docker.Parameters[0].Key)
	assert.Equal(t, "xxx", m.Container.Docker.Parameters[0].Value)
	assert.Equal(t, "b-docker-option", m.Container.Docker.Parameters[1].Key)
	assert.Equal(t, "yyy", m.Container.Docker.Parameters[1].Value)
	assert.Equal(t, 1, len(m.Container.Volumes))
	assert.Equal(t, "/etc/prod", m.Container.Volumes[0].ContainerPath)
	assert.Equal(t, "/var/data/prod", m.Container.Volumes[0].HostPath)
	assert.Equal(t, "RO", m.Container.Volumes[0].Mode)
	assert.Equal(t, "/usr/local/lib/myLib", m.Environment["LD_LIBRARY_PATH"])
	assert.Equal(t, "role1", m.AcceptedResourceRoles[0])
	assert.Equal(t, "*", m.AcceptedResourceRoles[1])
	assert.Equal(t, "prod", m.Labels["environment"])
	assert.Equal(t, "I am new", m.Labels["newlabel"])
	assert.Equal(t, "https://raw.github.com/mesosphere/marathon/master/README.md", m.URIs[0])
	assert.Equal(t, 3, len(m.Dependencies))
	assert.Equal(t, "/product/db/mongo", m.Dependencies[0])
	assert.Equal(t, "/product/db", m.Dependencies[1])
	assert.Equal(t, "../../db", m.Dependencies[2])
	assert.Equal(t, 1, len(m.HealthChecks))
	assert.Equal(t, "HTTP", m.HealthChecks[0].Protocol)
	assert.Equal(t, "/bealth", m.HealthChecks[0].Path)
	assert.Equal(t, 1500, m.HealthChecks[0].GracePeriodSeconds)
	assert.Equal(t, 15, m.HealthChecks[0].IntervalSeconds)
	assert.Equal(t, 0, m.HealthChecks[0].PortIndex)
	assert.Equal(t, 5, m.HealthChecks[0].TimeoutSeconds)
	assert.Equal(t, 10, m.HealthChecks[0].MaxConsecutiveFailures)
	assert.Equal(t, 1, m.BackoffSeconds)
	assert.Equal(t, 1.15, m.BackoffFactor)
	assert.Equal(t, 3600, m.MaxLaunchDelaySeconds)
	assert.Equal(t, 0.2, *m.UpgradeStrategy.MaximumOverCapacity)
	assert.Equal(t, 0.5, *m.UpgradeStrategy.MinimumHealthCapacity)
}

func TestToJSON(t *testing.T) {
	m := models.MarathonApp{}
	m.LoadYAML(readFile(baseFile))
	m.LoadYAML(readFile(overlayFile))
	o := readFile(outputFile)

	assert.Equal(t, o, string(m.ToJSON()))
}

func TestExampleOne(t *testing.T) {
	m := models.MarathonApp{}
	m.LoadYAML(DefaultYAMLDocker)
	m.LoadYAML(readFile("test-data/example-1.yml"))
	o := readFile("test-data/example-1.json")

	assert.Equal(t, o, string(m.ToJSON()))
}

func TestExampleTwo(t *testing.T) {
	m := models.MarathonApp{}
	m.LoadYAML(DefaultYAMLDocker)
	m.LoadYAML(readFile("test-data/example-2.yml"))
	o := readFile("test-data/example-2.json")

	assert.Equal(t, o, string(m.ToJSON()))
}

func TestExampleThree(t *testing.T) {
	m := models.MarathonApp{}
	m.LoadYAML(DefaultYAMLDocker)
	m.LoadYAML(readFile("test-data/example-3.yml"))
	o := readFile("test-data/example-3.json")

	assert.Equal(t, o, string(m.ToJSON()))
}
