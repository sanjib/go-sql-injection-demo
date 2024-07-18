package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

var (
	db *sql.DB
)

func main() {
	// Open db
	err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Mux and serve
	http.HandleFunc("GET /", home)
	http.HandleFunc("POST /", homePost)
	fmt.Println("starting server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type homeTemplateData struct {
	PostMessage string
	Messages    []Message
}

func home(w http.ResponseWriter, r *http.Request) {
	msgs, err := getMessages()
	if err != nil {
		log.Println(err)
		http.Error(w, "error: get messages", http.StatusInternalServerError)
		return
	}

	tmplData := homeTemplateData{Messages: msgs}
	tmpl := template.Must(template.New("base").ParseFiles("home.tmpl"))
	tmpl.Execute(w, tmplData)
}

func homePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "error: parse form", http.StatusInternalServerError)
		return
	}
	msg := r.PostForm.Get("msg")

	err = insertMessage(msg)
	if err != nil {
		log.Println(err)
		http.Error(w, "error: insert msg", http.StatusInternalServerError)
		return
	}

	msgs, err := getMessages()
	if err != nil {
		log.Println(err)
		http.Error(w, "error: get messages", http.StatusInternalServerError)
		return
	}

	tmplData := homeTemplateData{PostMessage: msg, Messages: msgs}
	tmpl := template.Must(template.New("base").ParseFiles("home.tmpl"))
	tmpl.Execute(w, tmplData)
}

type Message struct {
	ID      int
	Message string
}

// Example SQL injection (include space after -- below) code:
// '); truncate messages; --
func insertMessage(msg string) error {
	// Allows SQL injection.
	// Uncomment following 2 lines to demo SQL injection.
	//q := fmt.Sprintf(`INSERT INTO messages (message) VALUES ('%s')`, msg)
	//_, err := db.Exec(q)

	// Prevents SQL injection.
	// Uncomment following 2 lines to prevent SQL injection.
	q := `INSERT INTO messages (message) VALUES (?)`
	_, err := db.Exec(q, msg)

	log.Println(q)
	if err != nil {
		return err
	}
	return nil
}

func getMessages() ([]Message, error) {
	q := "SELECT * FROM messages"
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	var msgs []Message
	for rows.Next() {
		var m Message
		err = rows.Scan(&m.ID, &m.Message)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func openDB() error {
	var err error
	db, err = sql.Open("mysql", "root@/va_test1?parseTime=true&multiStatements=true")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
