package models

type Error struct {
	Message string `json:"messsage"`
}

func (e *Error) Error() string {
	return e.Message
}
