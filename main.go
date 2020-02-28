package main // import "github.com/nutmegdevelopment/marathon-config-generator"

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	models "github.com/nutmegdevelopment/marathon-golang-common"
)

// DefaultYAMLDocker is...
const DefaultYAMLDocker = `
container:
  type: DOCKER
  docker:
    network: BRIDGE
instances: 1
env: 
  DOCKER_API_VERSION: "1.21"
  VAULT_SKIP_VERIFY: "true"
  VAULT_CA_PATH: "/tmp"
`

var (
	templateBuffer  bytes.Buffer
	environment     string
	verbose         bool
	baseConfig      string
	overlayConfig   string
	useDefaults     bool
	replacementVars = make(stringmap)
	configFiles     stringslice
)

func main() {
	flag.BoolVar(&verbose, "verbose", false, "verbose output.")
	flag.Var(&replacementVars, "var", "[] of replacement variables in the form of: key=value - multiple -var flags can be used, one per key/value pair.")
	flag.Var(&configFiles, "config-file", "[] of config files.")
	flag.BoolVar(&useDefaults, "use-defaults", true, "Use default settings for Docker containers")
	flag.Parse()

	if len(configFiles) < 1 {
		log.Println("Required parameter(s) missing: -config-file=...")
		log.Fatalln("Please use -h to see more information.")
	}

	if verbose {
		//log.Printf("Generating configuration for: %s", environment)
		log.Printf("Config files: %v", configFiles)
		log.Printf("Vars: %v", replacementVars)
	}

	// Create an instance of MarathonApp.
	t := models.MarathonApp{}

	// Apply the default Docker YAML.
	t.LoadYAML(DefaultYAMLDocker)

	// Apply the required config files.
	for _, f := range configFiles {
		if verbose {
			log.Printf("Applying config file: %s", f)
		}
		t.LoadYAML(readFile(f))
	}

	// Generate the JSON string.
	jsonString := t.ToJSON()

	// Let's see what we have created!
	fmt.Println(string(jsonString))
}

func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read the file: %s", filename)
	}
	before := string(b)
	after := before

	// Apply any replacement vars.
	for k, v := range replacementVars {
		if verbose {
			log.Printf("Replacing '%s' with '%s'", k, v)
		}
		k = fmt.Sprintf("{{%s}}", k)
		after = strings.Replace(after, k, v, -1)
	}
	// TODO:  Some sort of diff output
	return after
}
