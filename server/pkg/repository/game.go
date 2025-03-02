package repository

import (
	"database/sql"
	"fmt"
	"globetrotter/pkg/database"
	"globetrotter/pkg/models"
	"time"

	"gorm.io/gorm"
)

func NewGame(game *models.Game) error {
	return database.Get().Create(game).Error
}

func GetGame(sessionID int) (*models.Game, error) {
	var game models.Game

	err := database.Get().Model(&models.Game{SessionID: sessionID}).Scan(&game).Error

	return &game, err
}

func AddOpponent(gameID int, opponentUserID int) error {
	return database.Get().Model(&models.Game{ID: gameID}).Update("user_two_id", opponentUserID).Error
}

func CheckGameIsActiveAndNoOpponent(userID string, sessionID int) error {
	var count int64
	err := database.Get().Model(&models.Game{}).
		Where("session_id = ? AND is_completed = ? AND user_one_id = ? AND user_two_id = ?", sessionID, false, userID, 0).
		Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("no active game with session id")
	}

	return nil
}

func CheckGameIsActive(sessionID int) error {
	var count int64
	err := database.Get().Model(&models.Game{}).
		Where("session_id = ? AND is_completed = ?", sessionID, false).
		Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("no active game with session id")
	}

	return nil
}

func GetRandomCity() (models.City, error) {
	var city models.City

	err := database.Get().Model(&models.City{}).Order("RANDOM()").Limit(1).Scan(&city).Error
	if err != nil {
		return models.City{}, err
	}

	return city, nil
}

func GetClues(cityID int) ([]models.Clue, error) {
	var clues []models.Clue

	err := database.Get().Model(&models.Clue{CitiesID: cityID}).Scan(&clues).Error
	if err != nil {
		return nil, err
	}

	return clues, nil
}

func GetOptions(correctCityID int) ([]models.City, error) {
	var cities []models.City

	err := database.Get().Model(&models.City{}).Where("id != ?", correctCityID).Order("RANDOM()").Limit(3).Scan(&cities).Error
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func GetCorrectCity(clueID int) (models.City, error) {
	var city models.City

	err := database.Get().Model(&models.Clue{ID: clueID}).InnerJoins("cities on cities.id = clues.cities_id").Scan(&city).Error
	if err != nil {
		return models.City{}, err
	}

	return city, nil
}

func GetFacts(cityID int) ([]models.Fact, error) {
	var facts []models.Fact

	err := database.Get().Model(&models.Fact{CitiesID: cityID}).Select("fact").Scan(&facts).Error
	if err != nil {
		return nil, err
	}

	return facts, nil
}

func IncreaseScore(userID int, sessionID int) error {
	result := database.Get().Model(&models.Game{SessionID: sessionID, UserOneID: userID}).Update("score_one", gorm.Expr("score_one + 1"))
	affectedRows := result.RowsAffected
	err := result.Error

	if err != nil {
		return err
	}

	if affectedRows != 0 {
		return nil
	}

	return database.Get().Model(&models.Game{SessionID: sessionID, UserTwoID: userID}).Update("score_two", gorm.Expr("score_two + 1")).Error
}

func EndGame(sessionID int, completedByUserID int, winnerUserID int) error {
	return database.Get().Model(&models.Game{SessionID: sessionID}).Updates(&models.Game{
		CompletedAt:       sql.NullTime{Time: time.Now()},
		IsCompleted:       true,
		CompletedByUserID: completedByUserID,
		WinnerID:          winnerUserID,
	}).Error
}
