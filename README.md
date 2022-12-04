### How to run and call for soundslike
# Server

    $ go run cmd/soundslike/main.go 
    1
    2021/01/01 15:20:08 Serving soundslike at http://[::]:3000
    shoeshop
    shoeshop encodes to XXP
    lookup table has 176630 keys

# Client

send this

    curl --header "Content-Type: application/json" -XPOST -d '{ "word": "shoeshop"} ' http://localhost:3000/phononym

and the response is

    ["chichipe","shoeshop","showshop"]
