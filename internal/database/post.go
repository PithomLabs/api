package database

import "github.com/komfy/api/internal/structs"

func GetPostByID(id string) (*structs.Entity, error) {
	post := &structs.Entity{}
	// SELECT * FROM entities WHERE id = `id` LIMIT 1
	gErr := openDatabase.Instance.Where("id = ?", id).First(post).Error
	if gErr != nil {
		return nil, gErr
	}

	return post, nil
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
