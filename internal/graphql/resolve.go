package graphql

import (
	"log"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/structs"
)

// ==============> returnFunction TYPE FUNCTIONS <==================

// ====> Resolve Public Field <======

/*
	userTypeEntitiesResolve is the resolve function of the user's entities fields (posts, comments).

	It needs as first args the entity type which can be "post" or "comment".
*/
func userTypeEntitiesResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
	// Made the choice of ignoring the cast's verification
	user := params.Source.(*structs.User)

	// Same as before, I choose to not do the verification
	posts, gErr := database.GetAllEntitiesFromUser(user.ID, args[0])
	if gErr != nil {
		log.Println("TEst", gErr)
		return nil, err.ErrInDatabaseOccured
	}

	return posts, nil
}

/*
	userInfosTypeNsfwResolve will return `user.Settings.NSFWPage` coz `user.NSFW` doesn't actually exist.
*/
func userInfosTypeNsfwResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
	user := params.Source.(*structs.User)

	return user.Settings.NSFWPage, nil
}

/*
	entityTypeAuthorResolve will return entity's user infos based on `structs.Entity.UserID`
*/
func entityTypeAuthorResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
	entity := params.Source.(*structs.Entity)

	user, gErr := database.GetUserByID(entity.UserID)
	if gErr != nil {
		log.Println(gErr)
		return nil, err.ErrInDatabaseOccured
	}

	return user, nil
}

/*
	entityTypeCommentsResolve will return n-depth of comments recursively.

	A depth of 0 will return as much comments as a depth of 1.
*/
func entityTypeCommentsResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
	depthLevel := params.Args["depth"].(int)
	comments, rErr := database.GetRecursiveEntities(
		params.Source.(*structs.Entity).ID,
		uint(depthLevel),
		0,
	)

	if rErr != nil {
		return nil, rErr
	}

	return comments, nil
}

func rootTypeEntityResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
	idArg, idOk := params.Args["id"]
	if idOk {
		// We don't need to check cast result because GraphQL will return an error
		// if params.Args["id"] is anything else than an integer.
		uid := idArg.(int)

		post, gErr := database.GetEntityByID(uint(uid), args[0])
		if gErr != nil {
			log.Println(gErr)
			return nil, err.ErrInDatabaseOccured
		}

		return post, nil
	}

	return nil, nil
}

func rootTypeUserResolve(params graphql.ResolveParams, token map[string]string, fnArgs ...string) (interface{}, error) {
	idArg, idOk := params.Args["id"]
	if idOk {
		uid := idArg.(int)

		user, gErr := database.GetUserByID(uint(uid))
		if gErr != nil {
			log.Println(gErr)
			return nil, err.ErrInDatabaseOccured
		}

		return user, nil
	}

	uname, uOk := params.Args["username"]
	if uOk {
		sUser := uname.(string)

		user, gErr := database.GetUserByName(sUser)
		if gErr != nil {
			log.Println(gErr)
			return nil, err.ErrInDatabaseOccured
		}

		return user, nil
	}

	return nil, nil
}

func rootTypeEntitiesResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
	numPostArg, uOk := params.Args["last"]
	if uOk {
		nPost := numPostArg.(int)

		posts, gErr := database.GetLastNEntities(uint(nPost), args[0])
		if gErr != nil {
			log.Println(gErr)
			return nil, err.ErrInDatabaseOccured
		}

		return posts, nil
	}

	return nil, nil
}

// ====> Resolve Private Field <======

/*
	entityTypeLikedResolve will return true if the user,
	which we can obtain using `token["ID"]`, has liked the current entity.
*/
func entityTypeLikedResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
	tokenUserID, cErr := strconv.Atoi(token["ID"])
	if cErr != nil {
		return nil, cErr
	}

	eid := params.Source.(*structs.Entity).ID
	return database.IsEntityLikedBy(eid, uint(tokenUserID))
}

/*
	userTypeSettingsResolve will return the user's settings if the user,
	which we can obtain using `token["ID"]`, as the same id as the one for which we are asking the settings.
*/
func userTypeSettingsResolve(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error) {
	tokenUserID, cErr := strconv.Atoi(token["ID"])
	if cErr != nil {
		return nil, cErr
	}

	user := params.Source.(*structs.User)
	if uint(tokenUserID) != user.ID {
		return nil, err.ErrIDDoesntMatch
	}

	return &user.Settings, nil
}

// ==============> CLASSIC FUNCTIONS <==================

/*
	entityTypeAnswerOfResolve will check if `structs.Entity.AnswerOf` is equal to 0, if so
	it will send nil so frontend can identify the comments root
*/
func entityTypeAnswerOfResolve(params graphql.ResolveParams) (interface{}, error) {
	currentEntity := params.Source.(*structs.Entity)
	if currentEntity.AnswerOf == 0 {
		return nil, nil
	}

	return currentEntity.AnswerOf, nil
}
