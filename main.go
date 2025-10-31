package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

func getFlags() string {
	name := flag.String("n", "myapp", "App name")
	flag.Parse()
	return *name
}

func getTemplateData(name string) map[string]string {
	return map[string]string{
		"ProjectName": name,
	}
}

func getTemplates(embededFiles embed.FS) []*template.Template {
	const TEMPLATE_PATH = "templates"
	fileNames, err := embededFiles.ReadDir(TEMPLATE_PATH)

	if err != nil {
		fmt.Println("Templates not found")
	}

	var templates []*template.Template

	for _, file := range fileNames {
		tmpl, err := template.ParseFS(embededFiles, filepath.Join(TEMPLATE_PATH, file.Name()))
		if err == nil {
			templates = append(templates, tmpl)
		}
	}
	return templates
}

func initiateTemplates(templates []*template.Template, data map[string]string) {
	for _, template := range templates {
		fileName := strings.TrimSuffix(template.Name(), ".tmpl")
		file, _ := os.Create(fileName)
		template.Execute(file, data)
	}
}

func main() {
	name := getFlags()

	if name == "" {
		fmt.Println("Usage -n [projectName]")
	}

	mkdir(name)
	chdir(name)
	goModInit(name)

	data := getTemplateData(name)
	fileTemplates := getTemplates(templateFS)

	initiateTemplates(fileTemplates, data)
}
