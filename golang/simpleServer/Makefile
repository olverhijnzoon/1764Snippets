BINARY=simpleServer

run_go: build_go
	./$(BINARY)X

build_go:
	go build -o $(BINARY)X $(BINARY).go

test_go:
	curl http://localhost:8080/items
	echo "# Creating item"
	curl -X POST -H "Content-Type: application/json" -d '{"name":"test item", "elo":1764}' http://localhost:8080/items
	echo "# Fetching all items"
	curl http://localhost:8080/items
	echo "# Updating item"
	curl -X PUT -H "Content-Type: application/json" -d '{"name":"updated item", "elo":1764}' http://localhost:8080/items/1
	echo "# Fetching updated item"
	curl http://localhost:8080/items/1
	echo "# Deleting item"
	curl -X DELETE http://localhost:8080/items/1
	echo "# Fetching all items after deletion"
	curl http://localhost:8080/items

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X
