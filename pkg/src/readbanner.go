package src

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
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

func HashValid(banners, pathOfBanners string) error {
	hashesOfBanners := map[string]string{
		"shadow":     "a49d5fcb0d5c59b2e77674aa3ab8bbb1",
		"standard":   "a51f800619146db0c42d26db3114c99f",
		"thinkertoy": "8efd138877a4b281312f6dd1cbe84add",
	}

	if hashesOfBanners[banners] != Md5sum(pathOfBanners) {
		return fmt.Errorf("you have changed a file")
	}
	return nil
}

func Md5sum(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}
