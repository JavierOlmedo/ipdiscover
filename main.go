package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {

	var all bool
	var threads int
	flag.BoolVar(&all, "a", false, "Get all IPs found")
	flag.IntVar(&threads, "t", 23, "Set numbers of threads")// ¯\_(ツ)_/¯
	flag.Parse()

	var domains io.Reader
	domains = os.Stdin

	domain := flag.Arg(0)
	if domain != "" {
		domains = strings.NewReader(domain)
	}

	out := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)

		go func() {

			for u := range out {
				ip := getIP(u)

				if len(ip) == 0 {
					fmt.Println(strings.Join([]string{u, ";", "Unknown"}, ""))
					continue
				}

				if all {
					for _, i := range ip {
						fmt.Println(strings.Join([]string{u, ";", i}, ""))
					}

				} else {
					fmt.Println(strings.Join([]string{u, ";", ip[0:1][0]}, ""))
				}
			}

			wg.Done()
		}()
	}

	sc := bufio.NewScanner(domains)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		out <- domain
	}

	close(out)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	wg.Wait()
}

func getIP(domain string) []string {
	var array []string

	addrs, err := net.LookupIP(domain)
	if err != nil {
		return nil
	}

	for _, addr := range addrs {
		array = append(array, addr.String())
	}
	return array
}
