package tests

import (
	"proto-game-server/router"
	"reflect"
	"testing"
)

type PatternCase struct {
	inputPattern    string
	compiledPattern string
}

type ParseCase struct {
	pattern string
	url     string
	values  map[string]string
}

func TestPattern(t *testing.T) {
	cases := []PatternCase{
		{
			"/user/{id}/test/{val}",
			`\/user\/([^\/]+)\/test\/([^\/]+)`,
		},
		{
			"/user",
			`\/user`,
		},
	}

	for _, c := range cases {
		parser, err := router.NewApiUrlParser(c.inputPattern)
		if err != nil {
			t.Fatal(err)
		}

		compiledPattern := parser.Pattern()
		if compiledPattern != c.compiledPattern {
			t.Fatalf("Expected: %v\nGot: %v", c.compiledPattern, compiledPattern)
		}
	}
}

func TestParsing(t *testing.T) {
	cases := []ParseCase{
		{
			"/user/{id}",
			"/user/2",
			map[string]string{
				"id": "2",
			},
		},
		{
			"/user/{id}/value/{v}",
			"/user/42/value/ctulhu",
			map[string]string{
				"id": "42",
				"v":  "ctulhu",
			},
		},
	}

	for _, c := range cases {
		parser, err := router.NewApiUrlParser(c.pattern)
		if err != nil {
			t.Fatal(err)
		}

		if !parser.Match(c.url) {
			t.Fatalf("URL dosent match to: %v", parser.Pattern())
		}

		urlValues := parser.Parse(c.url)
		if !reflect.DeepEqual(c.values, urlValues) {
			t.Fatalf("Expected: %v\nGot: %v", c.values, urlValues)
		}
	}
}
