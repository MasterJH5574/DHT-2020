# Distributed Hash Table - PPCA 2020

## Overview

 - A DHT can be viewed as a dictionary service distributed over a network: it provides access to a common shared key-&gt;value data-store, distributed over participating nodes with great performance and scalability.
 - From a user perspective, a DHT essentially provides a map interface, with two main operations: <code>put(key, value)</code> and <code>get(key)</code>. Get will retrieve values stored at a certain key while put (often called announce) will store a value on the network. Note that many values can be stored under the same key.
 - There are many algorithms to implement DHT. For this project, you are required to <b>implement at least Chord protocol</b>. You are also required to <b>implement an application of DHT or implement another protocol</b>. Finally, you should <b>write a report for about one page</b>.

More info in [Wiki: DHT](https://en.wikipedia.org/wiki/Distributed_hash_table).

## Assignment

* Use Go to implement a Chord DHT with basic functions.
* Use this DHT to implement an easy application, or implement another protocol.

## Syllabus

* Learn GoLang.
* Implement Chord protocol.
* Implement an application of DHT or implement another protocol.

## Score

GitHub repository for test: [DHT-2020](https://github.com/MasterJH5574/DHT-2020)

- 70% for the Chord Test (60 + 10).
  - 60% + 10%: 60% for basic test and 10% for advance test.
  - Basic test: naive test without "force quit".
  - Advance test: "Force quit" will be tested. There will be some more complex tests.
- 20% for the application/second protocol.
- 10% for a short report and code review.

## Tests

Unluckily, DHT tests cannot run successfully under Windows. So if you want to run tests by yourself, it is recommended to run tests under a Linux virtual machine, or employ a remote server.

If you want to debug tests by yourself, you can still use GoLand under virtual machine, or use Delve (a GoLang debugger) + GoLand if you employed a remote server.

Contact TA if you find any bug in the test program, or if you have some test ideas, or if you think the tests are too hard and you want TA to make it easier.

### Basic Test

The current test procedure is:

* There are **5 rounds** of test in total.
* In each round,
  1. **20 nodes** join the network. Then **sleep for 10 seconds.**
  2. **Put 200 key-value pairs**, **query for 160 pairs**, and then **delete 100 pairs**. There is **no sleep time between contiguous two operations**.
  3. **10 nodes** quit from the network. Then **sleep for 10 seconds**.
  4. (The same as 2.) **Put 200 key-value pairs**, **query for 160 pairs**, and then **delete 100 pairs**. There is **no sleep time between contiguous two operations**.

### Advance Test

Not finished yet.

### How to Test Your Chord?

1. Clone this repository using `git clone https://github.com/MasterJH5574/DHT-2020.git`.
2. Set the environment variables `GOROOT` and `GOPATH` correctly.
3. Under `GOPATH`, run `go get -u -v github.com/fatih/color` in shell to install the color package.
4. Replace `$GOPATH/src/main/userdef.go` with your own `userdef.go`. Do not modify `$GOPATH/src/main/interface.go`.
5. Copy your Chord package(s) into `$GOPATH/src` directory.
6. Under `GOPATH`, run `go build main` to generate the executable `main`. Then use `./main -test basic` or `./main -test advance` or `./main -test all` to run the corresponding test. (Or you can use `go run main -test [testName]` to run tests directly without generating the executable `main`.) (Or you can use GoLand to run the tests.)

### About Go Remote Debug

Please reference this [guide](https://github.com/MasterJH5574/DHT-2020/blob/master/guide/Go-Remote-Debug.md).

## Reference

- Learn Go
[A tour of Go](https://tour.golang.org/)
[Go package docs](http://golang.org/pkg/)
[Books about Go](https://github.com/golang/go/wiki/Books)
- DHT models
[Chord](https://en.wikipedia.org/wiki/Chord_(peer-to-peer))
[Pastry](https://en.wikipedia.org/wiki/Pastry_(DHT))
[Kademlia](https://en.wikipedia.org/wiki/Kademlia)
- Related project framework
[Dixie](https://cit.dixie.edu/cs/3410/asst_chord.php)
[CMU](https://www.cs.cmu.edu/~dga/15-744/S07/lectures/16-dht.pdf)
[MIT](https://pdos.csail.mit.edu/papers/sit-phd-thesis.pdf)

