package database

import (
	"github.com/komfy/api/internal/structs"
)

func GetEntityByID(id, eType string) (*structs.Entity, error) {
	entity := &structs.Entity{}
	// SELECT * FROM entities WHERE entity_id = `id` LIMIT 1
	gErr := openDatabase.Instance.Where("entity_id = ? AND type = ?", id, eType).First(entity).Error
	if gErr != nil {
		return nil, gErr
	}

	aErr := GetAssetsForPost(entity)
	if aErr != nil {
		return nil, aErr
	}

	return entity, nil
}

func GetAllEntitiesFromUser(uid, eType string) (*[]*structs.Entity, error) {
	entities := &[]*structs.Entity{}
	// TODO: Change Join to Natural Join
	// TODO: Add the SQL raw query
	gErr := openDatabase.Instance.Joins("JOIN users ON entities.user_id = users.user_id").Where("entities.user_id = ? AND entities.type = ?", uid, eType).Find(entities).Error
	if gErr != nil {
		return nil, gErr
	}

	var aErr error
	for _, entity := range *entities {
		aErr = GetAssetsForPost(entity)
		if aErr != nil {
			return nil, aErr
		}
	}

	return entities, nil
}

func GetAssetsForPost(entity *structs.Entity) error {
	// SELECT assets.* FROM entities JOIN contain ON contain.entity_id = entities.entity_id JOIN assets ON assets.asset_id = contain.asset_id WHERE entities.entity_id = `entity.ID`;
	gErr := openDatabase.Instance.Table("entities").Joins("JOIN contain ON contain.entity_id = entities.entity_id").Joins("JOIN assets ON assets.asset_id = contain.asset_id").Where("contain.entity_id = ?", entity.ID).Select("assets.*").Find(&entity.Inside.Source).Error
	if gErr != nil {
		return gErr
	}

	return nil
}

func UserLikedEntity(uid, eid string) (bool, error) {
	count := 0
	openDatabase.Instance.Table("likes").Where("user_id = ? and entity_id = ?", uid, eid).Count(&count)

	return count == 1, nil
}

func GetLastNPosts(numOfPosts string) (*[]structs.Entity, error) {
	posts := &[]structs.Entity{}
	// SELECT * FROM entities WHERE type = 'post' ORDER BY created_at DESC LIMIT `numsOfPosts`
	gErr := openDatabase.Instance.Limit(numOfPosts).Where("type = 'post'").Order("created_at desc").Find(posts).Error
	if gErr != nil {
		return nil, gErr
	}

	return posts, nil
}
