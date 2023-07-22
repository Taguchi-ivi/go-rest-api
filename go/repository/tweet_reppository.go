package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITweetRepository interface {
	GetAllTweet(tweets *[]model.Tweet, userId uint) error
	GetTweetById(tweet *model.Tweet, userId uint, tweetId uint) error
	CreateTweet(tweet *model.Tweet) error
	UpdateTweet(tweet *model.Tweet, useId uint, tweetId uint) error
	DeleteTweet(userId, tweetId uint) error
}

type tweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) ITweetRepository {
	return &tweetRepository{db}
}

func (tr *tweetRepository) GetAllTweet(tweets *[]model.Tweet, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(tweets).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) GetTweetById(tweet *model.Tweet, userId uint, tweetId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).First(tweet, tweetId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) CreateTweet(tweet *model.Tweet) error {
	if err := tr.db.Create(tweet).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) UpdateTweet(tweet *model.Tweet, userId uint, tweetId uint) error {
	result := tr.db.Model(tweet).Clauses(clause.Returning{}).Where("user_id = ? AND id=?", userId, tweetId).Updates(map[string]interface{}{
		"title":   tweet.Title,
		"content": tweet.Content,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *tweetRepository) DeleteTweet(userId, tweetId uint) error {
	result := tr.db.Where("id = ? AND user_id", tweetId, userId).Delete(&model.Tweet{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

