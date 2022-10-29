package pkg

import (
	"main/pkg/src"
	"net/http"
)

func AsciiPrint(asciiTxt, banner string) (string, int) {
	pathOfBanners := "pkg/banners/" + banner + ".txt"
	// err := src.IsvalidArgs(input)
	// if err != nil {

	// 	log.Println(err)
	// 	return "", 400
	// }
	args := asciiTxt
	err1 := src.Isvalid(args)
	if err1 != 0 {
		return "", http.StatusBadRequest
	}
	symbols, err2 := src.ReadBanner(banner, pathOfBanners)
	if err2 != 0 {
		return "", http.StatusInternalServerError
	}
	Donetxt := src.ReadArgs(args, symbols)
	return Donetxt, 0
}
