package main

import (
	"flag"
	"fmt"
)

// SET UP GIT and basic CI/CD with git actions for the generated thing

func getFlags() string {
	name := flag.String("n", "myapp", "App name")
	flag.Parse()
	return *name
}

func main() {
	name := getFlags()

	if name == "" {
		fmt.Println("Usage -n [projectName]")
	}

	mkdir(name)
	chdir(name)
	goModInit(name)

	mainBoiler := "package main\n\nimport (\"fmt\")\n\nfunc main(){\n\tfmt.Println(\"Hello world!\")\n}"
	writeFile("main.go", []byte(mainBoiler))

	readMeBoiler := fmt.Sprintf("## {%s}\n### This project was initialised using goboiler CLI tool", name)
	writeFile("README.md", []byte(readMeBoiler))
}
