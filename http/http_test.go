package http

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {

	var params = url.Values{}
	params.Add("first", "1")
	b, err := Get("http://httpbin.org/get", params)
	if err != nil {
		t.Error(err)
	}
	var m = make(map[string]interface{})
	_ = json.Unmarshal(b, &m)
	if m["headers"].(map[string]interface{})["User-Agent"].(string) != DefaultAgent {
		t.Error("not match user-agent")
	}
	if m["args"].(map[string]interface{})["first"].(string) != "1" {
		t.Error("param error")
	}
}

func TestPost(t *testing.T) {
	var params = url.Values{}
	params.Add("first", "1")
	b, err := Post("http://httpbin.org/post", params.Encode(), nil)
	if err != nil {
		t.Error(err)
	}
	var m = make(map[string]interface{})
	_ = json.Unmarshal(b, &m)
	if m["headers"].(map[string]interface{})["User-Agent"].(string) != DefaultAgent {
		t.Error("not match user-agent")
	}
	if m["form"].(map[string]interface{})["first"].(string) != "1" {
		t.Error("param error")
	}

	fmt.Println(string(b))
}
