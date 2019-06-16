package database

type User struct {
	userId int64 `gorm:"primary_key"`
	name   string
	NSFW   bool
}
