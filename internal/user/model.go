package user

type User struct {
	Email    string `gorm:"index"`
	Password string
	Name     string
}
