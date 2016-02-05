package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type Id struct {
	timer, counter int64
	mutex          sync.Mutex
}

func (id *Id) Get() int64 {
	if id.timer != time.Now().Unix() {
		id.timer = time.Now().Unix()
		id.counter = 0
	}

	id.mutex.Lock()
	id.counter++
	newId := (id.timer << 32) + id.counter
	id.mutex.Unlock()

	return newId
}

var id Id

func main() {
	id = Id{}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/id", generateId)

	router.Run("127.0.0.1:8080")
}

func generateId(c *gin.Context) {
	c.String(200, "%d\n", id.Get())
}
