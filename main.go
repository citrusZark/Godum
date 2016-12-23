// main
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"servertest/routers"

	"servertest/controllers"

	"github.com/codegangsta/negroni"
)

var addr *string
var n *negroni.Negroni

//var Delay *int

func main() {
	addr = flag.String("addr", ":8080", "The addr of the application.")
	delay := flag.Duration("delay", 0, "Delay Request.")
	status := flag.Int("status", 200, "HTTP Response Status.")
	isLog := flag.Bool("log", false, "Print Log?")
	isPrintVal := flag.Bool("printval", false, "Print JSON Value?")
	filename := flag.String("datafile", "", "filename.")
	flag.Parse() // parse the flags

	var fixedData []byte
	var isFixedData bool
	if *filename == "" {
		isFixedData = false
	} else {
		isFixedData = true
	}
	controllers.SetDelay(*delay)
	controllers.SetStatus(*status)
	controllers.SetIsPrintVal(*isPrintVal)
	controllers.SetIsFixedData(isFixedData)
	if isFixedData {
		dataByte, err := ioutil.ReadFile(*filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		fixedData = dataByte
		controllers.SetFixedData(fixedData)
	}

	//get mux router object
	router := routers.InitRoutes()

	//create negroni instance
	if *isLog {
		n = negroni.Classic()
	} else {
		n = negroni.New()
	}
	//n := negroni.New()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    *addr,
		Handler: n,
	}
	fmt.Println("...GoDumm : Dummy REST Server Simulator")
	fmt.Println("...by Pangeran Muhammad Thoha")
	fmt.Println("...github.com/carterpillar")
	fmt.Println("...usage example : godumm -addr localhost:8080 -delay 5s -status 200 -log -printval")
	fmt.Println("... -log for display server log")
	fmt.Println("... -printval for display json value")
	log.Printf("Listening at %s with delay %s.....", server.Addr, *delay)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server Start Error:", err)
		fmt.Println("Program exiting")
		return
	}
}
