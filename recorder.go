package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
  "flag"
)

func main() {
  var port = *flag.Int("port", 9090, "the port to listen for githook messages on")
  flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("usage:\n\tgithook-recorder [--port=9090] <session name>\n")
    flag.PrintDefaults()
		return
	}

	var playthoughName = flag.Arg(0)
	var githookSequence = 0
  var addr = fmt.Sprintf(":%d", port)
	fmt.Println("recording session named", playthoughName, "on", addr)

	if _, err := os.Stat(playthoughName); os.IsExist(err) {
		fmt.Println("session", playthoughName, "already exists.")
		return
	}

	if err := os.Mkdir(playthoughName, 0755); err != nil {
		fmt.Println("could not create session dir:", err)
		return
	}

	http.HandleFunc("/githook", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			io.WriteString(res, "thanks")
		} else {
			fmt.Println("received payload #", githookSequence)
			ioutil.WriteFile(fmt.Sprintf("%s/%d.json", playthoughName, githookSequence), []byte(req.FormValue("payload")), 0755)
			githookSequence++
		}
	})

	http.ListenAndServe(addr, nil)
}
