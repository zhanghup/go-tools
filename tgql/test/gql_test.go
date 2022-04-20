package main_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/zhanghup/go-tools/tgql/test/graph"
	"github.com/zhanghup/go-tools/tgql/test/graph/generated"
	"testing"
)

func TestName(t *testing.T) {

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))
	res := map[string]any{}
	c.MustPost(`
	query Todo{
	todos{
		id
		text
		done
		user{id name}
	}
}
`, &res)

}
