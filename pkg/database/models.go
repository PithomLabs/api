package database

type User struct {
	UserId int64 `gorm:"primary_key"`
	Name   string
	NSFW   bool
}
