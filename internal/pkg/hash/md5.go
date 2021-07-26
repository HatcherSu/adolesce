package hash

import (
	"crypto/md5"
	"fmt"
	"os"
)

func MD5Hash(str string) string {
	hash := md5.New()
	if _, err := hash.Write([]byte(str)); err != nil {
		fmt.Printf("hash write error :%s", err.Error())
		os.Exit(-1)
	}
	result := hash.Sum(nil)
	return fmt.Sprintf("%x", result)
}
