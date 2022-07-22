package web

type UserProfileResponse struct {
	Id       int    `json:"owner_id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
