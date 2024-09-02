package main

import (
	"daily/cmd"
	"log"
)

func main() {
	cmd.Execute()
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
}

// func fetchFirstLine(url string, results chan<- string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	rand := rd.Intn(5000) + 1000
// 	time.Sleep(time.Duration(rand) * time.Millisecond)
// 	results <- fmt.Sprintf("%v ms -> %v", rand, url)
// }

// func main() {
// 	urls := []string{
// 		"https://jsonplaceholder.typicode.com/posts/1",
// 		"https://jsonplaceholder.typicode.com/posts/2",
// 		"https://jsonplaceholder.typicode.com/posts/3",
// 		"https://jsonplaceholder.typicode.com/posts/4",
// 		"https://jsonplaceholder.typicode.com/posts/5",
// 		"https://jsonplaceholder.typicode.com/posts/6",
// 		"https://jsonplaceholder.typicode.com/posts/7",
// 		"https://jsonplaceholder.typicode.com/posts/8",
// 		"https://jsonplaceholder.typicode.com/posts/9",
// 	}

// 	var wg sync.WaitGroup
// 	lines := make(chan string, len(urls))

// 	for _, url := range urls {
// 		wg.Add(1)
// 		go fetchFirstLine(url, lines, &wg)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(lines)
// 	}()

// 	for line := range lines {
// 		fmt.Println(line)
// 	}
// }
