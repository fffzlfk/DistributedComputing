package main

import (
	"encoding/binary"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"

	"github.com/fffzlfk/singal/subscriber"
	"github.com/montanaflynn/stats"
	"github.com/nsqio/go-nsq"
)

const N = 100

type myMessageHandler struct {
	nums chan float64
}

func (h *myMessageHandler) processMessage(m []byte) error {

	bits := binary.LittleEndian.Uint64(m)
	num := math.Float64frombits(bits)
	// fmt.Println(num)
	h.nums <- num
	if len(h.nums) == N {
		go func() {
			nums := make([]float64, N)
			for i := range nums {
				nums[i] = <-h.nums
			}

			mean, err := stats.Mean(nums)
			if err != nil {
				log.Fatal("could not calculate the mean of nums:", mean)
			}

			sd, err := stats.StandardDeviation(nums)
			if err != nil {
				log.Fatal("could not calculate the standard deviation of nums:", sd)
			}

			for _, v := range nums {
				if v > mean+3.0*sd || v < mean-3.0*sd {
					log.Println("异常点:", v)
				}
			}
		}()
	}
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

	sub, err := subscriber.NewSubscriber("outlier", &myMessageHandler{make(chan float64, N)}, "127.0.0.1:4161")
	if err != nil {
		log.Fatal("could not create a subscriber:", err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the subscriber.
	sub.Stop()
}
