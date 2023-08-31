package controller

type Controller struct {
	User    UserController
	Todo    TodoController
	Contest ContestController
	Game    GameController
}
