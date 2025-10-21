package main

import (
    "fmt"       // 導入 fmt 包，用於格式化輸出，類似 Java 的 System.out.println
    "log"       // 導入 log 包，用於記錄致命錯誤，類似 Java 的日誌框架
    "net/http"  // 導入 net/http 包，提供 HTTP 伺服器功能，類似 Java Servlet
)

// formHandler 處理表單提交的請求，類似 Java Servlet 的 doPost 方法
func formHandler(w http.ResponseWriter, r *http.Request) {
    // 解析表單數據，如果失敗則回應錯誤訊息
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err) // 寫入錯誤訊息到回應，類似 response.getWriter().println
        return
    }
    fmt.Fprintf(w, "Post request successful") // 回應表單提交成功

    // 從表單中提取 name 和 address 欄位值，類似 request.getParameter
    name := r.FormValue("name")
    address := r.FormValue("address")

    // 將提取的表單數據寫入回應
    fmt.Fprintf(w, "name = %s\n", name)
    fmt.Fprintf(w, "address = %s\n", address)
}

// helloHandler 處理 /hello 路徑的 GET 請求，類似 Java Servlet 的 doGet 方法
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Received method:", r.Method) // 記錄請求方法到終端，方便除錯，類似 System.out.println

    // 檢查路徑是否為 /hello，若不是則返回 404 錯誤
    if r.URL.Path != "/hello" {
        http.Error(w, "404 Not Found", http.StatusNotFound) // 回應 404 錯誤，類似 response.sendError
        return
    }

    // 檢查請求方法是否為 GET，若不是則返回 405 錯誤
    if r.Method != "GET" {
        http.Error(w, "Method is not supported ", http.StatusMethodNotAllowed) // 回應 405 錯誤
        return
    }

    // 回應成功訊息
    fmt.Fprintf(w, "Hello ! Welcome to my website")
}

func main() {
    // 創建靜態檔案伺服器，服務 static 資料夾，類似 Java Web 應用中的靜態資源映射
    fileserver := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileserver) // 將根路徑 / 及以下請求映射到靜態檔案伺服

    // 註冊 /form 路徑的處理函數，類似 Java 的 @WebServlet("/form")
    http.HandleFunc("/form", formHandler)
    // 註冊 /hello 路徑的處理函數，類似 Java 的 @WebServlet("/hello")
    http.HandleFunc("/hello", helloHandler)

    fmt.Println("Starting server at port 8080\n") // 輸出伺服器啟動訊息

    // 啟動 HTTP 伺服器，監聽 8080 端口，nil 表示使用預設路由器
    // 類似 Java 中 Tomcat 的啟動，但 Go 無需容器
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err) // 如果啟動失敗，記錄致命錯誤並退出
    }
}