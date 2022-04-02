package main

import (
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"log"

	"github.com/go-openapi/runtime/middleware"

	"github.com/vorsprung/soundslike"
	"github.com/vorsprung/soundslike/gen/restapi"
	"github.com/vorsprung/soundslike/gen/restapi/operations"
)

var word = flag.String("word", "superb", "synonym word")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewSoundslikeAPI(swaggerSpec)
	server := restapi.NewServer(api)
	var o soundslike.Phonic
	o.Load("words.txt")
	fmt.Printf("%v\n", len(o.Setup))
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = 3000

	// TODO: Set Handle
	// listWordsHandler greets the given name,
	// in case the name is not given, it will default to World
	api.ListWordsHandler = operations.ListWordsHandlerFunc(
		func(params operations.ListWordsParams) middleware.Responder {
			myword := params.Word.(map[string]interface{})["word"]
			fmt.Printf("%s\n", myword)
			wordlist := o.SoundsLike(myword.(string))
			return operations.NewListWordsOK().WithPayload(wordlist)
		})
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
