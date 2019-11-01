package engine

//Request，存储url、解析器方法
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//解析器方法的返回值结构体，存储request数组、object数组
type ParseResult struct {
	Requests []Request
	//表示任意类型的数组
	Objects []interface{}
}

//空的解析器函数，临时给没有解析器的url使用
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
