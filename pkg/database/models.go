package database

// User is the struct used by db and gql
type User struct {
	UserID   int    `gorm:"AUTO_INCREMENT"` // This is the primary key
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string
	NSFW     bool
	Checked  bool
}

// Post is the struct used by db and gql
type Post struct {
	PostID      int `gorm:"AUTO_INCREMENT"`
	UserID      int
	Type        int
	Description string
	Likes       int
	Liked       bool
}
