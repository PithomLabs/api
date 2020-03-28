package graphql

import (
	"github.com/graphql-go/graphql"
)

// TODO: Change all the database errors to more related errors

func settings() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "UserSettings",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"show_likes": &graphql.Field{
				Type: graphql.Boolean,
			},
			"show_nsfw": &graphql.Field{
				Type: graphql.Boolean,
			},
			"nsfw_page": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})
}

func user() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			// GQL default types
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"fullname": &graphql.Field{
				Type: graphql.String,
			},
			"bio": &graphql.Field{
				Type: graphql.String,
			},
			"avatar_url": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.Int,
			},
			// TODO: Make settings a private field
			"settings": &graphql.Field{
				Type: settings(),
			},
			"posts": &graphql.Field{
				Type: graphql.NewList(entity()),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(
						params,
						userTypeEntitiesResolve,
						"post",
					)
				},
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(entity()),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(
						params,
						userTypeEntitiesResolve,
						"comment",
					)
				},
			},
		},
	})
}

func asset() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Asset",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"width": &graphql.Field{
				Type: graphql.Int,
			},
			"height": &graphql.Field{
				Type: graphql.Int,
			},
			"resource_type": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"secure_url": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
}

func content() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Content",
		Fields: graphql.Fields{
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"source": &graphql.Field{
				Type: graphql.NewList(asset()),
			},
			"nsfw": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})
}

func userInfos() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "UserInfos",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"avatar_url": &graphql.Field{
				Type: graphql.String,
			},
			"nsfw": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(
						params,
						userInfosTypeNsfwResolve,
					)
				},
			},
		},
	})
}

func entity() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Entity",
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"created_at": &graphql.Field{
					Type: graphql.Int,
				},
				"edited_at": &graphql.Field{
					Type: graphql.Int,
				},
				"likes": &graphql.Field{
					Type: graphql.Int,
				},
				"liked": &graphql.Field{
					Type: graphql.Boolean,
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return resolvePrivateField(
							params,
							entityTypeLikedResolve,
						)
					},
				},
				"answer_of": &graphql.Field{
					Type: graphql.Int,
					// TODO: Make this cleaner and do all the required checks
					Resolve: entityTypeAnswerOfResolve,
				},
				"inside": &graphql.Field{
					Type: content(),
				},
				"author": &graphql.Field{
					Type: userInfos(),
					/* Resolve Function */
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return resolvePublicField(
							params,
							entityTypeAuthorResolve,
						)
					},
				},
				"comments": &graphql.Field{
					Type: graphql.NewList(entity()),
					Args: graphql.FieldConfigArgument{
						"depth": &graphql.ArgumentConfig{
							Type:         graphql.Int,
							DefaultValue: 3,
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return resolvePublicField(
							params,
							entityTypeCommentsResolve,
						)
					},
				},
			}
		}),
	})
}

func Root() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/*
				Arguments for queries
					@arg id
					@arg username
			*/
			"user": &graphql.Field{
				Name: "RootUser",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Type: user(),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(
						params,
						rootTypeUserResolve,
					)
				},
			},
			/*
				Arguments for queries
					@arg id
			*/
			"post": &graphql.Field{
				Name: "RootPost",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Type: entity(),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(
						params,
						rootTypeEntityResolve,
						"post",
					)
				},
			},
			/*
				Arguments for queries
					@arg id
			*/
			"comment": &graphql.Field{
				Name: "RootComment",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Type: entity(),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(
						params,
						rootTypeEntityResolve,
						"comment",
					)
				},
			},
			/*
				Arguments for queries
					@arg last 'n': return the n most recent posts
			*/
			"posts": &graphql.Field{
				Name: "RootPosts",
				Args: graphql.FieldConfigArgument{
					"last": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Type: graphql.NewList(entity()),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(
						params,
						rootTypeEntitiesResolve,
						"post",
					)
				},
			},
		},
	})
}
