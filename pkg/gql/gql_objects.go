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
			Type: userWithPostGQL,
			Args: graphql.FieldConfigArgument{
				"userid": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				contextProvider := parameters.Context.Value("context_provider").(ContextProvider)

				// Check token, it must be valid in order to use the graphql queries
				_, err := jwt.IsTokenValid(contextProvider.Token)
				if err != nil {
					return nil, err

				}

				id, ok := parameters.Args["userid"].(int)
				if !ok {
					log.Print("userid should be an integer")
					return nil, nil
				}

				// Get the user from the ID
				strID := fmt.Sprintf("%v", id)
				user := contextProvider.Database.AskUserByID(strID)

				if user.Username == "" {
					return nil, nil
				}

				return user, nil
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
				contextProvider := parameters.Context.Value("context_provider").(ContextProvider)

				_, err := jwt.IsTokenValid(parameters.Context.Value("token").(string))
				if err != nil {
					return nil, err

				}

				id, ok := parameters.Args["postid"].(int)
				if !ok {
					log.Print("postid should be a integer.")

				}
				strID := fmt.Sprintf("%v", id)

				post := contextProvider.Database.AskPostByID(strID)
				if post.PostID == 0 {
					return nil, nil

				}

				return post, nil
			},
		},
	},
})

// User graphql object
var userWithPostGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"userid": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"NSFW": &graphql.Field{
			Type: graphql.Boolean,
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
				contextProvider := parameters.Context.Value("context_provider").(ContextProvider)

				_, err := jwt.IsTokenValid(contextProvider.Token)
				if err != nil {
					return nil, err

				}

				return nil, nil
			},
		},
	},
})

var basicUserGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"userid": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
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

// Post graphql object
var postGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"postid": &graphql.Field{
			Type: graphql.ID,
		},
		"user": &graphql.Field{
			Type: basicUserGQL,
			Resolve: func(parameters graphql.ResolveParams) (interface{}, error) {
				contextProvider := parameters.Context.Value("context_provider").(ContextProvider)

				_, err := jwt.IsTokenValid(contextProvider.Token)
				if err != nil {
					return nil, err

				}

				// The variable this represent the current post
				this := parameters.Source.(*db.Post)
				id := this.UserID
				strID := fmt.Sprintf("%v", id)

				user := contextProvider.Database.AskUserByID(strID)
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
