package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	/*
		Examples
		localhost:8000/list
		localhost:8000/read?item=socks
		localhost:8000/read?item=pants
		localhost:8000/update?item=socks&price=75
		localhost:8000/update?item=socks&price=75a
		localhost:8000/update?item=pants&price=75
		localhost:8000/create?item=pants&price=60
		localhost:8000/create?item=pants&price=60a
		localhost:8000/create?item=hat&price=60a
		localhost:8000/delete?item=pants
		localhost:8000/delete?item=hat
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/read", db.read)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	mux.HandleFunc("/create", db.create)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	price_updated, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}
	db[item] = dollars(price_updated)
	fmt.Fprintf(w, "%s: %s\n", item, dollars(price_updated))
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "deleted item %s: %s\n", item, price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusAlreadyReported) // 208
		fmt.Fprintf(w, "item already exist: %q\n", item)
		return
	}
	price_created, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}
	db[item] = dollars(price_created)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s: %s\n", item, dollars(price_created))
}
