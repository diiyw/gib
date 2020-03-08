package dom

import "testing"

func TestQuery(t *testing.T) {
	result, err := Query(
		Link("https://echo.labstack.com/"),
		WithDOM(DOM{
			Name:     "",
			Text:     false,
			Html:     false,
			Attr:     "",
			Selector: "body > div.w3-content > div > div.w3-row-padding > div > div.hero > h1",
			Markdown: false,
		}),
	)
	if err != nil {
		t.Error(err)
		return
	}
	if result["0"]["text"] != "Echo" {
		t.Error("unexpect result: ", result)
	}
}
