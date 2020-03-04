package graphql

import (
	"log"
	"strconv"

	"github.com/graphql-go/graphql"

	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/structs"
)

var settings *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Settings",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
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

var user *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		// GQL default types
		"id": &graphql.Field{
			Type: graphql.ID,
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
		// Custom GQL types
		"settings": &graphql.Field{
			Type: settings,
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(entity),
			/* Resolve Function */
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return resolvePublicField(params,
					func(token interface{}) (interface{}, error) {
						user, ok := params.Source.(*structs.User)
						if !ok {
							log.Println("params.Source couldn't be parse as structs.User")
							return nil, err.ErrServerSide
						}

						sUID := strconv.Itoa(int(user.ID))
						posts, gErr := database.GetAllEntitiesFromUser(sUID, "post")
						if gErr != nil {
							log.Println(gErr)
							return nil, err.ErrInDatabaseOccured
						}

						return posts, nil
					},
				)
			},
		},
		"comments": &graphql.Field{
			Type: graphql.NewList(entity),
			/* Resolve Function */
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return resolvePublicField(params,
					func(token interface{}) (interface{}, error) {
						user, ok := params.Source.(*structs.User)
						if !ok {
							log.Println("params.Source couldn't be parse as structs.User")
							return nil, err.ErrServerSide
						}

						sUID := strconv.Itoa(int(user.ID))
						comments, gErr := database.GetAllEntitiesFromUser(sUID, "comment")
						if gErr != nil {
							log.Println(gErr)
							return nil, err.ErrInDatabaseOccured
						}

						return comments, nil
					},
				)
			},
		},
	},
})

var content *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Content",
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"source": &graphql.Field{
			Type: graphql.String,
		},
		"NSFW": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var entityUser *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "EUser",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"avatar_url": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var entity *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Entity",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
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
				return resolvePrivateField(params,
					func(token interface{}) (interface{}, error) {
						return nil, nil
					},
				)
			},
		},
		"inside": &graphql.Field{
			Type: content,
		},
		"user_infos": &graphql.Field{
			Type: entityUser,
			/* Resolve Function */
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return resolvePublicField(params,
					func(token interface{}) (interface{}, error) {
						entity, ok := params.Source.(*structs.Entity)
						if !ok {
							log.Println("params.Source couldn't be parse as structs.Entity")
							return nil, err.ErrServerSide
						}

						sUID := strconv.Itoa(int(entity.UserID))
						user, gErr := database.GetUserByID(sUID)
						if gErr != nil {
							log.Println(gErr)
							return nil, err.ErrInDatabaseOccured
						}

						return user, nil
					},
				)
			},
		},
	},
})

var root *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		/*
			Arguments for queries
				@arg id
				@arg username
		*/
		"user": &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Type: user,
			/* Resolve Function */
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return resolvePublicField(params,
					// Anonymous function that will be executed by resolvePublicField
					func(token interface{}) (interface{}, error) {
						id, idOk := params.Args["id"]
						if idOk {
							sID, sOk := id.(string)
							if !sOk {
								return nil, err.CreateArgumentsError("id", "string")
							}

							user, gErr := database.GetUserByID(sID)
							if gErr != nil {
								log.Println(gErr)
								return nil, err.ErrInDatabaseOccured
							}

							return user, nil
						}

						uname, uOk := params.Args["username"]
						if uOk {
							sUser, sOk := uname.(string)
							if !sOk {
								return nil, err.CreateArgumentsError("username", "string")
							}

							user, gErr := database.GetUserByName(sUser)
							if gErr != nil {
								log.Println(gErr)
								return nil, err.ErrInDatabaseOccured
							}

							return user, nil
						}

						return nil, nil
					},
				)
			},
		},
		/*
			Arguments for queries
				@arg id
		*/
		"post": &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Type: entity,
			/* Resolve Function */
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return resolvePublicField(params,
					func(token interface{}) (interface{}, error) {
						id, idOk := params.Args["id"]
						if idOk {
							sID, sOk := id.(string)
							if !sOk {
								return nil, err.CreateArgumentsError("id", "string")
							}

							post, gErr := database.GetEntityByID(sID, "post")
							if gErr != nil {
								log.Println(gErr)
								return nil, err.ErrInDatabaseOccured
							}

							return post, nil
						}

						return nil, nil
					},
				)
			},
		},
		/*
			Arguments for queries
				@arg last 'n': return the n most recent posts
		*/
		"posts": &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"last": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Type: graphql.NewList(entity),
			/* Resolve Function */
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return resolvePublicField(params,
					func(token interface{}) (interface{}, error) {
						nPost, uOk := params.Args["last"]
						if uOk {
							sNPost, sOk := nPost.(string)
							if !sOk {
								return nil, err.CreateArgumentsError("last", "string")
							}

							posts, gErr := database.GetLastNPosts(sNPost)
							if gErr != nil {
								log.Println(gErr)
								return nil, err.ErrInDatabaseOccured
							}

							return posts, nil
						}

						return nil, nil
					},
				)
			},
		},
		/*
			Arguments for queries
				@arg id
		*/
		"comment": &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Type: entity,
			/* Resolve Function */
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return resolvePublicField(params,
					func(token interface{}) (interface{}, error) {
						id, idOk := params.Args["id"]
						if idOk {
							sID, sOk := id.(string)
							if !sOk {
								return nil, err.CreateArgumentsError("id", "string")
							}

							comment, gErr := database.GetEntityByID(sID, "comment")
							if gErr != nil {
								log.Println(gErr)
								return nil, err.ErrInDatabaseOccured
							}

							return comment, nil
						}

						return nil, nil
					},
				)
			},
		},
	},
})
