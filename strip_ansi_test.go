package stripansi

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	foo            = "\x1b[30mfoo\x1b[0m"
	bar            = "\x1b[36mbar\x1b[0m"
	baz            = "\x1b[96mbaz\x1b[0m"
	withUnderlines = "\u001B[31m\u001B[4mi like cake\u001B[24m\u001B[39m"
	inverse        = "\u001B[7minverse\u001B[27m"
	jim            = "\u001B[31mJ\u001B[39m\u001B[33mi\u001B[39m\u001B[32mm\u001B[39m \u001B[35mS\u001B[39m\u001B[31mc\u001B[39m\u001B[33mh\u001B[39m\u001B[32mu\u001B[39m\u001B[34mb\u001B[39m\u001B[35me\u001B[39m\u001B[31mr\u001B[39m\u001B[33mt\u001B[39m"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{name: "strips single escapes", src: foo, want: "foo"},
		{name: "strips multiple escapes", src: fmt.Sprintf("%s %s %s", foo, bar, baz), want: "foo bar baz"},
		{name: "strips underlines", src: withUnderlines, want: "i like cake"},
		{name: "strips inverse", src: inverse, want: "inverse"},
		{name: "strips ansi complex", src: jim, want: "Jim Schubert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.src); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes(t *testing.T) {
	tests := []struct {
		name string
		src  []byte
		want []byte
	}{
		{name: "strips single escapes", src: []byte(foo), want: []byte("foo")},
		{name: "strips multiple escapes", src: []byte(fmt.Sprintf("%s %s %s", foo, bar, baz)), want: []byte("foo bar baz")},
		{name: "strips underlines", src: []byte(withUnderlines), want: []byte("i like cake")},
		{name: "strips inverse", src: []byte(inverse), want: []byte("inverse")},
		{name: "strips ansi complex", src: []byte(jim), want: []byte("Jim Schubert")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes(tt.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
