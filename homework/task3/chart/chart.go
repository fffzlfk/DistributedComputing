package main

import (
	"encoding/binary"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fffzlfk/singal/subscriber"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
)

const (
	N = 100
)

var globalNums []float64

type myMessageHandler struct{}

func (h *myMessageHandler) processMessage(m []byte) error {
	bits := binary.LittleEndian.Uint64(m)
	num := math.Float64frombits(bits)
	// fmt.Println(num)
	if len(globalNums) > 300 {
		globalNums = globalNums[1:]
	}
	globalNums = append(globalNums, num)

	return nil
}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	err := h.processMessage(m.Body)
	return err
}

func main() {
	// Instantiate a consumer that will subscribe to the provided channel.

	sub, err := subscriber.NewSubscriber("chart", &myMessageHandler{}, "127.0.0.1:4161")
	if err != nil {
		log.Fatal("could not create a subscriber:", err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": globalNums,
		})
	})
	r.Run(":8080")

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the subscriber.
	sub.Stop()
}
