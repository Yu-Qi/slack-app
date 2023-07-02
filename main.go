package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/slack-go/slack"
)

// 你的應用程式資訊
const (
	RedirectURL = "https://b86c-114-32-26-14.ngrok-free.app/callback" // 根據你的需求設定
)

const (
	// MyExampleWorkflowStepCallbackID is configured in slack (api.slack.com/apps).
	// Select your app or create a new one. Then choose menu "Workflow Steps"...
	MyExampleWorkflowStepCallbackID = "msg123"
)

// 授權端點
const (
	AuthEndpoint     = "https://slack.com/oauth/v2/authorize"
	AccessTokenURL   = "https://slack.com/api/oauth.v2.access"
	TokenExchangeURL = "https://slack.com/api/oauth.v2.exchange"
)

var slackClient *slack.Client
var ClientID string
var ClientSecret string

func init() {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	ClientID = os.Getenv("SLACK_CLIENT_ID")
	ClientSecret = os.Getenv("SLACK_CLIENT_SECRET")

	slackClient = slack.New(botToken)
}

// 重新導向端點處理程序
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("callbackHandler, Request: %+v\n", r)

	// 解析查詢字串中的授權碼
	code := r.URL.Query().Get("code")

	// 交換授權碼以獲取訪問令牌
	data := url.Values{}
	data.Set("client_id", ClientID)
	data.Set("client_secret", ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", RedirectURL)

	resp, err := http.PostForm(AccessTokenURL, data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 讀取訪問令牌的回應
	// 在這裡你可以處理回應，例如儲存訪問令牌以供後續使用
	// 以下僅簡單印出回應的內容
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ok, Request: %+v\n", r)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ok":true}`))
	return
}

func main() {
	// 設定授權端點的參數
	authParams := url.Values{}
	authParams.Set("client_id", ClientID)
	authParams.Set("redirect_uri", RedirectURL)
	authParams.Set("scope", "YOUR_SCOPES") // 根據你的需求設定，例如：channels:history,channels:read

	// 註冊重新導向端點處理程序
	http.HandleFunc("/oauth/callback", callbackHandler)
	http.HandleFunc("/event", handleMyWorkflowStep)
	http.HandleFunc("/interactivity", handleInteraction)
	http.HandleFunc("/ok", ok)

	// 啟動伺服器
	log.Printf("Server started. Listening on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
