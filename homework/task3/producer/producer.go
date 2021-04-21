package main

import (
	"encoding/binary"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/nsqio/go-nsq"
)

// 产生随机数
func generateNumber() float64 {
	return rand.Float64()
}

func main() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	topicName := "topic"

	doChan := make(chan *nsq.ProducerTransaction)

	go func() {
		for {
			select {
			case res := <-doChan:
				if res.Error != nil {
					log.Println("send error:", res.Error.Error())
				}
			}
		}
	}()

	// 发送1000个随机数
	for i := 0; i < 1000; i++ {
		num := generateNumber()
		bits := math.Float64bits(num)
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, bits)
		// err = producer.Publish(topicName, bytes)
		err := producer.PublishAsync(topicName, bytes, doChan)
		if err != nil {
			log.Fatal("could not pulish:", err)
		}

		// 休眠10毫秒
		time.Sleep(time.Millisecond * 10)
	}

	producer.Stop()
}
