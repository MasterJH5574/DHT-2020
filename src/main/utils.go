package main

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/fatih/color"
)

const (
	firstPort int = 20000

	lengthOfKeyValue int = 50

	afterTestSleepTime = 30 * time.Second

	basicTestRoundNum               int     = 5   // 5 rounds in total.
	basicTestNodeSize               int     = 100 // Nodes are numbered with 0 ~ 100 (101 nodes in total).
	basicTestRoundJoinNodeSize      int     = 20
	basicTestRoundQuitNodeSize      int     = 10
	basicTestRoundPutSize           int     = 150
	basicTestRoundGetSize           int     = 120
	basicTestRoundDeleteSize        int     = 70
	basicTestMaxFailRate            float64 = 0.01
	basicTestAfterRunSleepTime              = 200 * time.Millisecond
	basicTestJoinQuitSleepTime              = time.Second
	basicTestAfterJoinQuitSleepTime         = 10 * time.Second

	forceQuitNodeSize           int     = 50
	forceQuitPutSize            int     = 500
	forceQuitRoundNum           int     = 10
	forceQuitMaxFailRate        float64 = 0.15
	forceQuitRoundQuitNodeSize          = forceQuitNodeSize / forceQuitRoundNum
	forceQuitAfterRunSleepTime          = 200 * time.Millisecond
	forceQuitJoinSleepTime              = time.Second
	forceQuitAfterJoinSleepTime         = 10 * time.Second
	forceQuitFQSleepTime                = 500 * time.Millisecond

	QASNodeSize           int     = 50
	QASPutSize            int     = 500
	QASMaxFailRate        float64 = 0.01
	QASGetSize            int     = 20
	QASAfterRunSleepTime          = 200 * time.Millisecond
	QASJoinSleepTime              = time.Second
	QASAfterJoinSleepTime         = 10 * time.Second
	QASQuitSleepTime              = 80 * time.Millisecond
)

var (
	green  = color.New(color.FgGreen)
	red    = color.New(color.FgRed)
	yellow = color.New(color.FgYellow)
	cyan   = color.New(color.FgCyan)
)

var (
	localAddress string
)

var (
	wg *sync.WaitGroup
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	localAddress = GetLocalAddress()
}

// function to get local address(ip address)
func GetLocalAddress() string {
	var localaddress string

	ifaces, err := net.Interfaces()
	if err != nil {
		panic("init: failed to find network interfaces")
	}

	// find the first non-loopback interface with an IP address
	for _, elt := range ifaces {
		if elt.Flags&net.FlagLoopback == 0 && elt.Flags&net.FlagUp != 0 {
			addrs, err := elt.Addrs()
			if err != nil {
				panic("init: failed to get addresses for network interface")
			}

			for _, addr := range addrs {
				ipnet, ok := addr.(*net.IPNet)
				if ok {
					if ip4 := ipnet.IP.To4(); len(ip4) == net.IPv4len {
						localaddress = ip4.String()
						break
					}
				}
			}
		}
	}
	if localaddress == "" {
		panic("init: failed to find non-loopback interface with valid address on this node")
	}

	return localaddress
}

func randString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func portToAddr(ip string, port int) string {
	return fmt.Sprintf("%s:%d", ip, port)
}

func removeFromArray(s []int, idx int) []int {
	s[len(s)-1], s[idx] = s[idx], s[len(s)-1]
	return s[:len(s)-1]
}

/* ------ Struct "testInfo" ------ */
type testInfo struct {
	msg       string
	failedCnt int
	totalCnt  int
}

func (info *testInfo) success() {
	info.totalCnt++
}

func (info *testInfo) fail() {
	info.totalCnt++
	info.failedCnt++
}

func (info *testInfo) finish(failedCnt *int, totalCnt *int) {
	*failedCnt += info.failedCnt
	*totalCnt += info.totalCnt
	info.printInfo()
}

func (info *testInfo) printInfo() {
	if info.failedCnt > 0 {
		_, _ = red.Printf("%s failed with error rate %.4f\n", info.msg,
			float64(info.failedCnt)/float64(info.totalCnt))
	} else {
		_, _ = green.Printf("%s passed.\n", info.msg)
	}
}
