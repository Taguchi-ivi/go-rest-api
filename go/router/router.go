package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController, ttc controller.ITweetController, tttc controller.ITodoController) *echo.Echo {
	e := echo.New()

	// corsのmiddlewareを設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	// csrfのmiddlewareを設定
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// postmainで動作確認をする際はnonemodeだとエラーになるため、defaultmodeに変更
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge: 60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	t := e.Group("/tasks")
	// useを使うことでエンドポイントにミドルウェアを設定できる
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	tw := e.Group("/tweets")
	tw.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	tw.GET("", ttc.GetAllTweet)
	tw.GET("/:tweetId", ttc.GetTweetById)
	tw.POST("", ttc.CreateTweet)
	tw.PUT("/:tweetId", ttc.UpdateTweet)
	tw.DELETE("/:tweetId", ttc.DeleteTweet)

	td := e.Group("/todos")
	td.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	td.GET("", tttc.GetAllTodos)
	td.GET("/:todoId", tttc.GetTodoById)
	td.POST("", tttc.CreateTodo)
	td.PUT("/:todoId", tttc.UpdateTodo)
	td.DELETE("/:todoId", tttc.DeleteTodo)

	return e
}
