//Package sieve主要是处理
package sieve

import (
	"regexp"
)

type Sie interface {
	DealWith(src string) (string, error)
}

//number0
type Number0 struct {
	src string
}

//
func (n0 Number0) DealWith(src string) (string, error) {
	_, err := regexp.MatchString(Pattern["number0"], src)
	if err == nil {
		return src, err
	} else {
		return "error", err
	}
}

type Number1 struct {
	src string
}

//
func (n1 Number1) DealWith(src string) (string, error) {
	_, err := regexp.MatchString(Pattern["number1"], src)
	if err == nil {
		return src, err
	} else {
		return "error", err
	}
}
