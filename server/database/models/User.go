package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null json:firstname"`
	LastName  string `gorm:"not null" json:"lastname"`
	Verified  bool   `gorm:"default:false"`
	CreatedAt int64
	UserLogin UserLogin `gorm:"embedded"`
}

type UserLogin struct {
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type Users struct {
	Users []User `json:"users"`
}
