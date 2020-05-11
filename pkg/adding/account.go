package adding

type Account struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Note     string `json:"note"`
}
