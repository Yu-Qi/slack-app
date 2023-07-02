# 目標

# 各 endpoint 使用場景

1. `/oauth/callback`

   - 用戶安裝 app 時，slack server 會跟用戶確認，其 app 所需要的授權，接著 slack server 會與 app(oauth client)，透過 oauth 完成用戶資料驗證並且取得

2. `/event`

   - `workflow_step_execute`: ，我們會收到用戶根據我們自訂的這個 step 需要填寫的資料

3. `/interactivity`
   - `workflow_step_edit: 當用戶建立 workflow 時加入 step 時，由於我們有訂閱 `workflow_step_edit` 這個事件，我們要回傳需要用戶的哪些資料，他會有一個格式，可以設計像是填空、下拉式選單等

# 使用方式

1. 複製 config/env.sh.example，並且填入相關資訊
2. go run main.go
3. ngrok http 8888
