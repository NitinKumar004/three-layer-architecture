package user

type User struct {
	UserID    int    `json:"userid"`
	UserName  string `json:"username"`
	UserPhone string `json:"userphone"`
	UserEmail string `json:"useremail"`
}
