package structs

type Soc_auth_data struct {
	Username  string `json:"username", db:"username"`
	Email     string `json:"email", db:"email"`
	BirthDate string `json:"birthdate", db:"birthdate"`
}
