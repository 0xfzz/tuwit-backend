package entity

type UserAccountEntity struct {
	Id       int
	Username string
	Email    string
	Password []byte
}
type UserEntity struct {
	Id             int    `json:"id,omitempty"`
	Username       string `json:"username,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       []byte
	DisplayName    string `json:"display_name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	Banner         string `json:"banner,omitempty"`
}
