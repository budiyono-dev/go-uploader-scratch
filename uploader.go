package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
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
	//doRequest()
	type data struct {
		
	}
	dat, err := ioutil.ReadFile("./conf.json")
	json.Decode()
	check(err)
	fmt.Println(string(dat))
}
