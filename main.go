package simple_amzn_url

import (
	"fmt"
	"regexp"
	"errors"
)

func RemoveUrl(url string) (string,  error) {
	str := url
	if str == "" {
		return "", errors.New("urlが指定されていません")
	}

	encodeStr := regexp.MustCompile(`%2F`)
	str = encodeStr.ReplaceAllString(str, "/")

	removeBookNameRep := regexp.MustCompile(`/%[a-zA-Z0-9%-]*`)
	removeBookNameStr := removeBookNameRep.ReplaceAllString(str, "")

	rep := regexp.MustCompile(`/ref[a-zA-Z0-9%&=_?-]*`)
	compStr := rep.ReplaceAllString(removeBookNameStr, "")
	fmt.Println(compStr)
	return compStr, nil


	fmt.Println(removeBookNameStr)





	return "", nil
}

