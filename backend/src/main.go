package main

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	cors "github.com/itsjamie/gin-cors"
)

var ctx = context.Background()

const REFRESH_RATE = 300

func main() {

	rdb := initRedis()

	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json", []byte(getBubbles(rdb)))
	})

	r.Run("0.0.0.0:8888")
}

func getBubbles(rdb *redis.Client) string {

	for {
		var bubbles string = "{"
		keys, cursor, err := rdb.Scan(ctx, 0, "bubble:*", 100).Result()
		if err != nil {
			panic(err)
		}
		last_i := len(keys) - 1
		for i, key := range keys {

			form, err := rdb.Get(ctx, key).Result()
			if err != nil {
				panic(err)
			}
			//fmt.Println(string(form))
			bubbles += "\"" + key + "\":" + string(form)
			if i < last_i {
				bubbles += ","
			}
		}
		if cursor == 0 {
			return bubbles + "}"
		}
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
	form.Data.Color = "blue"

	return form
}

func initRedis() *redis.Client {
	// connexion
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb

	//listen for events on new connexion
	// notify-keyspace-events g$xE
	// subscribe to __keyevent@0__:set

}
