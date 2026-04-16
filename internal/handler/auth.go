package handler

type AuthApiHandler struct {
	name string
}

func NewAuth() *AuthApiHandler {
	return &AuthApiHandler{}
}

func (a *AuthApiHandler) Login() {

}
