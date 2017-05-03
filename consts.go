package main

import (
	"net/http"
	"os"
)

// ResultQQwry 归属地信息
type ResultQQwry struct {
	IP      	string `json:"ip"`
	IPSegment	string `json:"ip_segment"`
	Address    	string `json:"address"`
}

// Response 向客户端返回数据的
type Response struct {
	r *http.Request
	w http.ResponseWriter
}
