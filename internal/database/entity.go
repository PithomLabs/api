package database

import (
	"github.com/komfy/api/internal/structs"
)

func GetEntityByID(id int64, eType string) (*structs.Entity, error) {
	entity := &structs.Entity{}
	// SELECT * FROM entities WHERE entity_id = `id` LIMIT 1
	gErr := openDatabase.Instance.Where("entity_id = ? AND type = ?", id, eType).First(entity).Error
	if gErr != nil {
		return nil, gErr
	}

	aErr := GetAssetsForEntity(entity)
	if aErr != nil {
		return nil, aErr
	}

	return entity, nil
}

func GetAllEntitiesFromUser(uid int64, eType string) (*[]*structs.Entity, error) {
	entities := &[]*structs.Entity{}
	// SELECT entities.* FROM entities JOIN users on entities.user_id = users.user_id WHERE entities.user_id = `uid` AND entities.type = `eType`
	gErr := openDatabase.Instance.Joins("JOIN users on entities.user_id = users.user_id").Where("entities.user_id = ? AND entities.type = ?", uid, eType).Select("entities.*").Find(entities).Error
	if gErr != nil {
		return nil, gErr
	}

	var aErr error
	for _, entity := range *entities {
		aErr = GetAssetsForEntity(entity)
		if aErr != nil {
			return nil, aErr
		}
	}

	return entities, nil
}

func GetAssetsForEntity(entity *structs.Entity) error {
	// If the content type of the entity is text then we don't need
	// to fetch the sources, because there aren't any
	if entity.Inside.Type == "text" {
		return nil
	}

	// SELECT assets.* FROM entities JOIN assets ON assets.entity_id = entities.entity_id WHERE entities.entity_id = `entity.ID`;
	gErr := openDatabase.Instance.Table("entities").Joins("JOIN assets ON assets.entity_id = entities.entity_id").Where("assets.entity_id = ?", entity.ID).Select("assets.*").Find(&entity.Inside.Source).Error
	if gErr != nil {
		return gErr
	}

	return nil
}

func IsEntityLikedBy(eid, uid uint) (bool, error) {
	count := 0
	openDatabase.Instance.Table("likes").Where("user_id = ? and entity_id = ?", uid, eid).Count(&count)

	return count == 1, nil
}

func GetLastNEntities(numOfEntities uint, eType string) (*[]*structs.Entity, error) {
	entities := &[]*structs.Entity{}
	// SELECT * FROM entities WHERE type = `eType` ORDER BY created_at DESC LIMIT `numOfEntities`
	gErr := openDatabase.Instance.Limit(numOfEntities).Where("type = ?", eType).Order("created_at desc").Find(entities).Error
	if gErr != nil {
		return nil, gErr
	}

	for _, entity := range *entities {
		GetAssetsForEntity(entity)
	}

	return entities, nil
}

func GetEntitiesByAnswerOf(eid int64) ([]*structs.Entity, error) {
	currentDepthEntities := []*structs.Entity{}

	// SELECT * FROM entities WHERE answer_of=`eid`;
	gErr := openDatabase.Instance.Where("answer_of = ?", eid).Find(&currentDepthEntities).Error
	if gErr != nil {
		return nil, gErr
	}

	return currentDepthEntities, nil
}
