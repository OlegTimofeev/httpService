package models

type CanSetResponse interface {
	SetResponse(id int, response *TaskResponse, err error) error
}

type WorkerPool interface {
	AddRequest(*FetchTask)
}
