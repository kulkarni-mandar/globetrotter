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
	Correct bool   `json:"correct"`
	Facts   []Fact `json:"facts,omitempty"`
}

type ResponseEndGame struct {
	WinnerUserName string `json:"winner_user,omitempty"`
	Score          int    `json:"score"`
	Completed      bool   `json:"completed,omitempty"`
}
