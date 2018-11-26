package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	// LockTestTheadResult defines result for each thread form lock test
	LockTestTheadResult struct {
		Message  string  `json:"message"`
		FailTime float64 `json:"fail_time"`
	}

	// LockTestResult defines lock test results
	LockTestResult struct {
		Status        string                `json:"status"`
		UsedResource  string                `json:"used_resource"`
		KeepAlive     bool                  `json:"keep_alive"`
		TimeToLive    int                   `json:"time_to_live,omitempty"`
		TimeSpent     time.Time             `json:"time_spent,omitempty"`
		ThreadResults []LockTestTheadResult `json:"thread_results"`
	}

	// LockTestParams defines incomming request params for lock testing
	LockTestParams struct {
		Resource           string `json:"resource"`
		TimeToLive         int    `json:"time_to_live"`
		LockedTestInterval int    `json:"locked_test_interval"`
		KeepAlive          bool   `json:"keep_alive"`
		KeepAliveTime      int    `json:"keep_alive_time"`
		ThreadSize         int    `json:"thread_size"`
		TreadStartInterval int    `json:"thread_start_interval"`
	}

	CreateParams struct {
		Username     string   `json:"username"`
		Guests       []Person `json:"guests"`
		RoomType     string   `json:"roomType"`
		CheckinDate  string   `json:"checkinDate"`
		CheckoutDate string   `json:"checkoutDate"`
	}

	Person struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	TestParams struct {
		Resource   string `json:"resource"`
		TimeToLive int    `json:"time_to_live"`
		KeepAlive  bool   `json:"keep_alive"`
		ThreadSize int    `json:"thread_size"`
	}
)

var startedTime time.Time

//var timeToLive  int

func lock(rs string, ttl int) bool {
	if time.Time.IsZero(startedTime) {
		startedTime = time.Now()
		return true
	}
	sinceSecs := int(time.Since(startedTime).Seconds())
	if sinceSecs <= ttl {
		return false
	}
	startedTime = time.Now()
	return true

}

func execKeepAlive(times int, keepAliveTime int, start *time.Time, checkAliveInterval int) {
	for i := 1; i <= times; i++ {

		for {
			cai := time.Duration(checkAliveInterval)
			time.Sleep(cai * time.Millisecond)
			sinceSecs := int(time.Since(*start).Seconds())
			if sinceSecs >= keepAliveTime {
				*start = time.Now()
				output := fmt.Sprintf("%s : %d", "lock reset at", sinceSecs)
				fmt.Println(output)
				break
			}
		}

	}
}

func lockTest(start *time.Time, wg *sync.WaitGroup, maxSecs int, LockedTestInterval int, ltr *[]LockTestTheadResult) {

	for {

		//test the lock

		milliSecs := time.Duration(LockedTestInterval)
		time.Sleep(milliSecs * time.Millisecond)
		sinceSecs := int(time.Since(*start).Seconds())
		sinceMilli := time.Since(*start).Seconds()
		if sinceSecs >= maxSecs {
			output := fmt.Sprintf("%s : %d seconds", "lock fail at", sinceSecs)
			fmt.Println(output)
			*ltr = append(*ltr, LockTestTheadResult{Message: "Lock test fail", FailTime: sinceMilli})
			wg.Done()
			return
		}
	}

}

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		//var createParams TestParams

		/* fmt.Println(lock("", 5))

		for i := 0; i <= 5; i++ {
			fmt.Println(lock("", 5))
			time.Sleep(2 * time.Second)
		}
		fmt.Println("TERMINA") */

		routinesSize := 1
		maxSecs := 2
		lockTestInterval := 10
		threadInterval := 100
		keepAliveTime := 2
		keepAliveLoops := 2
		checkAliveInterval := 100
		keepAlive := false

		mes := &struct {
			Resource               string `json:"resource" binding:"exists,alphanum"`
			TimeToLive             int    `json:"time_to_live" binding:"required"`
			LockedTestInterval     int    `json:"locked_test_interval"`
			ThreadSize             int    `json:"thread_size"`
			TreadsStartInterval    int    `json:"threads_start_interval"`
			KeepAlive              bool   `json:"keep_alive"`
			KeepAliveTime          int    `json:"keep_alive_time"`
			KeepAliveLoops         int    `json:"keep_alive_loops"`
			KeepAliveCheckInterval int    `json:"keep_alive_check_interval"`
		}{}

		err := c.BindJSON(mes)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
			//log.Fatal(err)
		}

		// validations
		if mes.Resource == "" {
			mes.Resource = "stringuete"

		}

		fmt.Println(mes.Resource)
		fmt.Println(mes.TimeToLive)
		fmt.Println(mes.LockedTestInterval)
		fmt.Println(mes.ThreadSize)
		fmt.Println(mes.TreadsStartInterval)
		fmt.Println(mes.KeepAlive)
		fmt.Println(mes.KeepAliveTime)
		fmt.Println(mes.KeepAliveLoops)
		fmt.Println(mes.KeepAliveCheckInterval)

		var countDown sync.WaitGroup
		countDown.Add(routinesSize)
		threadResults := []LockTestTheadResult{}

		start := time.Now()
		for i := 1; i <= routinesSize; i++ {
			output := fmt.Sprintf("%s : %d", "started routin:", i)
			fmt.Println(output)
			go lockTest(&start, &countDown, maxSecs, lockTestInterval, &threadResults)
			ti := time.Duration(threadInterval)
			time.Sleep(ti * time.Millisecond)
		}

		if keepAlive {
			go execKeepAlive(keepAliveLoops, keepAliveTime, &start, checkAliveInterval)
		}

		countDown.Wait()

		usi := LockTestResult{
			Status:        "status",
			UsedResource:  "un string",
			KeepAlive:     true,
			TimeToLive:    10,
			TimeSpent:     time.Now(),
			ThreadResults: threadResults,
		}

		c.JSON(http.StatusOK, usi)

	})

	r.Run(":8000")

	// request
	// curl http://localhost:8000 -d @request.json
}
