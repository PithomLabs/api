package database

// User is the struct used by db and gql
type User struct {
	UserID   int    `gorm:"AUTO_INCREMENT"` // This is the primary key
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string
	NSFW     bool
}

// Post is the struct used by db and gql
type Post struct {
	PostID      int `gorm:"AUTO_INCREMENT"`
	Type        int
	Description string
	Likes       int
	Liked       bool
}
