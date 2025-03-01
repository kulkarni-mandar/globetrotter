package models

import "time"

type City struct {
	ID      int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	City    string `gorm:"column:city" json:"city"`
	Country string `gorm:"column:country" json:"country"`
}

type Clue struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CitiesID int    `gorm:"column:cities_id" json:"cities_id"`
	Clue     string `gorm:"column:clue" json:"clue"`
}

type Fact struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CitiesID int    `gorm:"column:cities_id" json:"cities_id"`
	Fact     string `gorm:"column:fact" json:"fact"`
}

type Game struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	StartedAt   time.Time `gorm:"column:started_at" json:"started_at"`
	CompletedAt time.Time `gorm:"column:completed_at" json:"completed_at"`
	SessionID   int       `gorm:"column:session_id" json:"session_id"`
	UserOneID   int       `gorm:"column:user_one_id" json:"user_one_id"`
	UserTwoID   int       `gorm:"column:user_two_id" json:"user_two_id"`
	ScoreOne    int       `gorm:"column:score_one" json:"score_one"`
	ScoreTwo    int       `gorm:"column:score_two" json:"score_two"`
	WinnerID    int       `gorm:"column:winner_id" json:"winner_id"`
}
