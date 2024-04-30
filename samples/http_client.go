package samples

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slog"
	"net/http"
)

func HttpClient() {
	slog.Info("======> HTTP Client")
	slog.Info("HTTP client is a client that is able to send a request to and get a response from the server using HTTP protocol.")
	slog.Info("")

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

}
