package webui

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ARC5RF/DiscordAnalyticsBot/webui/hub"
)

// func Index(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("HTTP %s %s%s\n", r.Method, r.Host, r.URL)

// 	log.Println(r.URL)
// 	if r.URL.Path != "/" {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "text/html")
// 	http.ServeFile(w, r, "index.html")

// 	w.Write([]byte(time.Now().Local().Format(time.RFC1123Z)))
// }

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	bp := filepath.Join("assets", "templates")
	lp := filepath.Join(bp, "layout.html")
	fp := filepath.Join(bp, "pages", filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	if info.IsDir() {
		fmt.Println("got dir")
		a := filepath.Join(fp, "index.html")
		fmt.Println(a)
		if i, err := os.Stat(a); !os.IsNotExist(err) {
			info = i
			fp = a
		}
		b := fp + ".html"
		fmt.Println(b)
		if i, err := os.Stat(b); !os.IsNotExist(err) {
			info = i
			fp = b
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	f := []string{lp, fp}
	m, err := filepath.Glob(filepath.Join(bp, "components") + string(os.PathSeparator) + "*.html")
	if err != nil {
		// Log the detailed error
		log.Print(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if len(m) > 0 {
		fmt.Println(m)
		f = append(f, m...)
	}

	tmpl, err := template.ParseFiles(f...)
	if err != nil {
		// Log the detailed error
		log.Print(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

type WebUIInstance struct {
	server *http.Server
	hub    *hub.Hub
}

func New() *WebUIInstance {
	self := &WebUIInstance{}
	self.server = &http.Server{
		Addr: ":3000",
	}
	self.hub = hub.New()
	return self
}

func (self *WebUIInstance) Start() {
	log.Println("Starting our simple http server.")

	go self.hub.Run()

	fs := http.FileServer(http.Dir(filepath.Join("assets", "public")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/events", self.hub.ServeWs)
	http.HandleFunc("/", ServeTemplate)

	log.Printf("Listening at http://localhost%s", ":3000")

	if err := self.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}
}

func (self *WebUIInstance) Receive(c func(m *hub.ClientCommand)) {
	self.hub.Receive(c)
}

func (self *WebUIInstance) Broadcast(message []byte) {
	self.hub.Broadcast(message)
}

func (self *WebUIInstance) Stop() {
	if err := self.server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}
}
