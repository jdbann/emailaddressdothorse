package micropub_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"encore.app/micropub"
	"gotest.tools/v3/assert"
)

func TestHandle(t *testing.T) {
	type testCase struct {
		name       string
		baseURL    string
		method     string
		body       string
		header     http.Header
		wantCode   int
		wantHeader http.Header
		wantBody   string
	}

	run := func(t *testing.T, tc testCase) {
		baseURL, err := url.Parse(tc.baseURL)
		if err != nil {
			t.Fatal(err)
		}

		svc := micropub.Service{
			FrontendBaseURL: baseURL,
		}

		response := httptest.NewRecorder()
		request := httptest.NewRequest(tc.method, "/micropub", strings.NewReader(tc.body))
		for k, vs := range tc.header {
			for _, v := range vs {
				request.Header.Add(k, v)
			}
		}

		svc.Handle(response, request)

		assert.Equal(t, tc.wantCode, response.Code)
		assert.DeepEqual(t, tc.wantHeader, response.HeaderMap)
		assert.Equal(t, tc.wantBody, response.Body.String())
	}

	testCases := []testCase{
		{
			name:    "create an h-entry post (form-encoded)",
			baseURL: "https://blog.example.com",
			method:  http.MethodPost,
			body:    "h=entry&content=Micropub+test+of+creating+a+basic+h-entry",
			header: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			wantCode: http.StatusCreated,
			wantHeader: http.Header{
				"Location": []string{"https://blog.example.com/entry/123"},
			},
			wantBody: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			run(t, tc)
		})
	}
}
