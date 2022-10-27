package src

import (
	"errors"
	"os"
	"strings"
)

func ReadArgs(input string, symbols map[int][]string) string {
	art := ""
	input = strings.ReplaceAll(input, "\\n", "\n")
	buf := make([]string, 8)
	artBuf := false
	for k := 0; k < len(input); k++ {
		if input[k] == '\n' {
			art = makeStr(buf, art)
			buf = make([]string, 8)
			continue
		}
		artBuf = true
		for i, r := range symbols[int(input[k])] {
			buf[i] += r
		}
		if k+1 >= len(input) {
			artBuf = false
			art = makeStr(buf, art)
			break
		}
	}
	if artBuf {
		art += "\n"
	}
	return art[:len(art)-1]
}

func makeStr(buf []string, s string) string {
	if buf[0] == "" {
		return s + "\n"
	}
	for _, v := range buf {
		s += v + "\n"
	}
	return s
}

func IsvalidArgs(s []string) error {
	if len(s) != 3 {
		err := errors.New("the arguments should be : go run . [STRING] [BANNER]")
		return err
	}

	if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
		err := errors.New("[BANNER] should be : standard, shadow, thinkertoy")
		return err
	}
	return nil
}
