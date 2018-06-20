package main

import (
	"github.com/ananagame/Olivia/supports"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", nil)
	supports.ChooseSupport()
}
