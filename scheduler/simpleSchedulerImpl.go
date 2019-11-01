package scheduler

import "shSpider_plus/engine"

type SimpleScheduler struct {
	//worker的输入管道
	WorkerIn chan engine.Request
}

//ConfigWorkerIn 设置worker的输入通道
func (scheduler *SimpleScheduler) ConfigWorkerIn(in chan engine.Request) {
	//为worker的输入管道赋值
	scheduler.WorkerIn = in
}

//Submit 将request输入worker
func (scheduler *SimpleScheduler) Submit(request engine.Request) {
	//把request输入进worker的输入管道 开协程来做这件事 避免处理过慢而造成循环等待 卡死
	go func() { scheduler.WorkerIn <- request }()
}
