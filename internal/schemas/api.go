package schemas

type UserAuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserID   int64  `json:"userID"`
}

type Article struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	ArticleID int64  `json:"articleID"`
	AuthorID  int64  `json:"authorID"`
}

type Response[T Article | []Article | string] struct {
	Error     string `json:"error,omitempty"`
	ErrorCode int    `json:"errorCode,omitempty"`
	Body      T      `json:"body,omitempty"`
}
