package ch48

import (
	"encoding/json"
	"strconv"
	"strings"
)

// 客户端生成了一个序列化好的Json请求
func createRequest() string {
	payload := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		payload[i] = i
	}
	req := Request{"demo_transaction", payload}
	v, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}
	return string(v)
}

//服务端处理请求,主要调优的是process的过程
func processRequestOld(reqs []string) []string {
	reps := []string{}
	for _, req := range reqs {
		reqObj := &Request{}
		json.Unmarshal([]byte(req), reqObj)
		ret := ""
		for _, e := range reqObj.PayLoad {
			ret += strconv.Itoa(e) + ","
		}
		repObj := &Response{reqObj.TransactionID, ret}
		repJson, err := json.Marshal(&repObj)
		if err != nil {
			panic(err)
		}
		reps = append(reps, string(repJson))
	}
	return reps
}

// 优化使用easyjson
// 优化用string.Builder
func processRequestNew(reqs []string) []string {
	reps := []string{}
	for _, req := range reqs {
		reqObj := &Request{}
		reqObj.UnmarshalJSON([]byte(req))
		// json.Unmarshal([]byte(req), reqObj)
		// ret := ""	// buf用不上了就
		var buf strings.Builder // 优化字符串拼接1
		for _, e := range reqObj.PayLoad {
			buf.WriteString(strconv.Itoa(e)) // 优化字符串拼接2
			buf.WriteString(",")             // 优化字符串拼接3
			// ret += strconv.Itoa(e) + ","
		}
		//repObj := &Response{reqObj.TransactionID, ret}
		repObj := &Response{reqObj.TransactionID, buf.String()} // 优化字符串拼接4
		//repJson, err := json.Marshal(&repObj)
		repJson, err := repObj.MarshalJSON()
		if err != nil {
			panic(err)
		}
		reps = append(reps, string(repJson))
	}
	return reps
}
