package database

import "github.com/komfy/api/internal/structs"

func PostByID(postid int) *structs.Post {
	post := &structs.Post{}
	openDatabase.Instance.First(&post, "post_id = ?", postid)

	return post
}
