package main

import "net/http"

func main() {
	http.HandleFunc("/compliance/sanctions", sanctions)
	http.HandleFunc("/compliance/ask_user", askUser)
	http.HandleFunc("/compliance/fetch_info", fetchInfo)
}

func fetchInfo(writer http.ResponseWriter, r *http.Request) {

}

func askUser(writer http.ResponseWriter, r *http.Request) {

}

func sanctions(writer http.ResponseWriter, r *http.Request) {

}
