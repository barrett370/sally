// sally is an HTTP service that allows you to serve
// vanity import paths for your Go packages.
package main // import "go.uber.org/sally"

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/barrett370/sally/config"
	"github.com/barrett370/sally/handler"
	tmpls "github.com/barrett370/sally/templates"
	"github.com/barrett370/sally/utils"
)

func main() {
	yml := flag.String("yml", "sally.yaml", "yaml file to read config from")
	tpls := flag.String("templates", "", "directory of .html templates to use")
	port := flag.Int("port", 8080, "port to listen and serve on")
	flag.Parse()

	log.Printf("Parsing yaml at path: %s\n", *yml)
	config, err := config.Parse(*yml)
	if err != nil {
		log.Fatalf("Failed to parse %s: %v", *yml, err)
	}

	var templates *template.Template
	if *tpls != "" {
		log.Printf("Parsing templates at path: %s\n", *tpls)
		templates, err = utils.GetCombinedTemplates(*tpls)
		if err != nil {
			log.Fatalf("Failed to parse templates at %s: %v", *tpls, err)
		}
	} else {
		tmpls.Templates = templates
	}

	log.Printf("Creating HTTP handler with config: %v", config)
	handler, err := handler.CreateHandler(config, templates)
	if err != nil {
		log.Fatalf("Failed to create handler: %v", err)
	}

	log.Printf(`Starting HTTP handler on ":%d"`, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
