package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetNetdataMetrics() (string, error) {
	resp, err := http.Get("http://localhost:19999/api/v1/allmetrics?format=json&names=short")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var raw map[string]interface{}
	json.Unmarshal(body, &raw)

	var sb strings.Builder
	count := 0
	for k, v := range raw {
		sb.WriteString(k + ": ")
		switch t := v.(type) {
		case map[string]interface{}:
			for sk, sv := range t {
				sb.WriteString(sk + "=" + toString(sv) + " ")
			}
		default:
			sb.WriteString(toString(v))
		}
		sb.WriteString("\n")
		count++
		if count > 50 { // alleen eerste 50 metrieken om de prompt niet te groot te maken
			break
		}
	}

	return sb.String(), nil
}

func toString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
