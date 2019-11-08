package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//访问间隔
var rateLimiter = time.Tick(900 * time.Millisecond)

func Fetcher(url string) ([]byte, error) {
	//限制执行速率
	<-rateLimiter
	//发起请求
	//resp, err := http.Get(url)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	for err!=nil {
		request, err = http.NewRequest(http.MethodGet, url, nil)
		if err==nil {
			request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0")
			break
		}
	}




	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	//打印一下失败时的http代码
	if response.StatusCode != http.StatusOK {
		log.Printf("http请求出错了")

		return nil, fmt.Errorf("http请求失败，code为：%d", response.StatusCode)
	}

	reader := bufio.NewReader(response.Body)
	//猜测文档编码
	e := determineEncoding(reader)

	utf8Reader := transform.NewReader(response.Body, e.NewDecoder())

	//读取解码内容
	return ioutil.ReadAll(utf8Reader)

}

//猜测文档编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {

	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher中determineEncoding出错：%v", err)
		return unicode.UTF8
	}

	//猜测文档的coding
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
