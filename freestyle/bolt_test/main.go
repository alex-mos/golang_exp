package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	//"log"
)

type User struct {
	Email string
	Password string
	Age uint8
}

func (user *User) save(db *bolt.DB) error {
	// Store the user model in the user bucket using the username as the key.
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("usersBucket"))
		if err != nil {
			return err
		}

		encoded, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return b.Put([]byte(user.Email), encoded)
	})
	return err
}

func main()  {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		fmt.Println("cant open db")
	}
	defer db.Close()

	alex := User{"alexander.mospan@gmail.com", "qwe123", 31}
	anna := User{"tsatsenkina@gmail.com", "qazwsxedc", 32}

	alex.save(db)
	anna.save(db)
}
