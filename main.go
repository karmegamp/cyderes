package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

const fetchURL = "http://jsonplaceholder.typicode.com/posts"

type userInfo struct {
	UserId     int
	Id         int
	Title      string
	Body       string
	IngestedAt time.Time `json:"ingested_at"`
	Source     string    `json:"source"`
}

func main() {

	fetch := func() ([]userInfo, error) {
		res, err := http.Get(fetchURL)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		var users []userInfo
		if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
			return nil, err
		}
		return users, nil
	}

	// 1. Collect data
	users, err := fetch()
	if err != nil {
		panic(err)
	}

	ch := make(chan userInfo)
	var wg sync.WaitGroup

	// 2. Tranform data
	tranform := func() {
		defer wg.Done()
		for _, user := range users {
			u := userInfo{user.UserId, user.Id, user.Title, user.Body, time.Now(), "placeholder_api"}
			ch <- u
		}
		close(ch)
	}
	wg.Add(1)
	go tranform()

	// 3. Store to cloud storage (don't have cloud account yet)
	store := func() {
		defer wg.Done()
		f, err := os.Create("./cloud_store.txt")
		if err != nil {
			panic(err)
		}
		for u := range ch {
			record := fmt.Sprintf("\n{user_id:%d, id:%d, title:%s, body:%s,  ingested_at:%v, source: %s}\n", u.UserId, u.Id, u.Title, u.Body, u.IngestedAt, u.Source)
			f.WriteString(record)
		}
		f.Close()
	}
	wg.Add(1)
	go store()

	wg.Wait()
}
