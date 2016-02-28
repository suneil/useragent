package useragent

import (
	"bufio"
	"math/rand"
	"strings"
	"time"
)

// run 'go generate' whenever useragents.txt is changed
//go:generate go run scripts/includetxt.go

var (
	uaPath        = "ua.txt"
	useragentList []string
)

func init() {
	if len(useragentList) == 0 {
		inFile := strings.NewReader(useragents)
		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			useragentList = append(useragentList, scanner.Text())
		}
	}
}

// GetRandomUserAgent get random user agent string
func GetRandomUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ua := useragentList[r.Intn(len(useragentList))]

	return ua
}
