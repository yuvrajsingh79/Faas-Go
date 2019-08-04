package main

import "net/http"

func main() {
	http.HandleFunc("/", GetAllK8sresources)
}
