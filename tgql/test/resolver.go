//go:generate go run github.com/99designs/gqlgen

package test

import "context"

func Init() Config {
	return Config{
		Resolvers: &Resolver{},
	}
}

type Resolver struct {
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (this queryResolver) Hello(ctx context.Context) (*string, error) {
	return nil, nil
}
func (this mutationResolver) Hello(ctx context.Context) (*string, error) {
	return nil, nil
}
