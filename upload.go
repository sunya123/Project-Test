package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	//"io/ioutil"
	"log"
)

func upload(url, file string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(file)
	if err != nil {
		return
	}
	fw, err := w.CreateFormFile("file", file)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	nw, err := w.CreateFormField("filename")
	if _, err := nw.Write([]byte("ddd.tar.gz")); err != nil {
		return
	}

	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return

}

func download(url, file string) {
	fp, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer fp.Close()
	var r *http.Request
	r, err = http.NewRequest("GET", url, nil)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return
	}
	defer res.Body.Close()
	//	var data []byte
	//	data,err=ioutil.ReadAll(res.Body)
	//	fp.Write(data)
	//	if err==nil{
	//		log.Println(string(data))
	//	}

	_, err = io.Copy(fp, res.Body)
	if err != nil {
		log.Println(err.Error())
	}

}

func main() {
	url := "http://localhost:8080/dbbak?name=leveldb.tar.gz"
	download(url, "/home/sunya/test2/level.tar.gz")

	//upload("http://localhost:8080/upload","/home/sunya/test2/1/db/000006.log")
}
