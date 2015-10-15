package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const MBSPURL = "http://172.16.1.29:5000/"

func newAgent() *Agent {
	return &Agent{
		Memories: make(map[time.Time]Event)}
}

func (t *Agent) respond(input string) string {

	input = strings.ToLower(input)
	re := regexp.MustCompile("[^a-z]")
	input = re.ReplaceAllLiteralString(input, " ")
	input = strings.TrimSpace(input)

	// t.Memories[time.Now()] = Event{What: input}

	//wurds := strings.Split(input, " ")
	//for _, w := range wurds {
	//	nl, _ := regexp.MatchString("^$", w)
	//	if nl {
	//		continue
	//	}
	//	//t.WurdMap[w] += 1
	//}

	//fmt.Println(t.Memories)
	MBSPnalyse(input)

	return input
}

func MBSPnalyse(s string) {
	var jsonnn = "{\"wurds\": \"" + strings.TrimSpace(s) + "\"}"
	var jsonStr = []byte(jsonnn)
	req, err := http.NewRequest("POST", MBSPURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Sideb0ard-Service", "v1")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	//if b.Debug {
	//	fmt.Println(jsonnn)
	//	fmt.Println("response Status:", resp.Status)
	//	fmt.Println("response Headers:", resp.Header)
	//	fmt.Println("response Body:", string(body))
	//}
	fmt.Println(string(body))
}
