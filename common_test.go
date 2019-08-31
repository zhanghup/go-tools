package tools

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	str, err := Str().Template(`
		mutation {{.User}}Create($input:New{{.User}}!){
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
		"User":"User",
		"user":"user",
	})
	fmt.Println(str, err)
}
