package larkvc

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

type botRequestClient struct {
	request *http.Request
}

func (client *botRequestClient) Do(req *http.Request) (*http.Response, error) {
	client.request = req
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{"code":0,"data":{"has_more":false,"events":[]}}`)),
	}, nil
}

func TestBotEventsWithSupportedTokens(t *testing.T) {
	cases := []struct {
		name   string
		option larkcore.RequestOptionFunc
	}{
		{name: "user", option: larkcore.WithUserAccessToken("test-only-token")},
		{name: "tenant", option: larkcore.WithTenantAccessToken("test-only-token")},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			httpClient := &botRequestClient{}
			config := &larkcore.Config{
				AppId:        "test-app",
				BaseUrl:      "https://sdk-test.invalid",
				HttpClient:   httpClient,
				Serializable: &larkcore.DefaultSerialization{},
			}
			larkcore.NewLogger(config)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			resp, err := New(config).Bot.Events(ctx,
				NewEventsBotReqBuilder().MeetingId("meeting-test").PageSize(20).PageToken("next-page").Build(),
				testCase.option)
			if err != nil {
				t.Fatalf("request bot events: %v", err)
			}
			if !resp.Success() || resp.Data == nil || resp.Data.HasMore == nil || *resp.Data.HasMore {
				t.Fatal("bot events response was not decoded correctly")
			}
			req := httpClient.request
			if req == nil || req.Method != http.MethodGet || req.URL.Path != "/open-apis/vc/v1/bots/events" {
				t.Fatal("unexpected bot events method or route")
			}
			if req.URL.Query().Get("meeting_id") != "meeting-test" || req.URL.Query().Get("page_size") != "20" || req.URL.Query().Get("page_token") != "next-page" {
				t.Fatal("bot events query parameters were not encoded correctly")
			}
			if req.Header.Get("Authorization") != "Bearer test-only-token" {
				t.Fatal("explicit token was not applied")
			}
		})
	}
}
