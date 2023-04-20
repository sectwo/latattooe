package main

import (
	"fmt"

	"DMS_v1.1/utils"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName        = "lattatoe"
	barcodeBucket = "barcode"
	idBucket      = "id"
	checkpoint    = "checkpoint"
)

var db *bolt.DB

func getDbName() string {
	//port := os.Args[2][6:]
	return fmt.Sprintf("%s.db", dbName)
}

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(getDbName(), 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(barcodeBucket))
			utils.HandleErr(err)
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func Close() {
	DB().Close()
}

func SaveBarcode(id string, barcode string) {
	//fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(barcodeBucket))
		err := bucket.Put([]byte(id), []byte(barcode))
		return err
	})
	utils.HandleErr(err)
}

func viewBarcode(id string) string {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(barcodeBucket))
		data = bucket.Get([]byte(id))
		return nil
	})
	return string(data)
}

func viewAllBarcodeInfo() map[string]string {
	roll := make(map[string]string)
	// 모든 버킷 목록 출력
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(barcodeBucket))
		bucket.ForEach(func(k, v []byte) error {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			roll[string(k)] = string(v)
			return nil
		})
		return nil
	})
	return roll
}
