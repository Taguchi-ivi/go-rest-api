package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITweetUsecase interface {
	GetAllTweet(userId uint) ([]model.TweetResponse, error)
	GetTweetById(userId uint, tweetId uint) (model.TweetResponse, error)
	CreateTweet(tweet model.Tweet) (model.TweetResponse, error)
	UpdateTweet(tweet model.Tweet, userId uint, tweetId uint) (model.TweetResponse, error)
	DeleteTweet(userId, tweetId uint) error
}

// tweetRepositoryの値を格納する構造体
type tweetUsecase struct {
	tr repository.ITweetRepository
	tv validator.ITweetValidator
}

func NewTweetUsecase(tr repository.ITweetRepository, tv validator.ITweetValidator) ITweetUsecase {
	return &tweetUsecase{tr, tv}
}

// repositoryのtweetRepositoryのGetAllTweetを呼び出す
func (tu *tweetUsecase) GetAllTweet(userId uint) ([]model.TweetResponse, error) {
	var tweets []model.Tweet
	if err := tu.tr.GetAllTweet(&tweets, userId); err != nil {
		return nil, err
	}
	resTweets := []model.TweetResponse{}
	for _, v := range tweets {
		t := model.TweetResponse{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTweets = append(resTweets, t)
	}
	return resTweets, nil
}

func (tu *tweetUsecase) GetTweetById(userId uint, tweetId uint) (model.TweetResponse, error) {
	tweet := model.Tweet{}
	// 成功するとtweetに値が入る(第一引数で渡したtweetのポインタに値が入る)
	if err := tu.tr.GetTweetById(&tweet, userId, tweetId); err != nil {
		return model.TweetResponse{}, err
	}
	resTweet := model.TweetResponse{
		ID:        tweet.ID,
		Title:     tweet.Title,
		Content:   tweet.Content,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}
	return resTweet, nil
}

// 更新処理をする際に、失敗したら構造体の0値と共にエラーを返す
func (tu *tweetUsecase) CreateTweet(tweet model.Tweet) (model.TweetResponse, error) {
	if err := tu.tv.TweetValidate(tweet); err != nil {
		return model.TweetResponse{}, err
	}
	if err := tu.tr.CreateTweet(&tweet); err != nil {
		return model.TweetResponse{}, err
	}
	resTweet := model.TweetResponse{
		ID:        tweet.ID,
		Title:     tweet.Title,
		Content:   tweet.Content,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}
	return resTweet, nil
}

func (tu *tweetUsecase) UpdateTweet(tweet model.Tweet, userId uint, tweetId uint) (model.TweetResponse, error) {
	if err := tu.tv.TweetValidate(tweet); err != nil {
		return model.TweetResponse{}, err
	}
	if err := tu.tr.UpdateTweet(&tweet, userId, tweetId); err != nil {
		return model.TweetResponse{}, err
	}
	resTweet := model.TweetResponse{
		ID:        tweet.ID,
		Title:     tweet.Title,
		Content:   tweet.Content,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}
	return resTweet, nil
}

func (tu *tweetUsecase) DeleteTweet(userId uint, tweetId uint) error {
	if err := tu.tr.DeleteTweet(userId, tweetId); err != nil {
		return err
	}
	return nil
}
