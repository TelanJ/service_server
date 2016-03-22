package main


import (
    "github.com/gorilla/mux"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
    "github.com/telanj/service"
    "log"
    "net/http"
)

type App string

type Args struct {
	Name string
	ID int
}

type Reply string

func (t *App) Greet(args *Args, reply *Reply) error {
	*reply = "Greetings!" + args.Name + "with id" + args.id
	return nil
}

func main() {
	s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")
    s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
    greet := new(Greet)
    s.RegisterService(greet, "")
    http.ListenAndServe("127.0.0.1:10001", nil)
}