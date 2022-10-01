# Goroutine

go를 붙여서 함수를 호출하게 되면 goroutinetest 호출과 현재 함수의 흐름은 메모리를 고유하는 논리적으로 별개의 흐름이됩니다.  
여기서 논리적으로 별개의 흐름이라고 한 이유는 물리적으로 별개의 흐름이 되는 것과는 구분 되기 때문입니다. 

```go
package main

import "fmt"

func main() {
	fmt.Println("I can get anything if I try to get it with all my effort")
    go goroutinetest(1,2,3)
	fmt.Println("All things happen whether you want or not")
}

func goroutinetest(x, y, z int) {
	fmt.Println("%s, %s, %s", x, y, z)
}
```

### 병렬성과 병행성 

커피를 마시면서 신문을 보고 있는 사람이 있다면 물리적으로 두 흐름이 동시에 수행되는 것은 아닙니다. 커피를 마시기 위하여 신문을 보는 것을 짧은 시간 
잠시 중단하고 커피를 한 모금 마신뒤에 다시 신문 보는 일로 돌아오는 경우는 물리적으로 두 흐름이 있지는 않지만, 동시에 두 가지를 하고는 있습니다. 
이를 동시성 혹은 병행성이라고 합니다. 물리적으로 동시에 수행되지는 않지만 순차적으로 수행되는 것과는 큰 차이가 있습니다. 즉, 커피 한잔을 모두 다 마시는
동안 신문을 전혀 보지 않고, 반드시 커피를 다 마시고 난 다음에 신문을 보기 시작해야 한다는 것이 아닙니다. 이 과정에서 신문 기사의 다섯번째 문단을 읽는 
것 중에서 어느 것을 먼저 해야하는지는 그다지 상관 없는 것이 됩니다.  

### 동시성과 병렬성은 다르지만, 

동시성이 있어야 병렬성이 생기게 됩니다. 서로 어느것이 먼저 되어야하는 의존 관계가 있는 것은 함께 진행될 수 없지요. 장갑을 끼면서 동시에 구두를 신을 수는 있겠지만, 
양말을 신으면서 동시에 구두를 신을 수는 없습니다. 양말을 신고 나야 구두를 신을 수 있으니까요, 이 둘 사이에 동시성이 없으므로 병렬성이 생기지 않는 것입니다. 
즉, 장값을 먼저 끼거나 구두를 먼저 신는 것은 어느 것을 먼저해도 상관이 없지만, 양말은 구두를 신기전에 신어야 합니다. 

### Main Goroutine는 내부의 서브 Goroutine의 처리 시간을 보장하지 않습니다.

아래의 코드에서 main 함수 내부의 sub goroutine의 처리시간을 보장하지 않습니다. main goroutine이 종료되면 서브는 강제적으로 종료됩니다.

```go
package main

import "fmt"

func main() {
	go func() {
		fmt.Println("In goroutine")
	}()
	fmt.Println("In main routine")
}
```

**그러면 어찌해야 할까요?**

### 고루틴 기다리기 

실제 코드로 한번 볼까요?

```go
package main 

func main() {
	urls := []string{
		"http://image.com/img01.jpg",
		"http://image.com/img02.jpg",
		"http://image.com/img03.jpg",
	}

	for _, url := range urls {
		go func(url string) {
			if _, err := Download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}

	filenames, err := filepath.Glob("*.jpg")

	if err != nil {
		log.Fatal(err)
	}

	err = WriteZip("images.zip", filenames)

	if err != nil {
		log.Fatal(err)
	}
}
```

위의 프로그램을 실행하면 어떻게 될까요? 원하는 대로 압축이 될 까요? ...?  
파일이 다운로드 되기도 전에 압축하려고 들 수 있기 때문에 압축이 되지 않습니다. 저 위의 코드에 따르면 각각의 파일을 다운로드 하는 것과 압축하는 것 중 
어느 것을 먼저 수행하여도 좋은 동시성있는 작업들입니다. 그러나 실제로는 그렇지 않습니다. 파일 다운로드가 완료되지 않으면 압축도 할수 없습니다.  
즉, 압축 작업은 파일들이 모두 다운로드될 때까지 기다린후 수행되야 되겠습니다. 즉, 압축 작업은 파일들이 모두 다운로드 될때까지 기다링 후 수행해야 되겠습니다. 
이럴 때 이용할 수 있는 것이 sync.WaitGroup입니다.

```go
package main

import "sync"

func main() {

	var wg sync.WaitGroup
	
	urls := []string{
		"http://image.com/img01.jpg",
		"http://image.com/img02.jpg",
		"http://image.com/img03.jpg",
	}

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			if _, err := Download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	
	wg.Wait()

	filenames, err := filepath.Glob("*.jpg")

	if err != nil {
		log.Fatal(err)
	}

	err = WriteZip("images.zip", filenames)

	if err != nil {
		log.Fatal(err)
	}
}
```

wg 에는 기본 값이 0으로 맞춰져 있는 카운터가 들어가 있습니다. Wait() 함수는 이 카운터가 0이 될때까지 기다립니다. Add() 함수는 호출될 때 마다 숫자를 더합니다. 
wg.Done()은 사실상 wg.Add(-1)과 같다고 보시면 됩니다.   

그렇기 때문에 처음에 WaitGroup을 만들자마자 URL 개수만큼 카운터를 증가시키고 각각의 고루틴에서는 작업이 완료될 때까지 카운터를 감소시킵니다. 모든 고루틴이 끝나면 
카운터가 0이 되는데, 이 상태가 되기전까지는 wg.Wait() 부분에서 멈춰있게 됩니다.

미리 고루틴이 몇 개 생길지 알기 때문에 이렇게 작성이 가능합니다만, 고루틴이 몇개 생길지 알기 어렵거나 따로 수를 세어야 알 수 있는 경우에는 고루틴 띄워보내기 전에 
wg.Add(1)을 수행하여 하나씩 카운터를 증가시킬 수 있습니다.

```go
package main

import "log"

func main() {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()
}
```

### 공유 메모리와 최솟값 찾기 

공유메모리를 이용하여 매우 병렬화가 잘되는 문제를 풀어보기 

```go
package main

import (
	"fmt"
	"sync"
)

func FindMinimumValue() {
	fmt.Println(parallelMin(
		[]int{
			83, 46, 49, 23, 97, 12, 11, 46, 49, 27, 5,
		}, 2))
}

func parallelMin(a []int, n int) int {
	if len(a) < n {
		return min(a)
	}

	mins := make([]int, n)
	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return min(mins)
}

func min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]

	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

```