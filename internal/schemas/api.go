package schemas

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserID   int    `json:"userID"`
}

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	ID    int64  `json:id`
	//Author Author `json:"Author"`
}
