package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	fetchURL   = "http://jsonplaceholder.typicode.com/posts"
	cloudStore = "cloud_store.txt"
)

type userInfo struct {
	UserId     int
	Id         int
	Title      string
	Body       string
	IngestedAt time.Time `json:"ingested_at"`
	Source     string    `json:"source"`
}

var ch = make(chan userInfo)
var wg sync.WaitGroup

func startDataTransformation(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		cmd := req.URL.Query().Get("cmd")
		if strings.Compare(cmd, "start") == 0 {
			users, err := fetch()
			if err != nil {
				panic(err)
			}
			wg.Add(3)
			go tranform(users)
			go store()
			go func() {
				defer wg.Done()
				wg.Wait()
				close(ch)
			}()

			res.WriteHeader(201)
			fmt.Fprintf(res, "Successfully collected data in %s", cloudStore)
		}
	}
}

// 1. Fetch data
func fetch() ([]userInfo, error) {
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

// 2. Tranform data
func tranform(users []userInfo) {
	defer wg.Done()
	for _, user := range users {
		u := userInfo{user.UserId, user.Id, user.Title, user.Body, time.Now(), "placeholder_api"}
		ch <- u
	}
}

// 3. Store to cloud storage (don't have personal cloud account)
func store() {
	defer wg.Done()
	f, err := os.Create(cloudStore)
	if err != nil {
		panic("store creation failed")
	}

	for u := range ch {
		record := fmt.Sprintf("\n{user_id:%d, id:%d, title:%s, body:%s,  ingested_at:%v, source: %s}\n", u.UserId, u.Id, u.Title, u.Body, u.IngestedAt, u.Source)
		f.WriteString(record)
	}
	f.Close()
}

func main() {
	http.HandleFunc("/datatx", startDataTransformation)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
