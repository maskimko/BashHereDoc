package pkg

import (
	"reflect"
	"testing"
)

func TestParseHereDocs(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want map[string][]byte
	}{
		{name: "one heredoc",
			args: args{content: []byte(`interactive-program <<LimitString
command #1
command #2
...
LimitString`)},
			want: map[string][]byte{"LimitString": []byte(`command #1
command #2
...`)}},
		{name: "three heredocs",
			args: args{content: []byte(`interactive-program <<LimitString
command #1
command #2
...
LimitString
alksdj klahhasd
adslfjlk
alskdhf00wl
another line <<TKN
useful content
in
a
few
lines
TKN end of useful data
a;ldskjfa
asdhhrh
<<lowercase_token
lowercase token should work too
  even indentation should be preserved
lowercase_token
other stuff here`)},
			want: map[string][]byte{"LimitString": []byte(`command #1
command #2
...`),
				"TKN": []byte(`useful content
in
a
few
lines`),
				"lowercase_token": []byte(`lowercase token should work too
  even indentation should be preserved`)}},
		{name: "no heredoc",
			args: args{content: []byte(`adskjfk
akdsjfh
htbt0zojc`)},
			want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseHereDocs(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHereDocs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseHereDoc(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "one heredoc",
			args: args{content: []byte(`interactive-program <<LimitString
command #1
command #2
...
LimitString`)},
			want: []byte(`command #1
command #2
...`)},
		{name: "one tab indented heredoc",
			args: args{content: []byte(`some text
	some indented text <<-EOC
line #1
line #2
line #3
	EOC
	tab indented text`)},
			want: []byte(`line #1
line #2
line #3`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseHereDoc(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHereDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseHereDocString(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "one heredoc",
			args: args{content: `interactive-program <<LimitString
command #1
command #2
...
LimitString`},
			want: `command #1
command #2
...`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseHereDocString(tt.args.content); got != tt.want {
				t.Errorf("ParseHereDocString() = %v, want %v", got, tt.want)
			}
		})
	}
}
