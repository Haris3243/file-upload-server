package main

import(
	"fmt"
	"os"
	"io"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

    //upload size
    err := r.ParseMultipartForm(200000) // grab the multipart form
    if err != nil {
        fmt.Fprintln(w, err)
    }
    
    //reading original file
    file, handler, err := r.FormFile("originalFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    
	resFile, err := os.Create("./data/"+handler.Filename)
	if err != nil {
		fmt.Fprintln(w, err)
	}
    defer resFile.Close()
    if err==nil {
    io.Copy(resFile,file)
    defer resFile.Close()
    fmt.Fprintf(w, "Successfully Uploaded Original File\n")
    }
}
//start server which will listen and server on pot 8083
func startServer() {

    http.HandleFunc("/uploadFile", uploadFile)
    
    //server is listening on port 8083
    http.ListenAndServe(":8083", nil)
}

func main() {
    
	fmt.Println("Hello World!")
	startServer()
}