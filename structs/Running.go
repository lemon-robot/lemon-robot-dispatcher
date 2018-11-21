package structs

import "container/list"

type Running struct {
	Logs *list.List
	CancelFunc func()
}