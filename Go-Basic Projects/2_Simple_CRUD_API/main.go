package main

import (
	"encoding/json" // Go 標準庫中用於處理 JSON 數據的套件
	"fmt"
	"log"
	"math/rand" // 生成隨機數的標準庫套件
	"net/http"  // 提供 HTTP 客戶端和服務器的實現, 可以用來創建 Web 服務器、處理 HTTP 請求和響應, 提供基本的路由功能
	"strconv"   // 用於字符串和基本數據類型之間轉換的標準庫套件

	"github.com/gorilla/mux" //第三方路由器套件, 提供更好的 URL 模式匹配和路由處理
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			//1.取 movies 中從開頭到 index-1 的元素（movies[:index]）。
			//2.取 movies 中從 index+1 到結尾的元素（movies[index+1:]）。
			//3.使用 append 將這兩部分合併，跳過索引為 index 的元素，實現刪除效果。
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	//Decode 方法返回一個 error 值，表示解碼是否成功（例如，如果 JSON 格式錯誤，會返回非 nil 的錯誤）。此處表示忽略ERROR
	//將 JSON 數據解碼並填充到指定的變數中。這裡的參數是 &movie，表示 movie 變數的記憶體地址（指標）。
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	//來自 strconv 套件，Itoa 是「Integer to ASCII」的縮寫，將整數轉換為字符串。
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "movie 1", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "456", Title: "movie 2", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
