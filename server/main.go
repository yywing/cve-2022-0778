package main

import (
	"cve-2022-0778/certfile"
	"cve-2022-0778/tls"
	_ "embed"
	"flag"
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	// 设置一下超时
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	//  触发一下握手
	conn.Write([]byte{1})
	fmt.Println("connection close")
}

var addr string

func Init() {
	flag.StringVar(&addr, "addr", "127.0.0.1:12345", "addr")
}

func main() {
	cert, err := tls.X509KeyPair(certfile.Cert, certfile.Key)
	if err != nil {
		fmt.Println(err)
		return
	}
	cert.Certificate = [][]byte{certfile.BadCert}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MaxVersion:   tls.VersionTLS12,
	}

	// 初始化变量 cliFlag
	Init()
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	ln, err := tls.Listen("tcp", addr, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("start server")

	for {
		// 等待客户端建立连接
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		fmt.Printf("accept %s\n", conn.RemoteAddr().String())
		// 启动一个单独的 goroutine 去处理连接
		go process(conn)
	}
}
