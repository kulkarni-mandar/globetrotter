package models

import "time"

type City struct {
	ID      int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	City    string `gorm:"column:city" json:"city"`
	Country string `gorm:"column:country" json:"country"`
}

type Clue struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CitiesID int32  `gorm:"column:cities_id" json:"cities_id"`
	Clue     string `gorm:"column:clue" json:"clue"`
}

type Fact struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CitiesID int32  `gorm:"column:cities_id" json:"cities_id"`
	Fact     string `gorm:"column:fact" json:"fact"`
}

type Game struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	StartedAt   time.Time `gorm:"column:started_at" json:"started_at"`
	CompletedAt time.Time `gorm:"column:completed_at" json:"completed_at"`
	UserOneID   int32     `gorm:"column:user_one_id" json:"user_one_id"`
	UserTwoID   int32     `gorm:"column:user_two_id" json:"user_two_id"`
	ScoreOne    int32     `gorm:"column:score_one" json:"score_one"`
	ScoreTwo    int32     `gorm:"column:score_two" json:"score_two"`
	WinnerID    int32     `gorm:"column:winner_id" json:"winner_id"`
}
