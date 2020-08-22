package notification

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	var title string = "title1"
	var text string = "<!channel> from test"
	var domain string = "vaddy.net"
	var qType string = "NS"

	var expect []byte = []byte("{\"username\":\"\",\"icon_emoji\":\"\",\"icon_url\":\"\",\"channel\":\"\",\"text\":\"\",\"attachments\":[{\"color\":\"warning\",\"title\":\"title1\",\"text\":\"Type: NS, Domain: vaddy.net\\n\\n\\u003c!channel\\u003e from test\"}]}")

	result := createSlackMessage(title, text, domain, qType)

	if bytes.Compare(result, expect) != 0 {
		fmt.Printf("result: %s", result)
		t.Fail()
	}
}
