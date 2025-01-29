package models

type CreateAction struct {
	Name string `form:"name"`
}

type Credential struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
