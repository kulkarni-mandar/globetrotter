package models

type Error struct {
	Message string `json:"messsage"`
}

func (e Error) Error() string {
	return e.Message
}

type ResponseNewGame struct {
	SessionID int `json:"session_id"`
}

type ResponseJoinGame struct {
	SessionID        int    `json:"session_id"`
	OpponentUserName string `json:"opponent_user_name"`
}

type ResponseInviteToGame struct {
	SessionID int `json:"session_id"`
}

type ResponseNextQuestion struct {
	Clues   []Clue `json:"clues"`
	Options []City `json:"options"`
}

type ResponseValidateAnswer struct {
	Correct bool   `json:"bool"`
	Facts   []Fact `json:"facts,omitempty"`
}

type ResponseEndGame struct {
	WinnerUserName string `json:"winner_user,omitempty"`
	Completed      bool   `json:"completed,omitempty"`
}
