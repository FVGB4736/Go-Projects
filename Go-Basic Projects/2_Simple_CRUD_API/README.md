# 2_Simple_CRUD_API

簡介
---
這個專案是一個用 Go (Golang) 實作的簡單 CRUD（Create / Read / Update / Delete）範例 API，示範如何使用 `net/http` 與第三方路由器 `github.com/gorilla/mux` 建立 RESTful 介面。

資料模型
---
- Movie
  - `ID` (string)
  - `Isbn` (string)
  - `Title` (string)
  - `Director` (*Director)
- Director
  - `FirstName` (string)
  - `LastName` (string)

說明重點
---
- struct tag（例如 `json:"title"`）用於指定 struct 欄位在 JSON 的鍵名稱，以及額外選項（像 `omitempty` 或 `-`）。
- 欄位使用指標（例如 `Director *Director`）可以表示該欄位為可選（nil）、避免複製大型結構，並且在序列化時能更靈活地省略空值。

依賴
---
- Go (建議 >= 1.18)
- 第三方套件：`github.com/gorilla/mux`

快速啟動（Windows PowerShell）
---
1. 開啟 PowerShell，切換到專案目錄：

```powershell
cd "f:\Practice Project\Go Projects\Go-Basic Projects\2_Simple_CRUD_API"
```

2. 建立或啟用 go module（若尚未建立）：

```powershell
# 範例 module 名稱，可自行替換
go mod init example.com/2_simple_crud_api
# 取得依賴
go get github.com/gorilla/mux
# 或：
# go mod tidy
```

3. 執行程式：

```powershell
# 如果 main.go 已經啟動 HTTP server（通常使用 http.ListenAndServe）
go run main.go
```

4. 測試 API（範例 curl）：

```powershell
# 列出所有電影
curl http://localhost:8000/movies

# 取得單一電影
curl http://localhost:8000/movies/1

# 建立電影（JSON body）
curl -X POST -H "Content-Type: application/json" -d '{"isbn":"4444","title":"My Movie","director":{"firstname":"John","lastName":"Doe"}}' http://localhost:8000/movies
```

註：Windows 下的 curl 可能是系統內建的，或可改用 Postman / HTTPie / Invoke-RestMethod。

常見問題
---
- 如果 `json:"field"` 沒有用雙引號包住（例如 `json:field`），Go 的反射在解析 tag 時會無法正確讀取，會導致編碼/解碼行為不符合預期。
- 如果不想把某個欄位輸出到 JSON，可以使用 `json:"-"`。
- 若要在序列化時省略零值，可以使用 `json:"name,omitempty"`。

進階建議
---
- 建議在專案中加入簡單的路由定義與啟動伺服器（例如在 `main.go` 中呼叫 `http.ListenAndServe(":8000", r)`），並在 README 裡說明實際使用的 port。
- 可以加入範例測試 (Go unit tests) 以及 Postman collection 以方便測試 API。

檔案
---
- `main.go`：主要伺服器程式與路由設定

FVGB4736