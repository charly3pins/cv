package main

import (
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	datafileName   = "me.yaml"
	resumeFileName = "resume.html"
)

func main() {
	files := []string{
		"./tmpl/base.html",
		"./tmpl/partials/about.html",
		"./tmpl/partials/education.html",
		"./tmpl/partials/experience.html",
		"./tmpl/partials/keywords.html",
		"./tmpl/partials/languages.html",
		"./tmpl/partials/mission.html",
		"./tmpl/partials/projects.html",
		"./tmpl/partials/skills.html",
	}

	templates, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		return
	}

	data := map[string]interface{}{}
	buf, err := os.ReadFile(datafileName)
	if err != nil {
		log.Println("err reading data file", err)
		return
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		log.Println("err unmarshaling data", err)
		return
	}

	out, err := os.Create(resumeFileName)
	if err != nil {
		log.Println("error creating file", resumeFileName, err)
		return
	}
	err = templates.ExecuteTemplate(out, "base", data)
	if err != nil {
		log.Print(err.Error())
	}
}
