package db

import (
	"log"
	"os/exec"
	"time"

	"github.com/boltdb/bolt"
)

type Sites struct {
	Duration time.Duration
	url      string
}

var siteBucket = []byte("sites")
var db *bolt.DB

func Init(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists(siteBucket)
		return err
	})
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
			//fmt.Printf("key=%s, value=%s\n", k, v)
			sites = append(sites, Sites{
				Duration: btd(k),
				url:      string(v),
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
	var sites []Sites
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(siteBucket)
		time.AfterFunc(key, func() {
			for _, v := range sites {
				cmd := exec.Command("firefox", v.url)

				err := cmd.Run()

				if err != nil {
					log.Fatal(err)
				}
			}
			b.Delete(dtb(key))
		})
		return nil
	})

}

// function to change duration into byte of string format
func dtb(d time.Duration) []byte {
	return []byte(d.String())
}

//Function to change byte of string format into time.Duration
func btd(x []byte) time.Duration {
	v := string(x)
	//fmt.Println(v)
	u, _ := time.ParseDuration(v)
	return u
}
