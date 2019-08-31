package tools

import (
	"fmt"
	"testing"
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
