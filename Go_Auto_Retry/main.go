package main

import (
	"errors"
	"fmt"
	"github.com/adityarama1210/retrier"
	"log"
	"time"
)

type UserActivitySaver struct {
	UserEmail string
	Activity  string
}

func (u *UserActivitySaver) Exec() error {
	err := SaveToDB(u.UserEmail, u.Activity)
	if err != nil {
		return err
	}
	return nil
}

//goland:noinspection GoErrorStringFormat
func SaveToDB(email, activity string) error {
	fmt.Println("SAVING to DB for email", email, " with Activity", activity)
	time.Sleep(100 * time.Millisecond)
	if email == "user@test.com" {
		return errors.New("Mocking an error.")
	}
	return nil
}

func main() {
	uas := &UserActivitySaver{
		UserEmail: "john@gmail.com",
		Activity:  "Login Success",
	}
	r, err := retrier.New(uas, 5, func(err error) { log.Println("User Activity Saver error", err) })
	if err != nil {
		log.Fatal("ERR", err)
	}
	r.Start()
	fmt.Println("Start Invocation Done")
	time.Sleep(2 * time.Second)
	fmt.Println("DONE")
}
