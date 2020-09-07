package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port      string
	TCP       string
	TCPFilter bool
	UDP       string
	UDPFilter bool
}

// Scan port
func ScanPort(hostname string, port int) *ScanResult {
	address := hostname + ":" + strconv.Itoa(port)
	result := &ScanResult{
		Port:      strconv.Itoa(port),
		TCP:       "Closed",
		TCPFilter: false,
		UDP:       "Closed",
		UDPFilter: false,
	}

	// TCP
	connTCP, err := net.DialTimeout("tcp", address, 60*time.Second)
	if err == nil {
		result.TCP = "Open"
		result.TCPFilter = true
		defer connTCP.Close()
	}
	fmt.Printf("1")

	// UDP
	connUDP, err := net.DialTimeout("udp", address, 60*time.Second)
	if err == nil {
		result.UDP = "Open"
		result.UDPFilter = true
		defer connUDP.Close()
	}

	return result
}

// Loop through range
func InitialScan(hostname string, low, high int) []ScanResult {
	var wg *sync.WaitGroup
	wg.Add(high)

	var results []ScanResult
	for i := low; i <= high; i++ {
		fmt.Printf("0") // DEBUG
		go func(hostname string, port int) {
			defer wg.Done()
			scan := ScanPort(hostname, port)
			if scan.TCPFilter || scan.UDPFilter {
				results = append(results, *scan)
			}

			fmt.Printf("1") // DEBUG
		}(hostname, i)
	}

	wg.Wait()
	return results
}

func main() {
	fmt.Println("Port Scanner (tcp/udp):")

	low, high := 1, 1024

	results := InitialScan("localhost", low, high)
	for _, v := range results {
		fmt.Printf("port: %s | tcp: %s | udp: %s\n", v.Port, v.TCP, v.UDP)
	}
}
