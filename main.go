package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"text/template"
)

// https://dev.to/kirklewis/go-text-template-processing-181d

func main() {
	type EnvironmentConfigs struct {
		EnvName string
	}

	devEnv := EnvironmentConfigs{
		"Dev",
	}
	// Here we created the tmplContent object which contains a sample template to be renderred.
	tmplContent, err := template.New("test").Parse("The EnvName should be set to the {{ .EnvName }} value. \n")
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(tmplContent))
	fmt.Println(tmplContent)
	// here is the actual rendering takes place, and sends the rendere content to the output
	tmplContent.Execute(os.Stdout, devEnv)

	fmt.Println("")
	fmt.Println("")

	// Get a list of all files in the "SourceTemplates" folder
	files, err := ioutil.ReadDir("SourceTemplates/anotherfolder")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}

	// https://yourbasic.org/golang/list-files-in-directory/ - this is a common approach, it uses 'functions literal syntax'
	// https://xojoc.pw/blog/golang-file-tree-traversal.html

	/*
	  Now we're going to get a template content from a temlate file,
	  render it, then output it into another file.
	*/
	tmplContentFolder, err := template.ParseFiles("./dummy-template.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reflect.TypeOf(tmplContentFolder))
	fmt.Println(tmplContentFolder)
	fmt.Println(&tmplContentFolder)

	os.Mkdir("./rendered-content", 0755)
	renderedfile, err := os.Create("./rendered-content/rendered-dummy-template.yml")

	// here is the actual rendering takes place.
	fmt.Println(reflect.TypeOf(renderedfile))
	tmplContentFolder.Execute(renderedfile, devEnv)

}
