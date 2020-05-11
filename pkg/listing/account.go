package listing

type Account struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Note     string `json:"note"`
}
