package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Request struct {
	Num  int
	Resp chan Response
}

type Response struct {
	Num      int
	WorkerID int
}

func PlusOneService(reqs <-chan Request, workerId int) {
	fmt.Println("PlusOneService Start : " + time.Now().Format("YYYY MM DD HH24:MI:SS"))
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerId}
			fmt.Println("PlusOneService Start - go func : " + time.Now().Format("YYYY MM DD HH24:MI:SS") + " Value : " + strconv.Itoa(req.Num))
		}(req)
	}
}

func MappingRequestAndResponse() {
	fmt.Println("Start -------------- MappingRequestAndResponse --------------")

	reqs := make(chan Request)
	defer close(reqs)
	for i := 0; i < 3; i++ {
		go PlusOneService(reqs, i)
	}
	var wg sync.WaitGroup
	for i := 3; i < 53; i += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resps := make(chan Response)
			reqs <- Request{i, resps}
			fmt.Println(i, "=>", <-resps)
		}(i)
	}
	wg.Wait()
}
