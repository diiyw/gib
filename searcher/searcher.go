package searcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type IndexResult struct {
	Id   string
	Text string
}

type Searcher struct {
	Url  string
	Name string
}

func NewSearcher(name, url string) *Searcher {
	return &Searcher{url, name}
}

// 创建索引
func (s *Searcher) Create(id, text string) bool {
	resp, err := http.Get(s.Url + "/create/" + s.Name + "/" + id + "?text=" + text)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return false
	}
	v, ok := result["Message"].(string)
	return ok && v == "ok"
}

// 搜索
func (s *Searcher) Search(text string, offset int) (r []IndexResult) {
	resp, err := http.Get(s.Url + "/search/" + s.Name + "?text=" + text + "&p=" + strconv.Itoa(offset))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return
	}
	// 结果
	if result["Data"] == nil {
		return
	}

	index, ok := result["Data"].([]interface{})
	if !ok {
		return
	}

	for _, idx := range index {
		i, ok := idx.(map[string]interface{})
		if ok {
			r = append(r, IndexResult{
				Id:   i["DocId"].(string),
				Text: i["Content"].(string),
			})
		}

	}
	return
}
