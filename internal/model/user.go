package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null,unique"`
	Password string `json:"password"`
}

type CreateUserReq struct {
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null,unique"`
	Password string `json:"password" gorm:"not null"`
}

type UserRes struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
