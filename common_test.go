package tools

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	str, err := StrTemplate(`
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
		"user": "user",
	}, nil)
	fmt.Println(str, err)
}

func TestBolck(t *testing.T) {
	const (
		master = `{{ table "user" }}`
	)
	var (
		funcs = template.FuncMap{"table": func(str string) string{
			fmt.Println(str)
			return "ddddddddddddddd"
		}}
		guardians = "xxqq_user"
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}

	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}
