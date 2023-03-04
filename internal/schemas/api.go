package schemas

type Author struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	ID    int64  `json:id`
	//Author Author `json:"Author"`
}
