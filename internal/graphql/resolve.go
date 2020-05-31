package graphql

import (
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/komfy/api/internal/database"
	"github.com/komfy/api/internal/structs"
)

//RootResolver represents a root query for graphql
type RootResolver struct{}

//User resolves user type
func (r *RootResolver) User(args struct{ UserID graphql.ID }) (*UserResolver, error) {
	user, err := database.GetUserByID(string(args.UserID))
	if err != nil {
		return &UserResolver{&structs.User{}}, err
	}
	return &UserResolver{user}, err
}

//Posts resolves Posts field
func (r *RootResolver) Posts(args struct{ UserID graphql.ID }) ([]*PostResolver, error) {
	pResolvers := []*PostResolver{}
	posts, err := database.GetAllEntitiesFromUser(args.UserID, "post")
	if err != nil {
		return pResolvers, err
	}
	for _, post := range *posts {
		pResolvers = append(pResolvers, &PostResolver{post})
	}
	return pResolvers, err
}

//Post resolves single post
func (r *RootResolver) Post(args struct{ PostID graphql.ID }) (*PostResolver, error) {
	post, err := database.GetEntityByID(args.PostID, "post")
	if err != nil {
		return &PostResolver{post}, err
	}
	return &PostResolver{post}, nil
}

//Comment is a resolver for comment
//It returns PostResolver, because comment and post is the same entity.
//So, the way to resolve them is same.
func (r *RootResolver) Comment(args struct{ CommentID graphql.ID }) (*PostResolver, error) {
	return r.Post(struct{ PostID graphql.ID }{args.CommentID})
}

//Comments is root resolver for comments by userID or PostID
func (r *RootResolver) Comments(args struct {
	PostID graphql.ID
	UserID graphql.ID
}) ([]*PostResolver, error) {
	if args.PostID != "" {
		comments, err := database.GetEntitiesByAnswerOf(args.PostID)
		if err != nil {
			return nil, err
		}
		prs := make([]*PostResolver, 0)
		for _, comment := range comments {
			prs = append(prs, &PostResolver{comment})
		}
		return prs, nil
	}
	if args.UserID != "" {
		return r.Posts(struct{ UserID graphql.ID }{args.UserID})
	}
	return nil, fmt.Errorf("expected any of given args to be in query: PostID, UserID")
}

//UserInfo resolves info about author of post
func (r *RootResolver) UserInfo(args struct{ UserID graphql.ID }) (*UserResolver, error) {
	return r.User(struct{ UserID graphql.ID }{args.UserID})
}

//UserResolver ...
type UserResolver struct{ *structs.User }

//ID ...
func (ur *UserResolver) ID() graphql.ID {
	return ur.User.ID
}

//Username ...
func (ur *UserResolver) Username() string {
	return ur.User.Username
}

//Fullname ...
func (ur *UserResolver) Fullname() string {
	return ur.User.Fullname
}

//Bio ...
func (ur *UserResolver) Bio() string {
	return ur.User.Bio
}

//CreatedAt ...
func (ur *UserResolver) CreatedAt() uint64 {
	return ur.User.CreatedAt
}

//Posts ...
func (ur *UserResolver) Posts() ([]*PostResolver, error) {
	root := &RootResolver{}
	return root.Posts(struct{ UserID graphql.ID }{ur.User.ID})
}

//PostResolver ...
type PostResolver struct{ *structs.Entity }

//ID ...
func (pr *PostResolver) ID() graphql.ID {
	return pr.Entity.ID
}

//Likes ...
func (pr *PostResolver) Likes() uint {
	return pr.Entity.Likes
}

//Liked ...
func (pr *PostResolver) Liked() bool {
	return pr.Liked()
}

//Inside ...
func (pr *PostResolver) Inside() (*ContentResolver, error) {
	return &ContentResolver{&pr.Entity.Inside}, nil
}

//Comments is comment resolver for post
func (pr *PostResolver) Comments() ([]*PostResolver, error) {
	root := &RootResolver{}
	return root.Comments(struct {
		PostID graphql.ID
		UserID graphql.ID
	}{PostID: pr.Entity.ID})
}

//CreatedAt ...
func (pr *PostResolver) CreatedAt() uint64 {
	return pr.Entity.CreatedAt
}

//EditedAt ...
func (pr *PostResolver) EditedAt() uint64 {
	return pr.Entity.EditedAt
}

//ContentResolver resolves content of post
type ContentResolver struct{ *structs.Content }

//Type ...
func (cr *ContentResolver) Type() string {
	return cr.Content.Type
}

//Text ...
func (cr *ContentResolver) Text() string {
	return cr.Content.Text
}

//Nsfw ...
func (cr *ContentResolver) Nsfw() bool {
	return cr.Content.NSFW
}

//Author resolves author field for post
func (pr *PostResolver) Author(args struct{ UserID graphql.ID }) (*UserResolver, error) {
	root := &RootResolver{}
	return root.UserInfo(args)
}

//Source resolves
func (cr *ContentResolver) Source() []*AssetResolver {
	assetResolvers := make([]*AssetResolver, 0)
	for _, source := range cr.Content.Source {
		assetResolvers = append(assetResolvers, &AssetResolver{&source})
	}
	return assetResolvers
}

//AssetResolver ...
type AssetResolver struct{ *structs.Asset }

//ID ...
func (ar *AssetResolver) ID() graphql.ID {
	return ar.Asset.ID
}

//Width ...
func (ar *AssetResolver) Width() uint {
	return ar.Asset.Width
}

//Height ...
func (ar *AssetResolver) Height() uint {
	return ar.Asset.Height
}

//ResourceType ...
func (ar *AssetResolver) ResourceType() string {
	return ar.Asset.ResourceType
}

//Url ...
func (ar *AssetResolver) Url() string {
	return ar.Asset.URL
}

//SecureUrl ...
func (ar *AssetResolver) SecureUrl() string {
	return ar.Asset.SecureURL
}

//CreatedAt ...
func (ar *AssetResolver) CreatedAt() uint64 {
	return ar.Asset.CreatedAt
}

// TODO: Find out what it is
// func (ar *AssetResolver) Alt() string{
// 	return ar.Asset.A
// }

// // ==============> returnFunction TYPE FUNCTIONS <==================

// // ====> Resolve Public Field <======

// /*
// 	userTypeEntitiesResolve is the resolve function of the user's entities fields (posts, comments).

// 	It needs as first args the entity type which can be "post" or "comment".
// */
// func userTypeEntitiesResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
// 	// Made the choice of ignoring the cast's verification
// 	user := params.Source.(*structs.User)

// 	// Same as before, I choose to not do the verification
// 	posts, gErr := database.GetAllEntitiesFromUser(user.ID, args[0])
// 	if gErr != nil {
// 		log.Println("TEst", gErr)
// 		return nil, err.ErrInDatabaseOccured
// 	}

// 	return posts, nil
// }

// /*
// 	userInfosTypeNsfwResolve will return `user.Settings.NSFWPage` coz `user.NSFW` doesn't actually exist.
// */
// func userInfosTypeNsfwResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
// 	user := params.Source.(*structs.User)

// 	return user.Settings.NSFWPage, nil
// }

// /*
// 	entityTypeAuthorResolve will return entity's user infos based on `structs.Entity.UserID`
// */
// func entityTypeAuthorResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
// 	entity := params.Source.(*structs.Entity)

// 	user, gErr := database.GetUserByID(entity.UserID)
// 	if gErr != nil {
// 		log.Println(gErr)
// 		return nil, err.ErrInDatabaseOccured
// 	}

// 	return user, nil
// }

// /*
// 	entityTypeCommentsResolve will return n-depth of comments recursively.

// 	A depth of 0 will return as much comments as a depth of 1.
// */
// func entityTypeCommentsResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
// 	depthLevel := params.Args["depth"].(int)
// 	comments, rErr := database.GetRecursiveEntities(
// 		params.Source.(*structs.Entity).ID,
// 		uint(depthLevel),
// 		0,
// 	)

// 	if rErr != nil {
// 		return nil, rErr
// 	}

// 	return comments, nil
// }

// func rootTypeEntityResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
// 	idArg, idOk := params.Args["id"]
// 	if idOk {
// 		// We don't need to check cast result because GraphQL will return an error
// 		// if params.Args["id"] is anything else than an integer.
// 		uid := idArg.(int)

// 		post, gErr := database.GetEntityByID(uint(uid), args[0])
// 		if gErr != nil {
// 			log.Println(gErr)
// 			return nil, err.ErrInDatabaseOccured
// 		}

// 		return post, nil
// 	}

// 	return nil, nil
// }

// func rootTypeUserResolve(params graphql.ResolveParams, token map[string]string, fnArgs ...string) (interface{}, error) {
// 	idArg, idOk := params.Args["id"]
// 	if idOk {
// 		uid := idArg.(int)

// 		user, gErr := database.GetUserByID(uint(uid))
// 		if gErr != nil {
// 			log.Println(gErr)
// 			return nil, err.ErrInDatabaseOccured
// 		}

// 		return user, nil
// 	}

// 	uname, uOk := params.Args["username"]
// 	if uOk {
// 		sUser := uname.(string)

// 		user, gErr := database.GetUserByName(sUser)
// 		if gErr != nil {
// 			log.Println(gErr)
// 			return nil, err.ErrInDatabaseOccured
// 		}

// 		return user, nil
// 	}

// 	return nil, nil
// }

// func rootTypeEntitiesResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
// 	numPostArg, uOk := params.Args["last"]
// 	if uOk {
// 		nPost := numPostArg.(int)

// 		posts, gErr := database.GetLastNEntities(uint(nPost), args[0])
// 		if gErr != nil {
// 			log.Println(gErr)
// 			return nil, err.ErrInDatabaseOccured
// 		}

// 		return posts, nil
// 	}

// 	return nil, nil
// }

// // ====> Resolve Private Field <======

// /*
// 	entityTypeLikedResolve will return true if the user,
// 	which we can obtain using `token["ID"]`, has liked the current entity.
// */
// func entityTypeLikedResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
// 	tokenUserID, cErr := strconv.Atoi(token["ID"])
// 	if cErr != nil {
// 		return nil, cErr
// 	}

// 	eid := params.Source.(*structs.Entity).ID
// 	return database.IsEntityLikedBy(eid, uint(tokenUserID))
// }

// /*
// 	userTypeSettingsResolve will return the user's settings if the user,
// 	which we can obtain using `token["ID"]`, as the same id as the one for which we are asking the settings.
// */
// func userTypeSettingsResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
// 	tokenUserID, cErr := strconv.Atoi(token["ID"])
// 	if cErr != nil {
// 		return nil, cErr
// 	}

// 	user := params.Source.(*structs.User)
// 	if uint(tokenUserID) != user.ID {
// 		return nil, err.ErrIDDoesntMatch
// 	}

// 	return &user.Settings, nil
// }

// // ==============> CLASSIC FUNCTIONS <==================

// /*
// 	entityTypeAnswerOfResolve will check if `structs.Entity.AnswerOf` is equal to 0, if so
// 	it will send nil so frontend can identify the comments root
// */
// func entityTypeAnswerOfResolve(params graphql.ResolveParams) (interface{}, error) {
// 	currentEntity := params.Source.(*structs.Entity)
// 	if currentEntity.AnswerOf == 0 {
// 		return nil, nil
// 	}

// 	return currentEntity.AnswerOf, nil
// }
