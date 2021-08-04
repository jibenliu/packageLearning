package main

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
	"log"
	"strconv"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Update(func(tx *nutsdb.Tx) error {
		key := []byte("name")
		val := []byte("dj")
		if err := tx.Put("", key, val, 0); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	err = db.View(func(tx *nutsdb.Tx) error {
		key := []byte("name")
		if e, err := tx.Get("", key); err != nil {
			return err
		} else {
			fmt.Println(string(e.Value))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	bucket := "user_list"
	prefix := "user_"

	db.Update(func(tx *nutsdb.Tx) error {
		for i := 1; i <= 300; i++ {
			key := []byte(prefix + strconv.FormatInt(int64(i), 10))
			val := []byte("dj" + strconv.FormatInt(int64(i), 10))
			tx.Put(bucket, key, val, 0)
		}
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {
		entries, _ := tx.PrefixScan(bucket, []byte(prefix), 25)
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		return nil
	})
}
