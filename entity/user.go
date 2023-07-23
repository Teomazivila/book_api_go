package entity

type User struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	Token          string `json:"token"`
}
