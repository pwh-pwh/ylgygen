package ylgygen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func BrushScore(num int, token string, st int) {
	var wg = new(sync.WaitGroup)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go GetS(token, st, wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func BrushScore2(uid string, num int, st int) {
	fmt.Println(num)
	var wg = sync.WaitGroup{}
	token := GetToken(uid)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go GetS(token, st, &wg)
	}
	wg.Wait()
	fmt.Println("aaa")
}

func GetToken(uid string) string {
	//https://cat-match.easygame2021.com/sheep/v1/game/user_info?uid={uid}
	client := &http.Client{}
	ltoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ1MzQxMzcsIm5iZiI6MTY2MzQzMTkzNywiaWF0IjoxNjYzNDMwMTM3LCJqdGkiOiJDTTpjYXRfbWF0Y2g6bHQxMjM0NTYiLCJvcGVuX2lkIjoiIiwidWlkIjoxMzU5Njk1MiwiZGVidWciOiIiLCJsYW5nIjoiIn0.rxNp69Cy_UmYZt1uzsGkIKFBOZehW3vXzo3kltJtybY"
	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://cat-match.easygame2021.com/sheep/v1/game/user_info?uid=%v&t=%v", uid, ltoken), nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Host", "cat-match.easygame2021.com")
	req.Header.Set("t", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ1MzQxMzcsIm5iZiI6MTY2MzQzMTkzNywiaWF0IjoxNjYzNDMwMTM3LCJqdGkiOiJDTTpjYXRfbWF0Y2g6bHQxMjM0NTYiLCJvcGVuX2lkIjoiIiwidWlkIjoxMzU5Njk1MiwiZGVidWciOiIiLCJsYW5nIjoiIn0.rxNp69Cy_UmYZt1uzsGkIKFBOZehW3vXzo3kltJtybY")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 MicroMessenger/6.5.2.501 NetType/WIFI MiniGame WindowsWechat")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-us,en")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://servicewechat.com/wx141bfb9b73c970a9/15/page-frame.html")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	tt := gjson.Get(string(data), "data.wx_open_id").String()
	type Ut struct {
		Uid      string
		NickName string `json:"nick_name"`
		Avatar   string
		Sex      int
	}
	ut := Ut{
		tt,
		"1",
		"http://aaa.png",
		1,
	}
	marshal, _ := json.Marshal(ut)
	fmt.Println(string(marshal))
	reader := bytes.NewReader(marshal)
	rq, _ := http.NewRequest("POST", "https://cat-match.easygame2021.com/sheep/v1/user/login_oppo", reader)
	rq.Header.Set("Host", "cat-match.easygame2021.com")
	rq.Header.Set("t", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ1MzQxMzcsIm5iZiI6MTY2MzQzMTkzNywiaWF0IjoxNjYzNDMwMTM3LCJqdGkiOiJDTTpjYXRfbWF0Y2g6bHQxMjM0NTYiLCJvcGVuX2lkIjoiIiwidWlkIjoxMzU5Njk1MiwiZGVidWciOiIiLCJsYW5nIjoiIn0.rxNp69Cy_UmYZt1uzsGkIKFBOZehW3vXzo3kltJtybY")
	rq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 MicroMessenger/6.5.2.501 NetType/WIFI MiniGame WindowsWechat")
	rq.Header.Set("Accept", "*/*")
	rq.Header.Set("Sec-Fetch-Site", "cross-site")
	rq.Header.Set("Sec-Fetch-Mode", "cors")
	rq.Header.Set("Sec-Fetch-Dest", "empty")
	rq.Header.Set("Accept-Encoding", "gzip, deflate, br")
	rq.Header.Set("Accept-Language", "en-us,en")
	rq.Header.Set("Connection", "keep-alive")
	rq.Header.Set("Content-Type", "application/json")
	rq.Header.Set("Referer", "https://servicewechat.com/wx141bfb9b73c970a9/15/page-frame.html")
	resp1, _ := client.Do(rq)
	all, _ := ioutil.ReadAll(resp1.Body)
	fmt.Println(string(all))
	return gjson.Get(string(all), "data.token").String()
}

func GetS(token string, st int, wg *sync.WaitGroup) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://cat-match.easygame2021.com/sheep/v1/game/game_over?rank_score=1&rank_state=1&rank_time=%v&rank_role=1&skin=1", st), nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Host", "cat-match.easygame2021.com")
	req.Header.Set("t", token)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 MicroMessenger/6.5.2.501 NetType/WIFI MiniGame WindowsWechat")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-us,en")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://servicewechat.com/wx141bfb9b73c970a9/15/page-frame.html")
	defer wg.Done()
	//fmt.Println("get S")
	return client.Do(req)
}
