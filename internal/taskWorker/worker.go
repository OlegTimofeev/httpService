package taskWorker

import (
	"httpService/internal/models"
)

type TaskHandler interface {
	AddRequest(*models.FetchTask, models.CanSetResponse)
	ListenForTasks()
}
