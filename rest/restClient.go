package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAQuote() {
	resp, err := http.Get("http://api.theysaidso.com/qod.json")

	if err != nil {
		fmt.Println("Error1", err)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}
