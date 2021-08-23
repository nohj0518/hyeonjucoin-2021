package db

import (
	"github.com/boltdb/bolt"
	"github.com/nohj0518/hyeonjucoin-2021/utils"
)

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB{
	if db == nil {
		// init DB
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		
		err = db.Update(func(t *bolt.Tx) error {
			_, err:= t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_,err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}