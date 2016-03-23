package main


import (
    "net/rpc"
    "net/http"
    "net"
    "log"
    "github.com/telanj/service"
)


type Service string

type Args struct {
    Name string
}

type Reply string

func (t *Service) Greet(args *Args, reply *Reply) error {
    *reply = Reply("Greetings " +args.Name +"!")
    return nil
}


func main() {

    service := new(test.Service)
    err := rpc.Register(service)
    if err != nil {
        log.Fatalf("Format of service isn't correct. %s", err)
    }
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":10001")
    if e != nil {
        log.Fatalf("Couldn't start listening on port 10001. Error %s", e)
    }
    log.Println("Serving RPC handler")
    err = http.Serve(l, nil)
    if err != nil {
        log.Fatalf("Error serving: %s", err)
    }
}