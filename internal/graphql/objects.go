package graphql

import (
	"log"

	"github.com/graphql-go/graphql"

	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/structs"
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
					return resolvePublicField(params,
						func(token interface{}) (interface{}, error) {
							user, ok := params.Source.(*structs.User)
							if !ok {
								log.Println("params.Source couldn't be parse as structs.User")
								return nil, err.ErrServerSide
							}

							posts, gErr := database.GetAllEntitiesFromUser(user.ID, "post")
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
				Type: graphql.NewList(entity()),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(params,
						func(token interface{}) (interface{}, error) {
							user, ok := params.Source.(*structs.User)
							if !ok {
								log.Println("params.Source couldn't be parse as structs.User")
								return nil, err.ErrServerSide
							}

							comments, gErr := database.GetAllEntitiesFromUser(user.ID, "comment")
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

func entityUser() *graphql.Object {
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
					return resolvePublicField(params,
						func(token interface{}) (interface{}, error) {
							user, ok := params.Source.(*structs.User)
							if !ok {
								log.Println("params.Source couldn't be parse as structs.User")
								return nil, err.ErrServerSide
							}

							return user.Settings.NSFWPage, nil
						},
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
						return resolvePrivateField(params,
							func(token interface{}) (interface{}, error) {
								return nil, nil
							},
						)
					},
				},
				"answer_of": &graphql.Field{
					Type: graphql.Int,
					// TODO: Make this cleaner and do all the required checks
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if p.Source.(*structs.Entity).AnswerOf == 0 {
							return nil, nil
						}

						return p.Source.(*structs.Entity).AnswerOf, nil
					},
				},
				"inside": &graphql.Field{
					Type: content(),
				},
				"author": &graphql.Field{
					Type: entityUser(),
					/* Resolve Function */
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return resolvePublicField(params,
							func(token interface{}) (interface{}, error) {
								entity, ok := params.Source.(*structs.Entity)
								if !ok {
									log.Println("params.Source couldn't be parse as structs.Entity")
									return nil, err.ErrServerSide
								}

								user, gErr := database.GetUserByID(entity.UserID)
								if gErr != nil {
									log.Println(gErr)
									return nil, err.ErrInDatabaseOccured
								}

								return user, nil
							},
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
						return resolvePublicField(params, func(token interface{}) (interface{}, error) {
							comments, rErr := database.GetRecursiveEntities(
								params.Source.(*structs.Entity).ID,
								uint(params.Args["depth"].(int)),
								0,
							)

							if rErr != nil {
								return nil, rErr
							}

							return comments, nil
						})
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
					return resolvePublicField(params,
						// Anonymous function that will be executed by resolvePublicField
						func(token interface{}) (interface{}, error) {
							idArg, idOk := params.Args["id"]
							if idOk {
								uid, iOk := idArg.(int)
								if !iOk {
									return nil, err.CreateArgumentsError("id", "int")
								}

								user, gErr := database.GetUserByID(uint(uid))
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
				Name: "RootPost",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Type: entity(),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(params,
						func(token interface{}) (interface{}, error) {
							idArg, idOk := params.Args["id"]
							if idOk {
								uid, iOk := idArg.(int)
								if !iOk {
									return nil, err.CreateArgumentsError("id", "int")
								}

								post, gErr := database.GetEntityByID(uint(uid), "post")
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
				Name: "RootPosts",
				Args: graphql.FieldConfigArgument{
					"last": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Type: graphql.NewList(entity()),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(params,
						func(token interface{}) (interface{}, error) {
							numPostArg, uOk := params.Args["last"]
							if uOk {
								nPost, sOk := numPostArg.(int)
								if !sOk {
									return nil, err.CreateArgumentsError("last", "int")
								}

								posts, gErr := database.GetLastNEntities(uint(nPost), "post")
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
				Name: "RootComment",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Type: entity(),
				/* Resolve Function */
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return resolvePublicField(params,
						func(token interface{}) (interface{}, error) {
							idArg, idOk := params.Args["id"]
							if idOk {
								eid, sOk := idArg.(int)
								if !sOk {
									return nil, err.CreateArgumentsError("id", "int")
								}

								comment, gErr := database.GetEntityByID(uint(eid), "comment")
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
}
