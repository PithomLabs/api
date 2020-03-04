package database

import (
	"github.com/komfy/api/internal/structs"
)

func GetEntityByID(id, eType string) (*structs.Entity, error) {
	entity := &structs.Entity{}
	// SELECT * FROM entities WHERE id = `id` LIMIT 1
	gErr := openDatabase.Instance.Where("id = ? AND type = ?", id, eType).First(entity).Error
	if gErr != nil {
		return nil, gErr
	}

	return entity, nil
}

func GetAllEntitiesFromUser(uid, eType string) (*[]structs.Entity, error) {
	entities := &[]structs.Entity{}

	gErr := openDatabase.Instance.Joins("JOIN users ON entities.user_id = users.id").Where("entities.user_id = ? AND entities.type = ?", uid, eType).Find(entities).Error
	if gErr != nil {
		return nil, gErr
	}

	return entities, nil
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
