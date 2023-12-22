package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	// parse the raw url into structured url
	serverUrl, err := url.Parse(addr)

	handleErr(err)

	return &simpleServer{
		addr: addr,
		// redirect the url
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func (s *simpleServer) Address() string {
	// return the address of the simple server
	return s.addr
}

func (s *simpleServer) IsAlive() bool {

	// this should return alive or not in the real world
	return true
}

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

// get the next available server
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) serverProxy(rw http.ResponseWriter, req *http.Request) {

	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding Request to address %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func main() {

	servers := []Server{
		newSimpleServer("https://duckduckgo.com/"),
	}

	lb := NewLoadBalancer("8080", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serverProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving requests at localhost:%s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
