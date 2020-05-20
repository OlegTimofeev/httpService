package dataBase

import (
	"errors"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"httpService/internal/models"
	"log"
)

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
	return nil
}

func (db *PostgresDB) AddTaskResponse(res *models.TaskResponse) error {
	return db.pgdb.RunInTransaction(func(tx *pg.Tx) error {
		tr := models.TaskResponse{FetchTaskID: res.FetchTaskID}
		if err := db.pgdb.Delete(&tr); err != nil {
		}
		if err := db.pgdb.Insert(res); err != nil {
			return err
		}
		return nil
	})
}

func (db *PostgresDB) GetTaskResponseByFtID(taskId int) (*models.TaskResponse, error) {
	tr := models.TaskResponse{FetchTaskID: taskId}
	if err := db.pgdb.Model(&tr).Select(); err != nil {
		return nil, err
	}
	return &tr, nil
}
