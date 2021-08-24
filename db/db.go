package db

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type Sites struct {
	Duration time.Duration
	Url      string
}

var siteBucket = []byte("sites")
var db *bolt.DB

func Init(path string) []byte {
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists(siteBucket)
		return err
	})
	return siteBucket
}

func AddSite(url string, d time.Duration) (time.Duration, error) {

	err := db.Update(func(t *bolt.Tx) error {
		b := t.Bucket(siteBucket)
		Key := dtb(d)
		err := b.Put(Key, []byte(url))
		return err
	})

	if err != nil {
		return -1, err
	}

	return d, nil
}

func AllSites() ([]Sites, error) {
	var sites []Sites
	err := db.View(func(t *bolt.Tx) error {
		b := t.Bucket(siteBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			sites = append(sites, Sites{
				Duration: btd(k),
				Url:      string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return sites, nil
}

func DeleteSite(key time.Duration) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(siteBucket)
		return b.Delete(dtb(key))
	})
}

// function to change duration into byte of string format
func dtb(d time.Duration) []byte {
	return []byte(d.String())
}

//Function to change byte of string format into time.Duration
func btd(x []byte) time.Duration {
	v := string(x)
	fmt.Println(v)
	u, _ := time.ParseDuration(v)
	return u
}
