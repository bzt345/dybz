package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"novel/core"
	"strconv"
	"testing"
	"time"
)

func TestRabbitMq(t *testing.T) {
	go func() {
		//防止ip库还没运行
		time.Sleep(1)
		//爬虫获取小说列表章节列表进入rabbitmq
		re, _ := http.Get("http://localhost:8090/get")
		data, _ := ioutil.ReadAll(re.Body)
		if re != nil {
			var i IpInfo
			json.Unmarshal(data, &i)
			core.GetTitle("http://"+i.IP+":"+strconv.Itoa(i.Port), "http://www.diyibanzhu6.me/shuku/0-allvisit-0-1.html")
		}
		core.GetTitle("", "http://www.diyibanzhu6.me/shuku/0-allvisit-0-1.html")
		//章节详情进入队列
		//
	}()
	time.Sleep(2)
	defer func() {
		err := recover()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	for {
		go core.FixPage()
		go core.GetDetial()
		go core.GetContent()
		time.Sleep(time.Second * 100)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
