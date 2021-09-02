package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func FetchInputData(day int) string {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2019/day/%d/input", day), nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("AOC_SESSION")})

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	text, _ := ioutil.ReadAll(resp.Body)

	return strings.TrimSpace(string(text))
}
