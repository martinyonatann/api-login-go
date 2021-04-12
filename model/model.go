package model

type User struct {
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type ResponseResult struct {
	ResponseCode int    `json:"rc"`
	Error        string `json:"error"`
	Result       string `json:"result"`
}
