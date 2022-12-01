# Redis Real Time Bubbles

This project is a first training on Redis using Golang servers and Vuejs front-end to showcase real time data on a web app.


# Generator

The Golang generator will provide ephemeral data to the Redis database. This data is stored as JSON with the corresponding attributes for SVG circles. The circle data is generated at random time interval, with random values and a random TTL (expiration).
Created object keys are formated like this : bubble:circle:$id

### Run
`go run src/main.go`

# Redis Server

The Redis server is launched as a service on localhost:6379 without credentials.

### Run
`redis-server`


# Backend

The backend is a Golang HTTP server on localhost:8888.  It will scan all keys starting with "bubble:"  each time a GET request is sent on its root URL.

### Run
`go run src/main.go`

# Frontend

The frontend is a Vuejs server that will request the backend every 300 milliseconds to GET the state of the Redis server. The JSON is parsed in SVG and updated on the browser.

### Run
`npm run serve`