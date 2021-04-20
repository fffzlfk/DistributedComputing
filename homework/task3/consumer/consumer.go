package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/nsqio/go-nsq"
	chart "github.com/wcharczuk/go-chart/v2"
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
			fmt.Println("the mean: ", mean)
			sd, err := stats.StandardDeviation(nums)
			if err != nil {
				log.Fatal("could not calculate the mean of var:", sd)
			}

			for _, v := range nums {
				if v > mean+2.0*sd || v < mean-2.0*sd {
					log.Println("异常点:", v)
				}
			}

			fmt.Println("the sd: ", sd)
			time.Sleep(time.Second)

			mainSeries := chart.ContinuousSeries{
				Name:    "A test series",
				XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(), //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
				YValues: nums,                                                                                  //generates a []float64 randomly from 0 to 100 with 100 elements.
			}

			linRegSeries := &chart.LinearRegressionSeries{
				InnerSeries: mainSeries,
			}
			graph := chart.Chart{
				Series: []chart.Series{
					mainSeries,
					linRegSeries,
				},
			}
			// 绘制过去一段时间内随机信号的折线图到output.png
			f, _ := os.Create("output.png")
			defer f.Close()
			graph.Render(chart.PNG, f)
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
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&myMessageHandler{make(chan float64, N)})

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}
