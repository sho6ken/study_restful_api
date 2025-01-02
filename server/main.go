package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 地址
type Address struct {
	City string `json:"city,omitempty"` // 居住縣市
}

// 人員
type Person struct {
	ID   uint     `json:"id,omitempty"`   // 編號
	Name string   `json:"name,omitempty"` // 姓名
	Addr *Address `json:"addr,omitempty"` // 地址
}

// 成員列表
var members = make(map[string]Person)

// 入口
func main() {
	router := mux.NewRouter()

	// 假資料
	members["1"] = Person{ID: 1, Name: "A", Addr: &Address{City: "Taipei"}}
	members["2"] = Person{ID: 2, Name: "B", Addr: &Address{City: "Kaohsiung"}}

	// 註冊
	router.HandleFunc("/", getMembers).Methods("GET")
	router.HandleFunc("/person/{id}", getMember).Methods("GET")
	router.HandleFunc("/person/{id}", addMember).Methods("POST")
	router.HandleFunc("/person/{id}", delMember).Methods("DELETE")

	// 開始監聽
	log.Fatal(http.ListenAndServe(":9527", router))
}

// 取得所有成員
func getMembers(writer http.ResponseWriter, req *http.Request) {
	json.NewEncoder(writer).Encode(members)
}

// 取得成員
func getMember(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	if person, ok := members[id]; ok {
		json.NewEncoder(writer).Encode(person)
	}
}

// 新增成員
func addMember(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	if _, ok := members[id]; !ok {
		var person Person
		json.NewDecoder(req.Body).Decode(&person)
		members[id] = person
	}
}

// 刪除人員
func delMember(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	delete(members, id)
}
