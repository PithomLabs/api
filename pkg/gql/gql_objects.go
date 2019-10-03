package gql

import (
	"fmt"
	"log"

	"github.com/komfy/api/pkg/jwt"

	"github.com/graphql-go/graphql"
	db "github.com/komfy/api/pkg/database"
)

// Main graphql struct
// This is the main query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: user,
			Args: graphql.FieldConfigArgument{
				"userid": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				return generalResolveFunc(parameters, func(context ContextProvider, tokenInfos interface{}) (interface{}, error) {
					id, ok := parameters.Args["userid"].(int)
					if !ok {
						log.Print("userid should be an integer")
						return nil, nil
					}

					// Get the user from the ID
					strID := fmt.Sprintf("%v", id)
					user := context.Database.AskUserByID(strID)

					if user.Username == "" {
						return nil, nil
					}

					return user, nil
				})
			},
		},
		"post": &graphql.Field{
			Type: postGQL,
			Args: graphql.FieldConfigArgument{
				"postid": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				return generalResolveFunc(parameters, func(context ContextProvider, tokenInfos interface{}) (interface{}, error) {
					id, ok := parameters.Args["postid"].(int)
					if !ok {
						log.Print("postid should be a integer.")
						return nil, nil
					}
					strID := fmt.Sprintf("%v", id)

					post := context.Database.AskPostByID(strID)
					if post.PostID == 0 {
						return nil, nil

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
					sourceUser, ok := parameters.Source.(*db.User)
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
				context, cErr := getContext(parameters)
				if cErr != nil {
					return nil, cErr
				}

				_, err := jwt.IsTokenValid(context.Token)
				if err != nil {
					return nil, err

				}

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
				context, cErr := getContext(parameters)
				if cErr != nil {
					return nil, nil
				}

				// The variable sourcePost represent the current post
				sourcePost := parameters.Source.(*db.Post)
				id := fmt.Sprintf("%v", sourcePost.UserID)

				user := context.Database.AskUserByID(id)
				if user.Username == "" {
					return nil, nil

				}

				return user, nil
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
