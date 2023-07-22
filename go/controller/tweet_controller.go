package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITweetController interface {
	GetAllTweet(c echo.Context) error
	GetTweetById(c echo.Context) error
	CreateTweet(c echo.Context) error
	UpdateTweet(c echo.Context) error
	DeleteTweet(c echo.Context) error
}

type tweetController struct {
	tu usecase.ITweetUsecase
}

func NewTweetController(tu usecase.ITweetUsecase) ITweetController {
	return &tweetController{tu}
}

func (tc *tweetController) GetAllTweet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// claimsしたデータはany型で返却されるのでuint型に型アサーションする
	tweetRes, err := tc.tu.GetAllTweet(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tweetRes)
}

func (tu *tweetController) GetTweetById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)
	tweetRes, err := tu.tu.GetTweetById(uint(userId.(float64)), uint(tweetId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tweetRes)
}

func (tu *tweetController) CreateTweet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tweet := model.Tweet{}
	if err := c.Bind(&tweet); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tweet.UserId = uint(userId.(float64))
	tweetRes, err := tu.tu.CreateTweet(tweet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, tweetRes)
}

func (tu *tweetController) UpdateTweet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)

	tweet := model.Tweet{}
	if err := c.Bind(&tweet); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tweetRes, err := tu.tu.UpdateTweet(tweet, uint(userId.(float64)), uint(tweetId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tweetRes)
}

func (tu *tweetController) DeleteTweet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)

	err := tu.tu.DeleteTweet(uint(userId.(float64)), uint(tweetId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
