package pipefilter

// Request is the input of the filter
type Request interface{}

// Response is the output of the filter
type Response interface{}

// Filter interface is the definition of the data processing components
// Pipe-Filter structure
type Filter interface {
	Process(data Request) (Response, error)
}

/*
Pipe-Filter 模式
非常适合与数据处理及数据分析系统
Filter封装数据处理的功能
松耦合：Filter只跟数据(格式耦合)
Pipe用于连接Filter传递数据或者在一部处理过程中缓冲数据流进程内
同步调用时，pipe演变为数据在方法调用间传递


"1,2,3"这个字符串通过SpiltFilter分割成字符串数组["1","2","3"]
通过ToIntFilter形成int型数组[1,2,3]通过SumFilter求和得6.
*/
