package cbmiddle

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	gobreaker "github.com/mauryparra/go-circuit-breaker/src/go-circuit-breaker/gocb"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.Interval = time.Second * 60
	st.Timeout = time.Second * 20
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.TotalFailures >= 3 && failureRatio >= 0.4
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		fmt.Println("Realizando request a: " + url)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode >= 500 {
			fmt.Println("Response Code (INVALID): " + strconv.Itoa(resp.StatusCode))
			fmt.Println("Response Body: " + string(data))
			return nil, errors.New("La API no respondio como se esperaba")
		}

		fmt.Println("Response Code (VALID): " + strconv.Itoa(resp.StatusCode))
		fmt.Println("Response Body: " + string(data))
		return data, nil
	})

	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

// Cb middleware with circuit breaker
func Cb(context *gin.Context) {

	fmt.Println("")
	fmt.Println("+++++++++++++++++++++++++++++++++")

	url := context.Query("req")

	body, err := Get(url)
	if err != nil {
		fmt.Println(err.Error())
		context.String(200, "La API no respondio como se esperaba")
		fmt.Println("+++++++++++++++++++++++++++++++++")
		fmt.Println("")
		return
	}

	fmt.Println("La API respondio como se esperaba")
	fmt.Println("+++++++++++++++++++++++++++++++++")
	fmt.Println("")

	context.String(200, string(body))
}
