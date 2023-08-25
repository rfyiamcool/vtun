package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/net-byte/vtun/app"
	"github.com/net-byte/vtun/common/config"
)

var (
	_version   = "v1.7.1"
	_gitHash   = "nil"
	_buildTime = "nil"
	_goVersion = "nil"
)

func displayVersionInfo() {
	log.Printf("vtun version -> %s", _version)
	log.Printf("git hash -> %s", _gitHash)
	log.Printf("build time -> %s", _buildTime)
	log.Printf("go version -> %s", _goVersion)
}

func main() {
	var config = config.Config{}
	var configFile string

	flag.StringVar(&configFile, "f", "", "config file (support json/yaml)")
	flag.StringVar(&config.DeviceName, "dn", "", "device name")
	flag.StringVar(&config.CIDR, "c", "172.16.0.10/24", "tun interface cidr")
	flag.StringVar(&config.CIDRv6, "c6", "fced:9999::9999/64", "tun interface ipv6 cidr")
	flag.IntVar(&config.MTU, "mtu", 1500, "tun mtu")
	flag.StringVar(&config.LocalAddr, "l", ":3000", "local address")
	flag.StringVar(&config.ServerAddr, "s", ":3001", "server address")
	flag.StringVar(&config.ServerIP, "sip", "172.16.0.1", "server ip")
	flag.StringVar(&config.ServerIPv6, "sip6", "fced:9999::1", "server ipv6")
	flag.StringVar(&config.Key, "k", "freedom@2023", "key")
	flag.StringVar(&config.Protocol, "p", "udp", "protocol udp/tls/grpc/quic/utls/dtls/h2/http/tcp/https/ws/wss")
	flag.StringVar(&config.Path, "path", "/freedom", "path")
	flag.BoolVar(&config.ServerMode, "S", false, "server mode")
	flag.BoolVar(&config.GlobalMode, "g", false, "client global mode")
	flag.BoolVar(&config.Obfs, "obfs", false, "enable data obfuscation")
	flag.BoolVar(&config.Compress, "compress", false, "enable data compression")
	flag.IntVar(&config.Timeout, "t", 30, "dial timeout in seconds")
	flag.StringVar(&config.TLSCertificateFilePath, "certificate", "./certs/server.pem", "tls certificate file path")
	flag.StringVar(&config.TLSCertificateKeyFilePath, "privatekey", "./certs/server.key", "tls certificate key file path")
	flag.StringVar(&config.TLSSni, "sni", "", "tls handshake sni")
	flag.BoolVar(&config.TLSInsecureSkipVerify, "isv", false, "tls insecure skip verify")
	flag.BoolVar(&config.Verbose, "v", false, "enable verbose output")
	flag.BoolVar(&config.PSKMode, "psk", false, "enable psk mode (dtls only)")
	flag.StringVar(&config.Host, "host", "", "http host")
	flag.Parse()
	displayVersionInfo()
	if configFile != "" {
		err := config.LoadConfig(configFile)
		if err != nil {
			log.Fatalf("Failed to load config from file: %s", err)
		}
	}
	app := app.NewApp(&config, _version)
	app.InitConfig()
	go app.StartApp()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.StopApp()
}
