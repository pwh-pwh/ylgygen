package ylgygen

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func BrushScore(num int, token string, st int) {
	wg.Add(num)
	for i := 0; i < num; i++ {
		go GetS(token, st)
	}
	wg.Wait()
}
func GetS(token string, st int) (*http.Response, error) {
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
	return client.Do(req)
}
