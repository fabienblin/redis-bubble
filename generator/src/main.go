package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

const MAX_SLEEP = 1
const MAX_TTL = 30

func main() {
	rand.Seed(time.Now().UnixNano())
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	id := 1
	for {
		form := newCircleForm()
		jsonForm, err := json.Marshal(form)
		if err != nil {
			panic(err)
		}

		// send data to redis
		key := "bubble:" + form.Type + ":" + strconv.Itoa(id)
		err = rdb.SetEx(ctx, key, string(jsonForm), time.Duration(rand.Intn(MAX_TTL)+1)*time.Second).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println(key)
		fmt.Println(string(jsonForm))

		id++
		time.Sleep(time.Duration(rand.Intn(MAX_SLEEP)+1) * time.Second)
	}
}

type Form struct {
	Type string      `json:"type"`
	Data *CircleForm `json:"data"`
}

type CircleForm struct {
	Cx    int    `json:"cx"`
	Cy    int    `json:"cy"`
	R     int    `json:"r"`
	Color string `json:"color"`
}

func newCircleForm() *Form {
	form := new(Form)
	form.Type = "circle"
	form.Data = new(CircleForm)
	form.Data.Cx = rand.Intn(16)
	form.Data.Cy = rand.Intn(9)
	form.Data.R = rand.Intn(300)
	form.Data.Color = "hsla(" + strconv.Itoa(rand.Intn(360)) + ", 100%, 50%, " + strconv.FormatFloat(rand.Float64(), 'f', 2, 64) + ")"

	return form
}
