package webui

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ARC5RF/DiscordAnalyticsBot/webui/hub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var Broadcast = make(chan any)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP %s %s%s\n", r.Method, r.Host, r.URL)

	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")

	w.Write([]byte(time.Now().Local().Format(time.RFC1123Z)))
}

func Events(w http.ResponseWriter, r *http.Request) {

}

type WebUIInstance struct {
	s *http.Server
	h *hub.Hub
}

func New() *WebUIInstance {
	self := &WebUIInstance{}
	self.s = &http.Server{
		Addr: ":3000",
	}
	self.h = hub.New()
	return self
}

func (self *WebUIInstance) Start() {
	log.Println("Starting our simple http server.")

	go self.h.Run()

	http.HandleFunc("/", Index)
	http.HandleFunc("/events", self.h.ServeWs)

	log.Printf("Listening at http://%s", ":3000")

	if err := self.s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}
}

func (self *WebUIInstance) Broadcast(message []byte) {
	self.h.Broadcast(message)
}

func (self *WebUIInstance) Stop() {

	if err := self.s.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}
}
