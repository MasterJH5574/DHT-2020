# Go Remote Debug

1. 在服务器上安装 Go：
   * 下载 [GoLang 1.14.6](https://golang.org/dl/go1.14.6.linux-amd64.tar.gz)；
   * `sudo tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz` 将其解压至 `/usr/local`；
   * 将 `/usr/local/go/bin` 加到 `PATH` 环境变量中。
   
2. 服务器需要允许对于端口的监听（以阿里云为例，需要修改安全组以支持对端口的监听）；

3. 将你的代码从本地传至服务器（通过 GitHub、scp 等方式）；

4. 在服务器上设置好 `GOROOT` 与 `GOPATH`；

5. 在服务器上[安装 Delve](https://github.com/go-delve/delve/blob/master/Documentation/installation/linux/install.md)：
   * ```shell
     $ git clone https://github.com/go-delve/delve.git $GOPATH/src/github.com/go-delve/delve
     $ cd $GOPATH/src/github.com/go-delve/delve
     $ make install
     ```
     
   * 如果在 `make install` 出现如下超时报错：
     ```
     go: github.com/cosiner/argv@v0.1.0: Get "https://proxy.golang.org/github.com/cosiner/argv/@v/v0.1.0.mod": dial tcp 34.64.4.113:443: i/o timeout
     Makefile:10: recipe for target 'install' failed
     make: *** [install] Error 1
     ```
     则需要先通过 `go env -w GOPROXY=https://goproxy.cn` 修改 Go 模块代理，再次尝试 `make install`，应能顺利完成；
     
   * 此时 `$GOPATH/bin` 目录下应该存在名为 `dlv` 的可执行文件，需要将 `$GOPATH/bin` 添加至 `PATH` 环境变量中。
   
6. 在 GoLand 中添加 "Go Remote" 配置：点 "Edit Configuration"，添加一个 "Go Remote" 的 configuration，设置 "Host" 为服务器的公网 IP 即可；

7. 在服务器上 build 你的代码：假设需要 build 的是 package main，那么需要在 `$GOPATH` 下用 `go build -gcflags 'all=-N -l' main` 来 build 代码；

8. 在服务器上运行可执行文件：假设编译得到的可执行文件为 `main`，则通过 `dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./main -- -test=basic` 运行可执行文件（可将 `basic` 替换为 `advance` 或 `all`）；

9. 在本地 GoLand 中设置断点，点击 "Debug" 即可开始调试。

10. 注：似乎 dlv 并不能通过 Ctrl+C 终止运行，如果想要终止 dlv 的运行，可以使用 `killall dlv` 杀死这个进程。