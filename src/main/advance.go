package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func forceQuitTest() (bool, int, int) {
	_, _ = yellow.Println("Start Force Quit Test")

	forceQuitFailedCnt, forceQuitTotalCnt, panicked := 0, 0, false

	defer func() {
		if r := recover(); r != nil {
			_, _ = red.Println("Program panicked with", r)
		}
		panicked = true
	}()

	nodes := new([forceQuitNodeSize + 1]dhtNode)
	nodeAddresses := new([forceQuitNodeSize + 1]string)
	kvMap := make(map[string]string)
	nodesInNetwork := make([]int, 0, basicTestNodeSize+1)

	/* Run all nodes. */
	wg = new(sync.WaitGroup)
	for i := 0; i <= forceQuitNodeSize; i++ {
		nodes[i] = NewNode(firstPort + i)
		nodeAddresses[i] = portToAddr(localAddress, firstPort+i)

		wg.Add(1)
		go nodes[i].Run()
	}
	time.Sleep(forceQuitAfterRunSleepTime)

	/* Node 0 creates a new network. All notes join the network. */
	joinInfo := testInfo{
		msg:       "Force quit join",
		failedCnt: 0,
		totalCnt:  0,
	}
	nodes[0].Create()
	nodesInNetwork = append(nodesInNetwork, 0)
	_, _ = cyan.Printf("Start joining\n")
	for i := 1; i <= forceQuitNodeSize; i++ {
		addr := nodeAddresses[rand.Intn(i)]
		if !nodes[i].Join(addr) {
			joinInfo.fail()
		} else {
			joinInfo.success()
		}
		nodesInNetwork = append(nodesInNetwork, i)

		time.Sleep(forceQuitJoinSleepTime)
	}
	joinInfo.finish(&forceQuitFailedCnt, &forceQuitTotalCnt)

	time.Sleep(forceQuitAfterJoinSleepTime)

	/* Put. */
	putInfo := testInfo{
		msg:       "Force quit put",
		failedCnt: 0,
		totalCnt:  0,
	}
	_, _ = cyan.Printf("Start putting\n")
	for i := 0; i < forceQuitPutSize; i++ {
		key := randString(lengthOfKeyValue)
		value := randString(lengthOfKeyValue)
		kvMap[key] = value

		if !nodes[rand.Intn(forceQuitNodeSize+1)].Put(key, value) {
			putInfo.fail()
		} else {
			putInfo.success()
		}
	}
	putInfo.finish(&forceQuitFailedCnt, &forceQuitTotalCnt)

	/* 10 - 1 = 9 rounds in total. */
	for t := 1; t <= forceQuitRoundNum-1; t++ {
		_, _ = cyan.Printf("Force Quit Round %d\n", t)

		/* Force quit. */
		_, _ = cyan.Printf("Start force quitting (round %d)\n", t)
		for i := 1; i <= forceQuitRoundQuitNodeSize; i++ {
			idxInArray := rand.Intn(len(nodesInNetwork))

			nodes[nodesInNetwork[idxInArray]].ForceQuit()
			nodesInNetwork = removeFromArray(nodesInNetwork, idxInArray)

			time.Sleep(forceQuitFQSleepTime)
		}

		/* Get all data. */
		getInfo := testInfo{
			msg:       fmt.Sprintf("Get (round %d)", t),
			failedCnt: 0,
			totalCnt:  0,
		}
		_, _ = cyan.Printf("Start getting (round %d)\n", t)
		for key, value := range kvMap {
			ok, res := nodes[nodesInNetwork[rand.Intn(len(nodesInNetwork))]].Get(key)
			if !ok || res != value {
				getInfo.fail()
			} else {
				getInfo.success()
			}
		}
		getInfo.finish(&forceQuitFailedCnt, &forceQuitTotalCnt)
	}

	/* All nodes quit. */
	for i := 0; i <= forceQuitNodeSize; i++ {
		nodes[i].Quit()
	}

	return panicked, forceQuitFailedCnt, forceQuitTotalCnt
}

func quitAndStabilizeTest() (bool, int, int) {
	_, _ = yellow.Println("Start Quit & Stabilize Test")

	QASFailedCnt, QASTotalCnt, panicked := 0, 0, false

	defer func() {
		if r := recover(); r != nil {
			_, _ = red.Println("Program panicked with", r)
		}
		panicked = true
	}()

	nodes := new([QASNodeSize + 1]dhtNode)
	nodeAddresses := new([QASNodeSize + 1]string)
	kvMap := make(map[string]string)
	nodesInNetwork := make([]int, 0, QASNodeSize+1)

	/* Run all nodes. */
	wg = new(sync.WaitGroup)
	for i := 0; i <= QASNodeSize; i++ {
		nodes[i] = NewNode(firstPort + i)
		nodeAddresses[i] = portToAddr(localAddress, firstPort+i)

		wg.Add(1)
		go nodes[i].Run()
	}
	time.Sleep(QASAfterRunSleepTime)

	/* Node 0 creates a new network. All notes join the network. */
	joinInfo := testInfo{
		msg:       "Quit & Stabilize join",
		failedCnt: 0,
		totalCnt:  0,
	}
	nodes[0].Create()
	nodesInNetwork = append(nodesInNetwork, 0)
	_, _ = cyan.Printf("Start joining\n")
	for i := 1; i <= QASNodeSize; i++ {
		addr := nodeAddresses[rand.Intn(i)]
		if !nodes[i].Join(addr) {
			joinInfo.fail()
		} else {
			joinInfo.success()
		}
		nodesInNetwork = append(nodesInNetwork, i)

		time.Sleep(QASJoinSleepTime)
	}
	joinInfo.finish(&QASFailedCnt, &QASTotalCnt)

	time.Sleep(QASAfterJoinSleepTime)

	/* Put. */
	putInfo := testInfo{
		msg:       "Quit & Stabilize put",
		failedCnt: 0,
		totalCnt:  0,
	}
	_, _ = cyan.Printf("Start putting\n")
	for i := 0; i < QASPutSize; i++ {
		key := randString(lengthOfKeyValue)
		value := randString(lengthOfKeyValue)
		kvMap[key] = value

		if !nodes[rand.Intn(QASNodeSize+1)].Put(key, value) {
			putInfo.fail()
		} else {
			putInfo.success()
		}
	}
	putInfo.finish(&QASFailedCnt, &QASTotalCnt)

	/* All nodes quit. */
	getInfo := testInfo{
		msg:       "Quit & Stabilize Quit",
		failedCnt: 0,
		totalCnt:  0,
	}
	for t := 1; t <= QASNodeSize; t++ {
		/* Quit. */
		idxInArray := rand.Intn(len(nodesInNetwork))

		nodes[nodesInNetwork[idxInArray]].Quit()
		nodesInNetwork = removeFromArray(nodesInNetwork, idxInArray)

		time.Sleep(QASQuitSleepTime)

		/* Get some data. */
		getCnt := 0
		for key, value := range kvMap {
			ok, res := nodes[nodesInNetwork[rand.Intn(len(nodesInNetwork))]].Get(key)
			if !ok || res != value {
				getInfo.fail()
			} else {
				getInfo.success()
			}

			getCnt++
			if getCnt == QASGetSize {
				break
			}
		}
	}
	getInfo.finish(&QASFailedCnt, &QASTotalCnt)

	/* All nodes quit. */
	for i := 0; i <= QASNodeSize; i++ {
		nodes[i].Quit()
	}

	return panicked, QASFailedCnt, QASTotalCnt
}

/*
func testWhenStabAndQuit(rate time.Duration) {
	blue.Println("Start test StabAndQuit")
	info := make([]error, 4)
	defer func() {
		if r := recover(); r != nil {
			red.Println("Accidently end: ", r)
		}
		for _, inf := range info {
			totalCnt += inf.all
			totalFail += inf.cnt
		}
		if totalCnt == 0 {
			totalCnt++
			totalFail++
		}
	}()

	nodeGroup = new([maxNode]dhtNode)
	keyArray = new([maxData]string)
	datalocal = make(map[string]string)

	maxNodeSize = 100
	maxDataSize = 1200

	localIP = getIP()

	for i := 0; i < maxNodeSize; i++ {
		curport := config.Port + i
		nodeGroup[i] = NewNode(curport)

		go nodeGroup[i].Run()
	}
	time.Sleep(time.Millisecond * rate * 100)

	nodeGroup[0].Create()

	failcnt := 0
	cnt := 0
	for i := 1; i < maxNodeSize; i++ {
		curport := config.Port
		addr := portToAddr(localIP, curport)
		cnt++
		if !nodeGroup[i].Join(addr) {
			failcnt++
		}
		time.Sleep(time.Millisecond * 100 * rate)
	}
	info[0].initInfo("join", failcnt, cnt)
	info[0].finish()

	time.Sleep(time.Second * rate * 10)

	// fmt.Println("Force some node to quit")
	// for i := 150; i < maxNodeSize; i++ {
	// 	nodeGroup[i].ForceQuit()
	// 	time.Sleep(time.Millisecond * 200)
	// }
	// fmt.Println("Finish")

	failcnt = 0
	cnt = 0
	for i := 0; i < maxDataSize; i++ {
		k := randString(50)
		v := randString(50)
		keyArray[i] = k
		datalocal[k] = v

		cnt++
		if !nodeGroup[rand.Intn(maxNodeSize)].Put(k, v) {
			failcnt++
		}

		time.Sleep(time.Millisecond * rate)
	}
	info[1].initInfo("put", failcnt, cnt)
	info[1].finish()

	failcnt = 0
	cnt = 0
	for k, v := range datalocal {
		ok, ret := nodeGroup[rand.Intn(maxNodeSize)].Get(k)
		if !ok || ret != v {
			failcnt++
		}
		cnt++

		time.Sleep(time.Millisecond * rate)
	}
	info[2].initInfo("get", failcnt, cnt)
	info[2].finish()

	failcnt = 0
	cnt = 0
	for i := 0; i < maxNodeSize; i++ {
		for j := 1; j <= 10; j++ {
			rk := keyArray[rand.Intn(maxDataSize)]
			ok, ret := nodeGroup[i].Get(rk)

			cnt++
			if !ok || ret != datalocal[rk] {
				failcnt++
			}
			time.Sleep(time.Millisecond * rate * 10)
		}

		nodeGroup[i].Quit()
		time.Sleep(time.Millisecond * 100 * rate)
	}
	info[3].initInfo("get while quit", failcnt, cnt)
	info[3].finish()

	for i := 0; i < maxNodeSize; i++ {
		nodeGroup[i].Quit()
	}
}


func testRandom(rate time.Duration) {
	blue.Println("Start random test")
	info := make([]error, 4)
	defer func() {
		if r := recover(); r != nil {
			red.Println("Accidently end: ", r)
		}
		for _, inf := range info {
			totalCnt += inf.all
			totalFail += inf.cnt
		}
		if totalCnt == 0 {
			totalCnt++
			totalFail++
		}
	}()

	nodeGroup = new([maxNode]dhtNode)
	keyArray = new([maxData]string)
	datalocal = make(map[string]string)
	datamux := sync.Mutex{}
	maxNodeSize = 120

	localIP = getIP()

	for i := 0; i < maxNodeSize; i++ {
		// fmt.Println("run ", i)
		curport := config.Port + i
		nodeGroup[i] = NewNode(curport)

		go nodeGroup[i].Run()
	}
	time.Sleep(time.Millisecond * rate * 100)

	nodeGroup[0].Create()

	failcnt1 := 0
	cnt1 := 0
	running := true
	nodecnt := 1
	go func() {
		// fmt.Println("start join ")
		for running && nodecnt < maxNodeSize {
			curport := config.Port
			addr := portToAddr(localIP, curport)
			cnt1++
			if !nodeGroup[nodecnt].Join(addr) {
				failcnt1++
			}
			time.Sleep(time.Millisecond * 200 * rate)
			nodecnt++
		}
	}()

	//time.Sleep(time.Second * rate * 10)
	// fmt.Println("Force some node to quit")
	// for i := 150; i < maxNodeSize; i++ {
	// 	nodeGroup[i].ForceQuit()
	// 	time.Sleep(time.Millisecond * 200)
	// }
	// fmt.Println("Finish")

	failcnt2 := 0
	quitcnt := 1
	cnt2 := 0
	datacnt := 0
	time.Sleep(5 * time.Second)
	go func() {
		//fmt.Println("start put")
		for running && datacnt < maxDataSize {

			k := randString(50)
			v := randString(50)
			keyArray[datacnt] = k
			datamux.Lock()
			datalocal[k] = v
			datamux.Unlock()
			cnt2++
			if running && !nodeGroup[quitcnt+rand.Intn(nodecnt-quitcnt)].Put(k, v) {
				failcnt2++
			}
			datacnt++
			time.Sleep(time.Millisecond * 2 * rate)
		}
	}()

	failcnt3 := 0
	cnt3 := 0
	go func() {
		// fmt.Println("start get")
		for running {
			datamux.Lock()
			for k, v := range datalocal {
				if !running {
					break
				}
				tmp := quitcnt + rand.Intn(nodecnt-quitcnt)
				ok, ret := nodeGroup[tmp].Get(k)
				if !ok || ret != v {
					//fmt.Println("get fail:", k, " => ", v, " from ", tmp)
					failcnt3++
				}
				cnt3++
				time.Sleep(time.Millisecond * rate)
			}
			datamux.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	failcnt4 := 0
	cnt4 := 0
	time.Sleep(5 * time.Second)
	done := make(chan bool)
	go func() {
		// fmt.Println("start quit")
		for running {
			if quitcnt < nodecnt-1 {
				for j := 1; j <= 10; j++ {
					if datacnt <= 0 {
						time.Sleep(100 * time.Millisecond)
					}
					rk := keyArray[rand.Intn(datacnt)]
					tmp := quitcnt + rand.Intn(nodecnt-quitcnt)
					ok, ret := nodeGroup[tmp].Get(rk)
					cnt4++
					if !ok || ret != datalocal[rk] {
						//fmt.Println("get fail:", rk, " => ", datalocal[rk], " from ", tmp)
						failcnt4++
					}
					time.Sleep(time.Millisecond * rate * 100)
				}

				nodeGroup[quitcnt].Quit()
				quitcnt++
			} else if nodecnt == maxNodeSize {
				done <- true
			}
			time.Sleep(time.Second * rate)
		}
	}()*/
/*
	<-done
	running = false
	time.Sleep(5 * time.Second)
	nodeGroup[0].Quit()
	nodeGroup[maxNodeSize-1].Quit()

	info[0].initInfo("join", failcnt1, cnt1)
	info[0].finish()
	info[1].initInfo("put", failcnt2, cnt2)
	info[1].finish()
	info[2].initInfo("get", failcnt3, cnt3)
	info[2].finish()
	info[3].initInfo("get while quit", failcnt4, cnt4)
	info[3].finish()

}

*/
