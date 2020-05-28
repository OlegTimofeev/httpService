package dataBase

type ConfigDB struct {
	User          string
	Password      string
	Dbname        string
	StoreType     string
	PoolSize      int
	NatsUrl       string
	StanClusterID string
}
