package src

import (
	"bufio"
	"net/http"
	"os"
)

func ReadBanner(banners string, pathOfBanners string) (map[int][]string, int) {
	symbols := make(map[int][]string)
	var buf []string
	counter := 0
	key := 31

	file, err := os.Open(pathOfBanners)
	if err != nil {
		return map[int][]string{}, http.StatusInternalServerError
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			counter = 0
			key++
			buf = nil
			continue
		}
		buf = append(buf, scanner.Text())
		// fmt.Printf("[%v]\n",buf)
		counter++
		if counter == 8 {
			symbols[key] = buf
		}
	}
	return symbols, 0
}

func Isvalid(args string) int {
	for _, r := range args {
		if (r < 32 || r > 127) && r != 10 && r != '\r' {
			return http.StatusBadRequest
		}
	}
	return 0
}
