package main

import (
	"os"
	"text/template"
)

func main() {
	type EnvironmentConfigs struct {
		EnvName string
	}

	devEnv := EnvironmentConfigs{
		"Dev",
	}
	// Here we created the tmplContent object which contains a sample template to be renderred.
	tmplContent, _ := template.New("test").Parse("The EnvName should be set to the {{ .EnvName }} value. \n")
	// here is the actual rendering takes place.
	tmplContent.Execute(os.Stdout, devEnv)

	pwd, _ := os.Getwd()

	tmplContentFolder, _ := template.New("foldercontent").ParseFiles(pwd + "/MyTemplateFolder/dummy-template.yml")
	// here is the actual rendering takes place.
	tmplContentFolder.Execute(os.Stdout, devEnv)

}
