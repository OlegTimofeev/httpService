package nats

import (
	"encoding/json"
	"github.com/itimofeev/go-util"
	"github.com/nats-io/stan.go"
	uuid "github.com/satori/go.uuid"
	"httpService/internal/dataBase"
	"httpService/internal/models"
)

type Queue struct {
	Sc    stan.Conn
	Topic string
}

func NewStan(config dataBase.ConfigDB) *Queue {
	sc, err := stan.Connect(config.StanClusterID, uuid.NewV4().String(), stan.NatsURL(config.NatsUrl),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			util.CheckErr(reason, "stan connect lost")
		}))
	util.CheckErr(err, "stan.Connect")
	return &Queue{Sc: sc, Topic: "123"}
}

func (q *Queue) Publish(ft models.FetchTask) error {
	jsonTask, err := json.Marshal(ft)
	if err != nil {
		return err
	}
	if err := q.Sc.Publish(q.Topic, jsonTask); err != nil {
		return err
	}
	return nil
}

func (q *Queue) Subscribe(handler func(msg *stan.Msg)) {
	q.Sc.Subscribe(q.Topic, handler)
}
