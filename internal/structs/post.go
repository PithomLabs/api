package structs

type Post struct {
	PostID   int `gorm:"AUTO_INCREMENT"`
	UserID   int
	Inside   Content
	Likes    int
	Liked    bool
	NSFW     bool
	Comments []Comment
}
