// utils.go: Proxy loader and utilities
package attack

import (
	"fmt"
	"os"
)

func LoadProxies(filename string) []string {
	fmt.Println("Loading proxies from", filename)
	return []string{"http://1.1.1.1:80", "socks5://2.2.2.2:1080"}
}
