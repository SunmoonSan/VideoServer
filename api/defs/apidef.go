package defs

// requests
type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewComment struct {
	AuthorId int    `json:"author_id"`
	Content  string `json:"content"`
}

type NewVideo struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
}

// response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type UserSession struct {
	Username  string `json:"user_name"`
	SessionId string `json:"session_id"`
}

type UserInfo struct {
	Id int `json:"id"`
}

type SignedIn struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// Data model
type User struct {
	Id        int
	LoginName string
	Pwd       string
}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id       string
	VideoId  string
	AuthorId int
	Content  string
}

type SimpleSession struct {
	Username string // login name
	TTL      int64
}
