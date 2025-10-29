package commands

type LoginUserCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
