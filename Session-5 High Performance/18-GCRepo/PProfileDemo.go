//go get github.com/pkg/profile
//https://www.cedric-dumont.com/2017/02/01/install-apache-benchmarking-tool-ab-on-windows/
// ab -n 5000 -c 10 http://localhost:1234/
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"
)

func main() {
	Handler := func(w http.ResponseWriter, r *http.Request) {
		sleepDuration := r.URL.Query().Get("time")
		sleepDurationInt, err := strconv.Atoi(sleepDuration)
		if err != nil {
			fmt.Println("Incorrect value passed as a query string for time")
			return
		}
		sleep(sleepDurationInt)
		fmt.Fprintf(w, "Slept for %v Milliseconds", sleepDuration)
	}
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":1234", nil)
}
func sleep(sleepTime int) {
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println("Slept for ", sleepTime, " Milliseconds")
}
//curl -s "localhost:1234/debug/pprof/profile?seconds=10" > out.dump