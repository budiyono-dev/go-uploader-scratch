package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getMd5(content []byte) string {
	hashByte := md5.Sum(content)
	hashString := hex.EncodeToString(hashByte[:])
	return hashString
}

func getBase64(content []byte) string {
	base64Str := base64.StdEncoding.EncodeToString(content)
	return base64Str
}

func doRequest() {
	req, err1 := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err1 != nil {
		// handle error
	}
	//req.Header.Set("Content-Type", "application/json")
	req.Header = http.Header{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{"application/json"},
	}
	client := &http.Client{}
	res, err2 := client.Do(req)

	if err2 != nil {
		// handle error
	}
	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	fmt.Println("body", string(body))

}

func main() {
	fmt.Println("mulai")
	// duration := time.Duration(10)*time.Second
	// time.Sleep(duration)

	// dat, err := ioutil.ReadFile("./doc/hello.txt")
	// check(err)
	// fmt.Print(string(dat))
	//var uuidWithHyphen string
	//uuidWithHyphen = uuid.New().String()
	//fmt.Println(reflect.TypeOf(uuidWithHyphen))
	//files, err2 := ioutil.ReadDir("./doc/")
	//check(err2)
	//// fmt.Print(files)
	//for _, file := range files {
	//	fmt.Println(file.Name())
	//	dat, err := ioutil.ReadFile("./doc/" + file.Name())
	//	check(err)
	//	// fmt.Println("isi file ",string(dat))
	//	resultHash := getMd5(dat)
	//	fmt.Println("", resultHash)
	//	resultBase64 := getBase64(dat)
	//	fmt.Println("64 : ", resultBase64)
	//
	//}
	doRequest()

}
