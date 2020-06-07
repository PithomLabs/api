package structs

// User struct inside database
type User struct {
	ID       int64 `gorm:"column:user_id;primary_key"`
	Username string
	Password string
	Email    string
	Fullname string

	Bio       string
	AvatarURL string `gorm:"column:avatar_url;default:'default_komfy_profile_url'" json:"avatar_url"`
	CreatedAt uint64 `gorm:"column:created_at" json:"created_at"`
	Checked   bool
	// `-` means we ignore the settings field when working with gorm
	Settings Settings `gorm:"-"`
}

// Settings represent the user's account settings
type Settings struct {
	ID        uint  `gorm:"column:setting_id;primary_key"`
	UserID    int64 `gorm:"primary_key"`
	ShowLikes bool  `gorm:"column:show_likes" json:"show_likes"`
	ShowNSFW  bool  `gorm:"column:show_nsfw" json:"show_nsfw"`
	NSFWPage  bool  `gorm:"column:nsfw_page" json:"nsfw_page"`
}
