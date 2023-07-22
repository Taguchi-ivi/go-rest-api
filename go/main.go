package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

type dataFmt struct {
	names []string
	price int
	count int
}

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	tweetValidator := validator.NewTweetValidator()
	todoValidator := validator.NewTodoValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	tweetRepository := repository.NewTweetRepository(db)
	todoRepository := repository.NewTodoRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	tweetUsecase := usecase.NewTweetUsecase(tweetRepository, tweetValidator)
	todoUsecase := usecase.NewTodoUsecase(todoRepository, todoValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	tweetController := controller.NewTweetController(tweetUsecase)
	todoController := controller.NewTodoController(todoUsecase)
	e := router.NewRouter(userController, taskController, tweetController, todoController)

	e.Logger.Fatal(e.Start(":8080"))

	// pointer study
	// data := dataFmt{ []string{"apple", "banana"}, 100, 10 }
	// reset(&data)
	// fmt.Println(data)
}

// func reset(data *dataFmt) {
// 	data.names = append(data.names, "orange")
// 	data.price = 200
// 	data.count = 20
// }
