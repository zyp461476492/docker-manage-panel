package database

import (
	"github.com/asdine/storm"
	"github.com/zyp461476492/docker-app/types"
	"go.etcd.io/bbolt"
	"log"
	"sync"
	"time"
)

var db *storm.DB
var err error
var once sync.Once

func GetStorm(config types.Config) (*storm.DB, error) {
	once.Do(func() {
		db, err = storm.Open(
			config.FileLocation, storm.BoltOptions(0600, &bbolt.Options{Timeout: config.Timeout * time.Second}))
	})
	return db, err
}

func CloseStorm(db *storm.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("关闭数据库失败 %s", err.Error())
	}
}
