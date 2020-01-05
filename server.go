// +build !wasm

package main

//go:generate vugugen .

import (
	"flag"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/vugu/vugu/simplehttp"
)

func main() {
	dev := flag.Bool("dev", false, "Enable development features")
	dir := flag.String("dir", ".", "Project directory")
	host := flag.String("host", "0.0.0.0", "Listen for HTTP on this host")
	port := flag.String("port", "8877", "Listen for HTTP on this port")
	flag.Parse()
	wd, err := filepath.Abs(*dir)
	if err != nil {
		logrus.WithError(err).Errorln("join dir error")
		return
	}
	err = os.Chdir(wd)
	if err != nil {
		logrus.WithError(err).Errorln("change work dir error")
		return
	}
	address := net.JoinHostPort(*host, *port)
	_, err = net.ResolveTCPAddr("tcp", address)
	if err != nil {
		logrus.WithError(err).Errorln("resolve tcp addr error")
		return
	}
	logrus.Infoln("Starting HTTP Server at", address)

	simplehttp.DefaultStaticData["CSSFiles"] = []string{
		"/static/css/bootstrap.min.css",
	}
	simplehttp.DefaultStaticData["Title"] = "协议解析"
	h := simplehttp.New(wd, *dev)
	logrus.Fatal(http.ListenAndServe(address, h))
}
