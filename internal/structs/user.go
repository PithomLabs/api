package structs

// User is the representation of a database user
type User struct {
	UserID   int    `gorm:"AUTO_INCREMENT"` // This is the primary key
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string
	NSFW     bool
	NSFWPage bool
	Checked  bool
	Posts    []Post
}
