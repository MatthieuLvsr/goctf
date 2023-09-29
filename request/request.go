package request

import (
	// "crypto/tls"
	// "bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	// "os"
	"sync"
	"time"
)

const(
	CLIENT = "http://10.49.122.144:"
	TRIES = 60000
	TIMEOUT = 1 * time.Second
	ROUTINES = 1000
)

func Request() {
	client := &http.Client{
		Timeout: TIMEOUT,
	}

	var wg sync.WaitGroup
	ports := make(chan int,ROUTINES)
	for i := 0; i < ROUTINES; i++{
		wg.Add(1)

		go func(){
			defer wg.Done()
			for port := range ports{
				resp, err := client.Get(fmt.Sprintf("%s%d/ping",CLIENT,port))
				if err == nil && resp.StatusCode == http.StatusOK {
					defer resp.Body.Close()
					rbody, _ := io.ReadAll(resp.Body)
					fmt.Printf("port : %d -> %s\n",port,string(rbody))
					body := []byte(`{
						"User": "Matt"
					}`)
					ReqPost(port, "signup",body)
					ReqPost(port, "check",body)
					body = []byte(`{
						"User": "Matt",
						"Secret": "admin"
					}`)
					ReqPost(port, "getUserLevel",body)
					ReqPost(port, "getUserPoints",body)
					// founder(port)
					return
				}
			}
		}()
	}
	for port := 1; port <= TRIES; port ++{
		ports <- port
	}
	close(ports)
	wg.Wait()
}

func ReqPost(port int, path string, body []byte){
	
	r, err := http.NewRequest("POST", fmt.Sprintf("%s%d/%s",CLIENT,port,path), bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}


	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: TIMEOUT,
	}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	rbody, _ := io.ReadAll(res.Body)
	fmt.Println(string(rbody))
}

// func founder(port int){
// 	file, err := os.Open("./able.txt")

// 	if err != nil {
// 		return
// 	}

// 	fileScanner := bufio.NewScanner(file)

// 	for fileScanner.Scan(){
// 		ReqPost(port,fileScanner.Text())
// 	}

// 	if err := fileScanner.Err(); err != nil {
// 		return
// 	}

//     file.Close()
// }