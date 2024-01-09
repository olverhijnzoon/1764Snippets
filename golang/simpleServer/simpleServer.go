package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Elo  int    `json:"elo"`
}

var (
	store = make(map[int]Item)
	mu    sync.RWMutex
)

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		items := make([]Item, 0, len(store))
		for _, item := range store {
			items = append(items, item)
		}
		mu.RUnlock()

		json.NewEncoder(w).Encode(items)
	case http.MethodPost:
		var item Item
		json.NewDecoder(r.Body).Decode(&item)

		mu.Lock()
		item.ID = len(store) + 1
		store[item.ID] = item
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(item)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		item, ok := store[id]
		mu.RUnlock()

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(item)
	case http.MethodPut:
		var item Item
		json.NewDecoder(r.Body).Decode(&item)

		mu.Lock()
		store[id] = item
		mu.Unlock()
	case http.MethodDelete:
		mu.Lock()
		delete(store, id)
		mu.Unlock()
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/items/", itemHandler)

	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Simple Server")
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}
