package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/Sntree2mi8/go-graphql-federation-gateway/report/graph/model"
	"github.com/vektah/gqlparser/v2/formatter"
)

func (r *queryResolver) Service(_ context.Context) (*model.Service, error) {
	s := new(strings.Builder)
	f := formatter.NewFormatter(s)
	// parsedSchema is in the generated code
	f.FormatSchema(parsedSchema)

	service := model.Service{
		Name:    "report-service",
		Version: "0.1.0",
		Schema:  s.String(),
	}
	return &service, nil
}

func (r *queryResolver) Reports(_ context.Context) ([]*model.Report, error) {
	return []*model.Report{
		{
			ID:      "report:1",
			Content: "美味しかった",
		},
		{
			ID:      "report:2",
			Content: "楽しかった",
		},
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
