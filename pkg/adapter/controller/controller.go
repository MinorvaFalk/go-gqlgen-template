package controller

type Controller struct {
	User interface{ User }
	Todo interface{ Todo }
}
