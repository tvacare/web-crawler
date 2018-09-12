package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// MatchPage match responde according to regex passed
func MatchPage(resp *http.Response, regex string) ([][]string, error) {
	b, e := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if e != nil {
		return nil, e
	}

	defer resp.Body.Close()
	return regexp.MustCompile(regex).FindAllStringSubmatch(string(b), -1), nil
}

// SliceContains verify if contains any string in slice passed
func SliceContains(s string, properties []string) (contain bool, prop string) {
	for _, p := range properties {
		if strings.Contains(s, p) {
			return true, p
		}
	}
	return false, ""
}
