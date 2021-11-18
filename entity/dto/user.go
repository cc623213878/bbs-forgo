package dto

type User struct {
	Username string `binding:"gte=6,lte=24"`
	Password string `binding:"gte=8,lte=24"`
}

type UserInfo struct {
	User
	Phone string
	Email string
}
