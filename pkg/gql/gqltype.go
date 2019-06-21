package gql

import (
	"github.com/graphql-go/graphql"
)

// Main graphql struct
// This is the main query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: UserGQL,
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				// Connect to database and return the user
				// Matching the given user_id
				return nil, nil
			},
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(PostGQL),
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				// Connect to database
				// And return the first n-th posts
				return nil, nil
			},
		},
	},
})

// User graphql struct
var UserGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"user_id": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"NSFW": &graphql.Field{
			Type: graphql.Boolean,
		},
		"Posts": &graphql.Field{
			Type: graphql.NewList(PostGQL),
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					// Must give a number of this user's posts
					// Can be zero
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				// Return the n-th first post of this user
				return nil, nil
			},
		},
	},
})

// Post graphql struct
var PostGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"post_id": &graphql.Field{
			Type: graphql.ID,
		},
		"contentType": &graphql.Field{
			Type: contentTypeGQL,
		},
		"descr": &graphql.Field{
			Type: graphql.String,
		},
		"likes": &graphql.Field{
			Type: graphql.Int,
		},
		"liked": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

// Enum content type
// Don't actually know if this is gonna work
// Will see when we could do gql request to now
// :D
var contentTypeGQL = graphql.NewEnum(graphql.EnumConfig{
	Name: "contentType",
	Values: graphql.EnumValueConfigMap{
		"text": &graphql.EnumValueConfig{
			Value: 0,
		},
		"image": &graphql.EnumValueConfig{
			Value: 1,
		},
	},
})
