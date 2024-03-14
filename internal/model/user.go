package model

type RegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
