package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
)

var (
	templateBuffer bytes.Buffer
	environment    string
	verbose        bool
)

func main() {

	flag.StringVar(&environment, "env", "test", "the environment configuration to be generated.")
	flag.BoolVar(&verbose, "verbose", false, "verbose output.")
	flag.Parse()

	if verbose {
		log.Printf("Generating configuration for: %s", environment)
	}

	// Put it all here for now until we know that it's working as expected.

	// Config instance.
	//app := MarathonApp{}

	// Template filename.
	//templateFile := "test-data/marathon.json"
	baseYAML := readFile("test-data/marathon.yml")
	overlayYAML := readFile("test-data/marathon-prod.yml")
	//templateName := "marathon-test"

	// Yaml to struct.
	t := MarathonApp{}
	t.ParseYAML(baseYAML)
	t.ParseYAML(overlayYAML)
	// err := yaml.Unmarshal([]byte(yamlFile), &t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// if verbose {
	// 	spew.Dump(t)
	// }

	// Override the defaults.
	// err = yaml.Unmarshal([]byte(yamlFileProd), &t)
	// if err != nil {
	// 	log.Fatalf("error unmarshalling prod: %v", err)
	// }
	// if verbose {
	// 	spew.Dump(t)
	// }

	// Marshall it to JSON and see what we get...
	// marathonConfig, err := json.Marshal(t)
	// if err != nil {
	// 	log.Fatalf("error marshalling json: %s", err.Error())
	// }
	// fmt.Println(string(marathonConfig))
	log.Printf("JSON: %s", t.ToJSON())

	// Read in the template.
	//templateBody := readFile(templateFile)

	// Parse the template.
	//tmpl, err := template.New(templateName).Funcs(funcMap).Parse(templateBody)

	// Execute the template.
	//tmpl.Execute(&templateBuffer, app)
}

/*
var funcMap = template.FuncMap{
	"sampleFunc": sampleFunc,
}

func sampleFunc() string {
	return ""
}
*/
func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read the file: %s", filename)
	}
	return string(b[:])
}
