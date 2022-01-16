package ch48

import "testing"

func TestCreateRequest(t *testing.T) {
	str := createRequest()
	t.Log(str)
}

func TestProcessRequest(t *testing.T) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	reps := processRequestNew(reqs)
	t.Log(reps[0])
}

func BenchmarkProcessRequestNew(b *testing.B) {

	reqs := []string{}
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequestNew(reqs)
	}
	b.StopTimer()

}

func BenchmarkProcessReQuestOld(b *testing.B) {

	reqs := []string{}
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequestOld(reqs)
	}
	b.StopTimer()

}

/*
go test -bench=. -benchmem
enchmarkProcessRequest-8  33043   36347 ns/op  20258 B/op  222 allocs/op
go test -bench=. -cpuprofile=cpu.prof(可生成cpu.prof文件)

go tool pprof cpu.prof
top -cum
flat  flat%   sum%        cum   cum%
0     0%     0%      2.57s 66.93%  WorkSpace/go_learning/src/ch48.BenchmarkProcessRequest
0.03s  0.78%  0.78%      2.57s 66.93%  WorkSpace/go_learning/src/ch48.processRequest
list processRequest

 .      1.83s     26:           json.Unmarshal([]byte(req), reqObj)
 .          .     27:           ret := ""
 .          .     28:           for _, e := range reqObj.PayLoad {
 .      640ms     29:                   ret += strconv.Itoa(e) + ","
 .          .     30:           }
20ms     40ms     31:           repObj := &Response{reqObj.TransactionID, ret}
 .       50ms     32:           repJson, err := json.Marshal(&repObj)

exit 退出
*/

/*
优化一：
使用easyjson优化自带的效率底下的通过反射实现的json
第一步：生成easyjson   easyjson -all structs.go
第二步：修改代码
第三步：测试修改后的性能

go test -bench=. -benchmem

enchmarkProcessRequestNew-8  53592 19420 ns/op  18632 B/op  113 allocs/op

go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
top -cum

flat  flat%   sum%    cum   cum%
0     0%     0%      1.44s 61.02%  WorkSpace/go_learning/src/ch48.BenchmarkProcessRequestNew
0.03s  1.27%  1.27%  1.44s 61.02%  WorkSpace/go_learning/src/ch48.processRequestNew

list processRequestNew

 .      520ms     47:           reqObj.UnmarshalJSON([]byte(req))
 .          .     48:           //json.Unmarshal([]byte(req), reqObj)
 .          .     49:           ret := ""
 .          .     50:           for _, e := range reqObj.PayLoad {
30ms      830ms     51:                   ret += strconv.Itoa(e) + ","
 .          .     52:           }
 .          .     53:           repObj := &Response{reqObj.TransactionID, ret}
 .          .     54:           //repJson, err := json.Marshal(&repObj)
 .       80ms     55:           repJson, err := repObj.MarshalJSON()

go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof
top -cum

flat  	 flat%   sum%     cum        cum%
0     	 0%      0%       998.25MB   100%  WorkSpace/go_learning/src/ch48.BenchmarkProcessRequestNew
854.17MB 85.57%  85.57%   998.25MB   100%  WorkSpace/go_learning/src/ch48.processRequestNew

list processRequestNew

19.51MB   117.57MB     47:           reqObj.UnmarshalJSON([]byte(req))
.          .     	   48:           //json.Unmarshal([]byte(req), reqObj)
.          .    	   49:           ret := ""
.          .    	   50:           for _, e := range reqObj.PayLoad {
815.16MB   815.16MB    51:                   ret += strconv.Itoa(e) + ","

*/

/*
优化二：
经过测试耗时最多变成字符串相加，因为string不可变，存在不断的内存复制，原有内存数据对GC也有影响
使用buffer优化

测试修改后性能

go test -bench=. -benchmem

enchmarkProcessRequestNew-8  125336  8908 ns/op  4664 B/op  20 allocs/op

go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof

flat  flat%  sum%    cum   cum%
0     0%     0%      0.86s 61.87%  WorkSpace/go_learning/src/ch48.BenchmarkProcessRequestNew
0     0%     0%      0.86s 61.87%  WorkSpace/go_learning/src/ch48.processRequestNew


list processRequestNew

.      580ms     49:           reqObj.UnmarshalJSON([]byte(req))
.          .     50:           // json.Unmarshal([]byte(req), reqObj)
.          .     51:           // ret := ""    // buf用不上了就
.          .     52:           var buf strings.Builder // 优化字符串拼接1
.          .     53:           for _, e := range reqObj.PayLoad {
.      140ms     54:                   buf.WriteString(strconv.Itoa(e)) // 优化字符串拼接2
.       70ms     55:                   buf.WriteString(",")             // 优化字符串拼接3

go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof

0     	 0%     0%  	1267.08MB   100%  WorkSpace/go_learning/src/ch48.BenchmarkProcessRequestNew
205.57MB 16.22% 16.22%  1267.08MB   100%  WorkSpace/go_learning/src/ch48.processRequestNew

list processRequestNew

101.53MB 617.38MB 49:           reqObj.UnmarshalJSON([]byte(req))
.          .      50:           // json.Unmarshal([]byte(req), reqObj)
.          .      51:           // ret := ""    // buf用不上了就
.          .      52:           var buf strings.Builder // 优化字符串拼接1
.          .      53:           for _, e := range reqObj.PayLoad {
.   102.02MB      54:                   buf.WriteString(strconv.Itoa(e)) // 优化字符串拼接2
.   195.58MB      55:                   buf.WriteString(",")             // 优化字符串拼接3

*/

/*
优化前后性能对比
go test -bench=. -benchmem

enchmarkProcessRequestNew-8  145118  13102 ns/op   4664 B/op  20 allocs/op
BenchmarkProcessReQuestOld-8 48616   41664 ns/op   19294 B/op 120 allocs/op

不错子！
*/
