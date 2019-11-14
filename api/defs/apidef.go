package defs

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`

}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comments struct {
	Id       string
	VideoId  string
	AuthorId int
	Content  string
}

type SimpleSession struct {
	Username string // login name
	TTL      int64
}
