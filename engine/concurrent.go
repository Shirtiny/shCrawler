package engine

import (
	"fmt"
	"log"
	"shSpider_plus/fetcher"
)

//并发引擎 不再用
type Concurrent struct {
	//调度器
	Scheduler Scheduler
	//worker协程数
	WorkerCount int
}

//接口 调度器 接口方法的参数不需要名字
type Scheduler interface {
	//设置worker的输入通道
	ConfigWorkerIn(chan Request)
	//将request输入worker
	Submit(Request)
}

//将所有request作为参数依次传入seeds数组中
func (engine *Concurrent) Run(seeds ...Request) {

	//制作两个channel worker协程的通道
	//in负责对worker输入 传输Request
	in := make(chan Request)
	//设置调度器中的worker输入管道
	engine.Scheduler.ConfigWorkerIn(in)

	//out负责worker输出 传输ParseResult
	out := make(chan ParseResult)

	//创建workerCount个worker协程
	for i := 0; i < engine.WorkerCount; i++ {
		createWorker(in, out)
	}

	//将所有request交给调度器 调度器对worker输入
	for _, request := range seeds {
		engine.Scheduler.Submit(request)
	}

	//处理worker的输出结果
	for {
		//接收输出结果
		result := <-out
		//打印解析结果中的objects result.Objects
		for _, object := range result.Objects {
			fmt.Printf("worker输出的对象为：%v\n", object)
		}
		//处理解析出的新请求 result.Requests
		for _, request := range result.Requests {
			//将新请求提交给调度器处理 => worker输出中取request输入到scheduler
			engine.Scheduler.Submit(request)
		}
	}

}

//整合了fetcher和Parser
func worker(request Request) (ParseResult, error) {
	//使用fetcher发起请求，得到请求url后的数据
	bytes, err := fetcher.Fetcher(request.Url)
	fmt.Printf("请求的URL：%s\n", request.Url)
	if err != nil {
		//处理错误err
		log.Printf("fetcher请求失败,url：%s ；error：%v", request.Url, err)
		//停止后续操作，进入下一次循环
		//continue
		return ParseResult{}, err
	}
	//解析fetcher请求url返回的数据
	return request.ParserFunc(bytes), nil
}

//创建worker协程
func createWorker(in chan Request, out chan ParseResult) {
	//协程
	go func() {
		for {
			//从request管道中取出request请求
			request := <-in
			//调用worker处理请求
			result, e := worker(request)
			//出错 结束本轮循环，进入下一轮 不对result处理
			if e != nil {
				continue
			}
			//没出错则将result放入ParseResult管道
			out <- result
		}
	}()
}
