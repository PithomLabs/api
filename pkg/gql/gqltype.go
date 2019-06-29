package gql

import (
	"log"

	"github.com/graphql-go/graphql"
	db "github.com/komfy/api/pkg/database"
)

// Main graphql struct
// This is the main query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: userGQL,
			Args: graphql.FieldConfigArgument{
				"userid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				id, ok := parameters.Args["userid"].(string)
				if !ok {
					log.Fatal("UserID should be a string")
					return nil, nil
				}

				// Get the user from the ID
				user := db.AskUserByID(id)

				if user.Name != "" {
					return user, nil
				}

				return nil, nil
			},
		},
		"post": &graphql.Field{
			Type: postGQL,
			Args: graphql.FieldConfigArgument{
				"postid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				// Connect to database
				// And return the given post
				// Which is define by this postid
				return nil, nil
			},
		},
	},
})

// User graphql object
var userGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"userid": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"NSFW": &graphql.Field{
			Type: graphql.Boolean,
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(postGQL),
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					// Must give a number of this user's posts
					// Can be zero
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				// Return the n-th first post of this user
				post := []db.Post{
					db.Post{
						PostID:      15425,
						Description: "THIS IS A TEST",
						Type:        1,
					},
				}
				return post, nil
			},
		},
	},
})

// Post graphql object
var postGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"postid": &graphql.Field{
			Type: graphql.ID,
		},
		"type": &graphql.Field{
			Type: contentTypeGQL,
		},
		"description": &graphql.Field{
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
