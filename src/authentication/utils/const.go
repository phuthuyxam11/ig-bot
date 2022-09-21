package utils

const (
	DbTypeTodo  = 1
	DbTypeTodo2 = 2
	DbTypeTodo3 = 3
	DbTypeUser  = 4
)
const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() int
}
