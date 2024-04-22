package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("juan")
	fmt.Println(msg)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("juan") = %q, %v, quiere coincidencia para %#q,nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "",error`, msg, err)
	}

}
