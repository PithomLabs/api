package structs

type Comment struct {
	CommentID int `gorm:"AUTO_INCREMENT"`
	PostID    int
	Inside    Content
	Likes     int
	Liked     bool
}
