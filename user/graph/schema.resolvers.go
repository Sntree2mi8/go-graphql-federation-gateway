package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/vektah/gqlparser/v2/formatter"
	"strings"

	"github.com/Sntree2mi8/go-graphql-federation-gateway/user/graph/model"
)

func (r *queryResolver) Service(_ context.Context) (*model.Service, error) {
	s := new(strings.Builder)
	f := formatter.NewFormatter(s)
	// parsedSchema is in the generated code
	f.FormatSchema(parsedSchema)

	service := model.Service{
		Name:    "user-service",
		Version: "0.1.0",
		Schema:  s.String(),
	}
	return &service, nil
}

func (r *queryResolver) Users(_ context.Context) ([]*model.User, error) {
	return []*model.User{
		{
			ID:   "user:1",
			Name: "Federation Taro",
		},
		{
			ID:   "user:2",
			Name: "Go Taro",
		},
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
