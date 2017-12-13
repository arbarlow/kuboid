package kuboid

import (
	"context"

	graphql "github.com/neelance/graphql-go"
)

var schema = `
	schema {
		query: Query
		// mutation: Mutation
	}

	type Query {
		graph(): ServiceGraph
	}

	// type Mutation {
	// 	createReview(episode: Episode!, review: ReviewInput!): Review
	// }

	type ServiceGraph {
		name: String!
		// services: [String]!
		// connections: [Connection]!
	}

	type Connection {
		source: String!
		destination: String!
		rps: Float
	}
`

var Schema *graphql.Schema

func init() {
	Schema = graphql.MustParseSchema(schema, &Resolver{})
}

type ServiceGraph struct {
	name string
}

func (s ServiceGraph) Name() string {
	return s.name
}

type Resolver struct{}

func (r *Resolver) Graph(ctx context.Context) *ServiceGraph {
	return &ServiceGraph{name: "hello"}
}
