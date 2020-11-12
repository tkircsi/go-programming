package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type dollars float32
type database map[string]dollars

const listHTML = `<html>
<body>
	<table border="1">
	<tr>
		<th>Item</th><th>Price</th>
	</tr>
	{{range $key, $value := .}}
	<tr>
		<td>{{ $key }}</td><td> {{ $value }}</td>
	</tr>
	{{end}}
	</table>
</body>
</html>`

var temp, tempErr = template.New("list").Parse(listHTML)

func main() {
	db := database{"milk": 12, "apple": 3, "meat": 22, "hamburger": 3.5, "hot-dog": 2.2}
	mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price)) OR
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/delete", db.delete)
	mux.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, r *http.Request) {
	if tempErr == nil {
		w.WriteHeader(http.StatusOK)
		temp.Execute(w, db)
		return
	}

	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if len(item)*len(price) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invaid item or price\n")
		return
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item: %q already exists\n", item)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %s\n", price)
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "item: %q with price: %s added\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if len(item) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item name is required")
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "item %q is deleted", item)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if len(item)*len(price) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invaid item or price\n")
		return
	}

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item: %q does not exist\n", item)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %s\n", price)
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "item: %q with price: %s updated\n", item, db[item])
}
