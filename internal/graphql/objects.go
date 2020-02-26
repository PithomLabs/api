package graphql

import (
	"github.com/graphql-go/graphql"

	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
)

func Root() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			/*
				Arguments for queries
					@arg id: The user id inside users table in DB
					@arg:
					@arg:
			*/
			"user": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Type: User(),
				/*
					Resolve Function

				*/
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, iOk := params.Args["id"].(string)
					if !iOk {
						return nil, err.CreateArgumentsError("id", "string")
					}

					user, uErr := database.GetUserByID(id)
					if uErr != nil {
						return nil, uErr
					}

					return user, nil
				},
			},
			/*
				Arguments for queries
					@arg id: The post id inside the entities table in DB
					@arg:
					@arg:
			*/
			"post": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Type: Entity(),
				/*
					Resolve Function

				*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
			/*
				Arguments for queries
					@arg id: The comment id inside the entities table in DB
					@arg:
					@arg:
			*/
			"comment": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Type: Entity(),
				/*
					Resolve Function

				*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
		},
	})
}

func User() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
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
				Type: Settings(),
			},
			/*
				Arguments for queries
					@arg:
					@arg:
					@arg:
			*/
			"posts": &graphql.Field{
				Type: graphql.NewList(Entity()),
				/*
					Resolve Function

				*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
			/*
				Arguments for queries
					@arg:
					@arg:
					@arg:
			*/
			"comments": &graphql.Field{
				Type: graphql.NewList(Entity()),
				/*
					Resolve Function

				*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
		},
	})
}

func Settings() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
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
}

func Entity() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Entity",
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			return graphql.Fields{ // Start of FieldsThunk
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
				},
				"inside": &graphql.Field{
					Type: Content(),
				},
				/*
					Arguments for queries
						@arg:
						@arg:
						@arg:
				*/
				"user_infos": &graphql.Field{
					Type: EntityUser(),
					/*
						Resolve Function

					*/
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return nil, nil
					},
				},
				/*
					Arguments for queries
						@arg:
						@arg:
						@arg:
				*/
				"comments": &graphql.Field{
					Type: graphql.NewList(Entity()),
					/*
						Resolve Function

					*/
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return nil, nil
					},
				},
			} // End of FieldsThunk
		}),
	})
}

func Content() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Entity Content",
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
}

func EntityUser() *graphql.Object { // SAME AS UserInfos
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Entity User",
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
}
