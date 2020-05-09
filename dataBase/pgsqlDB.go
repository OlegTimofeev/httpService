package dataBase

import (
	"errors"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"httpService/models"
	"log"
)

var postgres *pg.DB
var pgOptions = pg.Options{User: user, Password: password, Database: dbname}

func connect() *pg.DB {
	postgres = pg.Connect(&pgOptions)
	return postgres
}

func (db *PostgresDB) InitDB() {
	db.pgdb = connect()
	err := db.pgdb.RunInTransaction(func(tx *pg.Tx) error {
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
		if cons := db.pgdb.PoolStats().Hits; cons < 1 {
			return *new(error)
		}
		return nil
	})
	log.Fatal(err)
}

func (db *PostgresDB) CheckConnection() error {
	if cons := db.pgdb.PoolStats().Hits; cons < 1 {
		return errors.New("")
	}
	return nil
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
