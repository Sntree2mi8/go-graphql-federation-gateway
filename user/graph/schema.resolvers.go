package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/Sntree2mi8/go-graphql-federation-gateway/user/graph/model"
	"github.com/vektah/gqlparser/v2/formatter"
)

func (r *queryResolver) Service(ctx context.Context) (*model.Service, error) {
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

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{
		{
			ID:   "user:1",
			Name: "Federation Taro",
			Reports: []*model.Report{
				{
					ID: "report:1",
				},
			},
		},
		{
			ID:   "user:2",
			Name: "Go Taro",
		},
	}, nil
}

func (r *queryResolver) Report(ctx context.Context, id string) (*model.Report, error) {
	return &model.Report{ID: id}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
