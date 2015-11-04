package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	templateBuffer bytes.Buffer
	environment    string
	verbose        bool
	baseConfig     string
	overlayConfig  string
)

func main() {
	flag.StringVar(&environment, "env", "test", "the environment configuration to be generated.")
	flag.StringVar(&baseConfig, "base-config", "", "the base yaml configuration.")
	flag.StringVar(&overlayConfig, "overlay-config", "", "the yaml file configuration to be overlayed.")
	flag.BoolVar(&verbose, "verbose", false, "verbose output.")
	flag.Parse()

	if baseConfig == "" {
		log.Println("Required parameter missing: -base-config=...")
		log.Fatalln("Please use -h to see more information.")
	}

	if verbose {
		log.Printf("Generating configuration for: %s", environment)
	}

	// Template filename.
	baseYAML := readFile(baseConfig)

	// Create an instance of MarathonApp.
	t := MarathonApp{}

	// Load the base YAML configuration.
	t.LoadYAML(baseYAML)

	// Overlay the environment specific YAML configuration if specified.
	if overlayConfig != "" {
		overlayYAML := readFile(overlayConfig)
		t.LoadYAML(overlayYAML)
	}

	// Let's see what we have created!
	fmt.Println(t.ToJSON())
}

func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read the file: %s", filename)
	}
	return string(b[:])
}
