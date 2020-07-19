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

