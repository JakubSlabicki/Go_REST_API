package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type OutputData struct {
	/* This is struct for one result data

	:param Stddev: A standard deviation of all integers from Data array
	:type Stddev: float64
	:param Data: Random integers, lenght of array depends of GET operation
	:type Data: aarray of int
	*/
	Stddev float64 `json:"Stddev"`
	Data   []int   `json:"Data"`
}

type OutputDatas []OutputData

/* This is a array of all OutputData to be shown
 */

func f_mean(data []int) float64 {
	/* Returns mean value of all integers from input array. If input data array is empty then return 0 as mean.

	:param data: array of integers for calculating mean value
	:type data: array of int
	:return: calculated mean value
	:rtype: float64
	*/
	if len(data) == 0 {
		return 0
	}
	var sum int
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	var mean_val float64
	mean_val = float64(float64(sum) / float64(len(data)))
	return mean_val
}

func f_stddev(data []int) float64 {
	/* Returns standard deviation value of all integers from input array. If input data array is empty then return 0 as mean. Using mean function

	:param data: array of integers for calculating standard deviation value value
	:type data: array of int
	:return: calculated standard deviation value
	:rtype: float64
	*/
	if len(data) == 0 {
		return 0
	}
	var _mean = f_mean(data)
	var sum float64
	for i := 0; i < len(data); i++ {
		var diff = float64(data[i]) - _mean
		sum += math.Pow(diff, 2)
	}
	var stddev float64
	stddev = math.Pow(sum/float64(len(data)), 0.5)
	return stddev
}

func homePage(w http.ResponseWriter, r *http.Request) {
	/* Function for printing Hello world string to make sure server and connection is on

	:param w: value of the http type used for sending responses to any connected HTTP clients
	:type w: http.ResponseWriter
	:param r: pointer to an http.Request, a web request
	:type r: *http.Request
	*/
	fmt.Fprint(w, "Hello world!")
}

func randomOutputPage(w http.ResponseWriter, r *http.Request) {
	/* Function for printing request data in json. Function GETs request and lenght data from url query, converts type for integers to calculate standard deviation in two loops. One for creating lenght of array of random integers (1 to 100) and second for creating exact requests number of data structs. At the end function calculate standard deviation for all data.

	:param w: value of the http type used for sending responses to any connected HTTP clients
	:type w: http.ResponseWriter
	:param r: pointer to an http.Request, a web request
	:type r: *http.Request
	*/
	requests := r.URL.Query().Get("requests")
	length := r.URL.Query().Get("length")
	random_data := OutputDatas{}
	loopRequests, err := strconv.Atoi(requests)
	if err != nil {
		fmt.Fprint(w, loopRequests, err, reflect.TypeOf(loopRequests))
	}
	loopLength, err := strconv.Atoi(length)
	if err != nil {
		fmt.Fprint(w, loopLength, err, reflect.TypeOf(loopLength))
	}
	var all_data []int
	var stddev float64
	for i := 1; i <= loopRequests; i++ {
		var _data []int
		for i := 1; i <= loopLength; i++ {
			var _temp_rand = rand.Intn(100)
			_data = append(_data, _temp_rand)
			all_data = append(all_data, _temp_rand)
		}
		stddev = f_stddev(_data)
		random_data = append(random_data, OutputData{Stddev: stddev, Data: _data})

	}
	stddev = f_stddev(all_data)
	random_data = append(random_data, OutputData{Stddev: stddev, Data: all_data})
	json.NewEncoder(w).Encode(random_data)

}

func main() {
	/* Main function.
	Creates flag which waits 15s for which wait for connection to finish. Creates HTTP outer from gorilla/mux repository which is used for building Go web servers. Listen for GET method in which passes requests and lenght values to calculate standard deviation. Setting localhost address, :8081 port and several timeouts and listen for any errors in goroutine. Making graceful shutdown when quit. Creates context cancellation with timeout wait that does not bock if no connecion.
	*/
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "duration time wait")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/random/mean", randomOutputPage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))

	srv := &http.Server{
		Addr:         "127.0.0.0:8081",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
