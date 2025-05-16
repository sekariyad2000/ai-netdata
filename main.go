package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// ============ MAIN FUNCTION ============

func main() {
	// Serve static files (CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ask", askHandler)

	fmt.Println("Server gestart op http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

// ============ INDEX HANDLER ============

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// ============ ASK HANDLER ============

func askHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	question := r.FormValue("question")
	metrics, err := GetNetdataMetrics()
	if err != nil {
		http.Error(w, "Error getting metrics", 500)
		return
	}

	fullPrompt := fmt.Sprintf("System data:\n%s\n\nUser question: %s", metrics, question)
	response, err := AskOllama(fullPrompt)
	if err != nil {
		http.Error(w, "Error calling Ollama", 500)
		return
	}

	w.Write([]byte(response))
}
