package graphql

import (
	"fmt"

	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/structs"

	"github.com/graphql-go/graphql"
)

// Main graphql struct
// This is the main query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: user,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				return generalResolveFunc(parameters,
					// The resolveFunc for this field
					func(context ContextProvider, tokenInfos interface{}) (interface{}, error) {
						// We do not need to verify if postid is really an integer
						// because graphql will do it for us
						id, _ := parameters.Args["userid"].(int)

						strID := fmt.Sprintf("%v", id)
						user := database.UserByID(strID)

						if user.Username == "" {
							return nil, err.ErrUserDoesntExist
						}

						return user, nil
					})
			},
		},
		"post": &graphql.Field{
			Type: postGQL,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				return generalResolveFunc(parameters, func(context ContextProvider, tokenInfos interface{}) (interface{}, error) {
					// We do not need to verify if postid is really an integer
					// because graphql will do it for us
					id, _ := parameters.Args["postid"].(int)

					post := database.PostByID(id)
					if post.PostID == 0 {
						return nil, err.ErrPostDoesntExist

					}

					return post, nil
				})
			},
		},
	},
})

// User graphql object with post field and
// a resolve function
var user = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"userid": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"NSFW": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				// NSFW is a private field, which need authentication
				return privatiseField(parameters, func(context ContextProvider) (interface{}, error) {
					sourceUser, ok := parameters.Source.(*structs.User)
					if !ok {
						return nil, nil
					}

					return sourceUser.NSFW, nil
				})
			},
		},
		"avatar": &graphql.Field{
			Type: graphql.String,
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(postGQL),
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					// Must give a number of this user's posts
					// Can be zero
					Type: graphql.Int,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				// Still need to be implemented
				return nil, nil
			},
		},
	},
})

// The basic user which is used inside the post object
// in order to get post's user infos
var basicUserGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"userid": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"NSFW": &graphql.Field{
			Type: graphql.Boolean,
		},
		"avatar": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// Post graphql object is the object which correspond to
// a post inside the database
var postGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"postid": &graphql.Field{
			Type: graphql.ID,
		},
		"user": &graphql.Field{
			Type: basicUserGQL,
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				// Still need to be implemented
				return nil, nil
			},
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
