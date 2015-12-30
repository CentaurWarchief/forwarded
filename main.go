package forwarded

import (
	"net/http"
	"regexp"
)

// Parse parses the `Forwarded` header and return all of its information about
// the forwarding
func Parse(r *http.Request) *Forwarded {
	forwarded := r.Header.Get("Forwarded")

	if forwarded == "" {
		return nil
	}

	regex := regexp.MustCompile("(?P<parameter>for|by|host|proto)=(?P<value>.[^;]+)")

	f := &Forwarded{}

	for _, match := range regex.FindAllStringSubmatch(forwarded, -1) {
		val := string(match[2])

		switch string(match[1]) {
		case "proto":
			f.Proto = val
			break
		case "by":
			f.By = val
			break
		case "for":
			f.For = val
			break
		case "host":
			f.Host = val
			break
		}
	}

	return f
}

// Was returns whether the request was forwarded by a proxy
func Was(r *http.Request) bool {
	return r.Header.Get("Forwarded") != ""
}
