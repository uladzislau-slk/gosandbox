package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://google.com", nil)
	request.Header.Add("Accept", "text/html")
	request.Header.Add("User-Agent", "MSIE/15.0")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}
