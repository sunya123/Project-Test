package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var uploadTemplate = template.Must(template.ParseFiles("index.html"))

func indexHandle(w http.ResponseWriter, r *http.Request) {
	if err := uploadTemplate.Execute(w, nil); err != nil {
		log.Fatal("Execute: ", err.Error())
		return
	}
}

func UploadHandle(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Fatal("FormFile:", err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			errHandle(w, r, 500)
			log.Fatal("Close:", err.Error())
			return
		}
	}()
	filename := r.FormValue("filename")
	var fp *os.File
	fp, err = os.OpenFile("/home/sunya/"+filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal("Open File Error")
		errHandle(w, r, 400)
	}
	defer fp.Close()
	_, err = io.Copy(fp, file)
	if err != nil {
		log.Fatal("ReadAll:", err.Error())
		http.Error(w, "internal error", 500)
		return
	}
	w.Write([]byte("{Opcode:Ok,msg:upload the file success}"))
}

func downloadHandle(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("name")
	log.Println(fileName)
	if len(fileName) > 0 {
		fp, err := os.OpenFile("/home/sunya/test3/"+fileName, os.O_RDWR, os.ModePerm)
		if err == nil {
			defer fp.Close()
			_, err = io.Copy(w, fp)
		} else {
			log.Println(err.Error())
		}

	}

	//w.Write([]byte(fileName))

}

func errHandle(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func main() {
	var dd map[int64]int64
	dd = make(map[int64]int64, 20)
	dd[int64(4)] = int64(98)
	dd[int64(1)] = int64(97)
	log.Println(dd[int64(1)])
	dd[10] = 89
	log.Println(dd[10])
	delete(dd, 10)
	dd[10] = 89
	dd[11] = 89
	for k, v := range dd {
		log.Println(k)
		log.Println(v)
	}

	//http.HandleFunc("/", indexHandle)
	//http.HandleFunc("/dbbak", downloadHandle)
	//http.HandleFunc("/upload", UploadHandle)
	//http.ListenAndServe("localhost:8080", nil)

}
