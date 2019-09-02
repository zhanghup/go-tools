package tools

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	str, err := Str().Template(`
		mutation {{title .user}}Create($input:New{{title .user}}!){
			{{.user}}_create(input:$input){
				id
				type
				account
				password
				name
				avatar
				i_card
				sex
				mobile
				admin
				weight
				status
				created
				updated
			}
		}
	`, map[string]interface{}{
		"user":"user",
	},nil)
	fmt.Println(str, err)
}

func TestBolck(t *testing.T) {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}