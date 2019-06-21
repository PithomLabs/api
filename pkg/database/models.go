package database

type User struct {
	UserId uint64 `gorm:"AUTO_INCREMENT"` // This is the primary key
	Name   string
	NSFW   bool
}
