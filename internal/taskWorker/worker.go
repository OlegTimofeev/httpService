package taskWorker

import (
	"httpService/internal/models"
)

type Worker interface {
	AddRequest(*models.FetchTask)
	ListenForTasks()
}
