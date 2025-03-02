package service

import (
	"database/sql"
	"errors"
	"globetrotter/pkg/models"
	"globetrotter/pkg/repository"
	"time"
)

func NewGame(userName string) (models.ResponseNewGame, error) {
	userID, err := repository.GetUserID(userName)
	if err != nil {
		return models.ResponseNewGame{}, err
	}

	game := &models.Game{
		CreatedAt: sql.NullTime{Time: time.Now()},
		SessionID: int(time.Now().Unix()),
		UserOneID: userID,
		UserTwoID: 0,
		ScoreOne:  0,
		ScoreTwo:  0,
		WinnerID:  0,
	}

	err = repository.NewGame(game)
	if err != nil {
		return models.ResponseNewGame{}, err
	}

	return models.ResponseNewGame{SessionID: game.SessionID}, nil
}

func JoinGame(userName string, sessionID int) (models.ResponseJoinGame, error) {
	userID, err := repository.GetUserID(userName)
	if err != nil {
		return models.ResponseJoinGame{}, err
	}

	game, err := repository.GetGame(sessionID)
	if err != nil {
		return models.ResponseJoinGame{}, err
	}

	if game.IsCompleted {
		return models.ResponseJoinGame{}, errors.New("joining finished game")
	}

	err = repository.AddOpponent(game.ID, userID)
	if err != nil {
		return models.ResponseJoinGame{}, err
	}

	opponentUserName, err := repository.GetUserName(game.UserOneID)
	if err != nil {
		return models.ResponseJoinGame{}, err
	}

	return models.ResponseJoinGame{
		SessionID:        sessionID,
		OpponentUserName: opponentUserName,
	}, nil
}

func InviteToGame(userName string, sessionID int) (models.ResponseInviteToGame, error) {
	err := repository.CheckGameIsActiveAndNoOpponent(userName, sessionID)
	if err != nil {
		return models.ResponseInviteToGame{}, err
	}

	// game is active and can be joined
	return models.ResponseInviteToGame{
		SessionID: sessionID,
	}, nil
}

func NextQuestion(userName string, sessionID int) (models.ResponseNextQuestion, error) {
	err := repository.CheckGameIsActive(sessionID)
	if err != nil {
		return models.ResponseNextQuestion{}, err
	}

	city, err := repository.GetRandomCity()
	if err != nil {
		return models.ResponseNextQuestion{}, err
	}

	clues, err := repository.GetClues(city.ID)
	if err != nil {
		return models.ResponseNextQuestion{}, err
	}

	options, err := repository.GetOptions(city.ID)
	if err != nil {
		return models.ResponseNextQuestion{}, err
	}

	options = append(options, city)

	return models.ResponseNextQuestion{
		Clues:   clues,
		Options: options,
	}, nil
}

func ValidateAnswer(userName string, sessionID int, clueID int, selectedCityID int) (models.ResponseValidateAnswer, error) {
	err := repository.CheckGameIsActive(sessionID)
	if err != nil {
		return models.ResponseValidateAnswer{}, err
	}

	userID, err := repository.GetUserID(userName)
	if err != nil {
		return models.ResponseValidateAnswer{}, err
	}

	correctCity, err := repository.GetCorrectCity(clueID)
	if err != nil {
		return models.ResponseValidateAnswer{}, err
	}

	isCorrect := correctCity.ID == selectedCityID

	if !isCorrect {
		return models.ResponseValidateAnswer{Correct: false}, nil
	}

	// answer is correct given by user
	facts, err := repository.GetFacts(selectedCityID)
	if err != nil {
		return models.ResponseValidateAnswer{}, err
	}

	err = repository.IncreaseScore(userID, sessionID)
	if err != nil {
		return models.ResponseValidateAnswer{}, err
	}

	return models.ResponseValidateAnswer{
		Correct: isCorrect,
		Facts:   facts,
	}, nil
}

func EndGame(userName string, sessionID int) (models.ResponseEndGame, error) {
	err := repository.CheckGameIsActive(sessionID)
	if err != nil {
		return models.ResponseEndGame{}, err
	}

	userID, err := repository.GetUserID(userName)
	if err != nil {
		return models.ResponseEndGame{}, err
	}

	game, err := repository.GetGame(sessionID)
	if err != nil {
		return models.ResponseEndGame{}, err
	}

	var winnerID int

	if game.ScoreOne > game.ScoreTwo {
		winnerID = game.UserOneID
	} else {
		winnerID = game.UserTwoID
	}

	winnerUserName, err := repository.GetUserName(winnerID)
	if err != nil {
		return models.ResponseEndGame{}, err
	}

	game.IsCompleted = true
	game.CompletedAt = sql.NullTime{Time: time.Now()}
	game.CompletedByUserID = userID

	err = repository.EndGame(sessionID, userID, game.WinnerID)
	if err != nil {
		return models.ResponseEndGame{}, err
	}

	currentScore := 0
	if game.UserOneID == userID {
		currentScore = game.ScoreOne
	} else {
		currentScore = game.ScoreTwo
	}

	return models.ResponseEndGame{
		WinnerUserName: winnerUserName,
		Completed:      true,
		Score:          currentScore,
	}, nil
}
