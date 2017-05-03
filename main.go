package main

import (
	"github.com/kayon/qqwry"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

var QQWry qw

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	datFile := flag.String("qqwry", "./qqwry.dat", "纯真 IP 库的地址")
	port	:= flag.String("port", "2060", "HTTP 请求监听端口号")

	flag.Parse()

	qw = qqwry.New(*datFile)

    var res qqwry.Result
    res = qw.Search("114.114.114.114")
    fmt.Printf("IP: %s\nBegin: %s\nEnd: %s\nCountry: %s\nArea: %s\n", res.IP, res.Begin, res.End, res.Country, res.Area)

	

	// 下面开始加载 http 相关的服务
	http.HandleFunc("/", findIP)

	log.Printf("开始监听网络端口:%s", *port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
		log.Println(err)
	}
}

// findIP 查找 IP 地址的接口
func findIP(w http.ResponseWriter, r *http.Request) {
	res := NewResponse(w, r)

	ip := r.Form.Get("ip")

	if ip == "" {
		res.ReturnError(http.StatusBadRequest, 200001, "请填写 IP 地址")
		return
	}

	ips := strings.Split(ip, ",")

	qqWry := NewQQwry()

	rs := map[string]ResultQQwry{}
	if len(ips) > 0 {
		for _, v := range ips {
			resWry := &ResultQQwry{}

			qqwryRes = qw.Search(v)

			resWry.IP = v
			resWry.IPSegment = qqwryRes.Begin + "-" + qqwryRes.End
			resWry.Address = qqwryRes.Country
			
			rs[v] = resWry
		}
	}

	res.ReturnSuccess(rs)
}
