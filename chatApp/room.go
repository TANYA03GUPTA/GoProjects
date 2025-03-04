package main

import (
	"chatApp/trace"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/stretchr/objx"
)

type room struct{
	clients map[*client]bool
	join chan *client
	leave chan *client
	forward chan *message
	tracer trace.Tracer
}

func newRoom() *room{
	return &room{
		forward: make(chan *message),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room)run(){
	for{
		select{
			//listening to diff channels
		case client := <- r.join:
			//mapping their values int he channel
			r.clients[client] = true
			r.tracer.Trace("New Clients Joined")
		case client := <- r.leave:
			delete(r.clients, client)
			close(client.receive)
			r.tracer.Trace("Client left")
		case msg := <- r.forward:
			for client := range r.clients{
				select{
				case client.receive <- msg: 
				    r.tracer.Trace(" -- sent to client")
				default:
				//fail o send
				delete(r.clients,client)
		        close(client.receive)	
		        r.tracer.Trace("Client left")
			}	
				}
		}
	}
}

const (
	socketBufferSize =  1024
	MessageBufferSize = 256
)

///to upgrade our normal http connection to socket sp it persists and soes not end
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize , WriteBufferSize: socketBufferSize}
//server func to call the upgraded socket and push to it
func (r *room)ServeHTTP(w http.ResponseWriter, req *http.Request){
	socket,err := upgrader.Upgrade(w,req, nil)
	if err != nil{
       log.Fatal("serveHttp :",err)
	   return 
	}
    
	authCookie,err := req.Cookie("auth")
	if err != nil {
		log.Fatal("fialed to get auth cookie:",err)
		return
	}

	client := &client{
		socket: socket,
		receive: make(chan *message,MessageBufferSize),
		room : r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() {r.leave <- client	}()
	go client.write()
	//we make this a go routine because the write func will keep waiting for a req to arrive INF 
	//therefore to make sure that we call it asoson as req arrives we call a gorouitne
	client.read()
}