package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	src "github.com/themycode/go-docker-env/src"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!

`)
}

type jsonStr struct {
	Env      string
	Language string
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	js := jsonStr{
		Env: "docker", Language: "golang",
	}
	data, error := json.Marshal(js)
	if error != nil {
		fmt.Println("error", error.Error())
	}

	mainConfig := src.LoadConfig("dev.json")
	fmt.Println("mainConfig", mainConfig.Address)

	fmt.Print(r.URL.RawQuery)
	fmt.Println("data:", string(data))
	fmt.Fprint(w, string(data))

}
func HttpReq(w http.ResponseWriter, r *http.Request) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/anything", nil)
	if err != nil {
		fmt.Println("get response err", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response err", err)
	}
	defer resp.Body.Close()
	context, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read context err ", err)
	}
	fmt.Println("context:", string(context))
	fmt.Fprint(w, string(context))
}

func main() {
	fmt.Println(src.Str("hello"))
	http.HandleFunc("/http", HttpReq)
	http.HandleFunc("/", handler)
	http.HandleFunc("/env", handlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// run RootCommand
	if err := src.RootCommand.Execute(); err != nil {
		fmt.Println("err:", os.Stderr, err)
		os.Exit(1)
	}
}
