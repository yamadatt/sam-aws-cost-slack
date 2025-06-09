package external

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/slack-go/slack"
)

func TestSlackClientPostMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chat.postMessage" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}))
	defer ts.Close()

	c := slack.New("test-token", slack.OptionAPIURL(ts.URL+"/"))
	s := SlackClient{client: c}

	err := s.PostMessage("test", slack.MsgOptionText("hello", false))
	if err != nil {
		t.Fatalf("post message failed: %v", err)
	}
}
