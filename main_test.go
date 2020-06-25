package main

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestRunRequiresOneArgument(t *testing.T) {
	is := is.New(t)

	for _, args := range [][]string{
		{},
		{"16", "16"},
		{"16", "16", "16", "16", "16", "16", "16", "16", "16", "16", "16"},
	} {
		err := run(args, nil)
		is.True(err != nil) // run returns an error
		is.True(strings.HasPrefix(err.Error(), "usage: "))
	}
}

func TestRunRequiresNumericalArg(t *testing.T) {
	is := is.New(t)

	for _, arg := range []string{
		"a", "", "-", "?", "a9", "9a",
	} {
		err := run([]string{arg}, nil)
		is.True(err != nil) // run returns an error
		is.True(strings.HasPrefix(err.Error(), "usage: "))
	}
}

func TestRun(t *testing.T) {
	is := is.New(t)

	for i := 0; i < 128; i++ {
		var b bytes.Buffer
		err := run([]string{strconv.Itoa(i)}, &b)

		is.NoErr(err) // should not return an error
		res, err := base64.StdEncoding.DecodeString(b.String())
		is.NoErr(err) // base64-decoding output failed

		is.Equal(len(res), i)
	}
}
