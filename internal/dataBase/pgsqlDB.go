package dataBase

import (
	"errors"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"httpService/internal/models"
	"log"
)

type PostgresDB struct {
	pgdb *pg.DB
}

func connect(config ConfigDB) *pg.DB {
	postgres := pg.Connect(&pg.Options{User: config.User, Password: config.Password, Database: config.Dbname})
	return postgres
}

func NewPGStore(config ConfigDB) *PostgresDB {
	db := new(PostgresDB)
	db.pgdb = connect(config)

	if err := db.pgdb.RunInTransaction(func(tx *pg.Tx) error {
		if err := db.pgdb.DropTable((*models.FetchTask)(nil), &orm.DropTableOptions{
			IfExists: true,
			Cascade:  true,
		}); err != nil {
			return err
		}
		if err := db.pgdb.CreateTable((*models.FetchTask)(nil), &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		}); err != nil {
			return err
		}
		if err := db.pgdb.DropTable((*models.TaskResponse)(nil), &orm.DropTableOptions{
			IfExists: true,
			Cascade:  true,
		}); err != nil {
			return err
		}
		if err := db.pgdb.CreateTable((*models.TaskResponse)(nil), &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		}); err != nil {
			return err
		}
		if cons := db.pgdb.PoolStats().Hits; cons < 1 {
			return errors.New("no connection")
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return db
}

func (db *PostgresDB) AddFetchTask(task *models.FetchTask) (*models.FetchTask, error) {
	if err := db.pgdb.Insert(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (db *PostgresDB) UpdateFetchTask(task models.FetchTask) error {
	if err := db.pgdb.Update(&task); err != nil {
		return err
	}
	return nil
}

func (db *PostgresDB) GetFetchTask(taskId int) (*models.FetchTask, error) {
	ft := models.FetchTask{ID: taskId}
	if err := db.pgdb.Select(&ft); err != nil {
		return nil, err
	}
	return &ft, nil
}

func (db *PostgresDB) GetAllTasks() ([]*models.FetchTask, error) {
	var tasks []*models.FetchTask
	if err := db.pgdb.Model(&tasks).Select(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (db *PostgresDB) DeleteFetchTask(taskId int) error {
	ft := models.FetchTask{ID: taskId}
	if err := db.pgdb.Delete(&ft); err != nil {
		return err
	}
	resp := models.TaskResponse{ID: taskId}
	if err := db.pgdb.Delete(&resp); err != nil {
		return err
	}
	return nil
}

func (db *PostgresDB) AddTaskResponse(res *models.TaskResponse) (*models.TaskResponse, error) {
	err := db.pgdb.Insert(res)
	return res, err
}

func (db *PostgresDB) GetTaskResponseByFtID(taskId int) (*models.TaskResponse, error) {
	tr := models.TaskResponse{ID: taskId}
	if err := db.pgdb.Select(&tr); err != nil {
		return nil, err
	}
	return &tr, nil
}

func (db *PostgresDB) SetResponse(id int, response *models.TaskResponse, errTask error) error {
	task, err := db.GetFetchTask(id)
	if err != nil {
		return err
	}
	task.Status = models.StatusInProgress
	if err := db.UpdateFetchTask(*task); err != nil {
		return err
	}
	response.ID = id
	if errTask != nil {
		response.Err = errTask.Error()
	}
	if err := db.pgdb.Insert(response); err != nil {
		return err
	}
	if errTask != nil {
		task.Status = models.StatusError
	} else {
		task.Status = models.StatusCompleted
	}
	if err := db.UpdateFetchTask(*task); err != nil {
		return err
	}
	return nil
}
