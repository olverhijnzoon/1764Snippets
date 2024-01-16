package main

import (
	"fmt"
	"math/rand"

	mystore "github.com/olverhijnzoon/1764Snippets/golang/interfaces/mystore"
	storecle "github.com/olverhijnzoon/1764Snippets/golang/interfaces/storecle"
)

/*
	"Declare an interface where it is used, not where it is implemented. Unless the interface is well discovered." https://twitter.com/inancgumus/status/1605116543553019905?s=20&t=fTo98MYIjsZ5nw44XnTmRw

	In Go, it is generally considered good practice to declare an interface in the same package where it is used, rather than where it is implemented. This is because interfaces are a way of specifying the behavior of a type, rather than the implementation of that type. By declaring an interface in the same package where it is used, you make it clear to users of the interface what they can expect from types that implement it.

	For example, suppose we have a package called mystore that defines an interface called Store for storing and retrieving data. If we declare the Store interface in the mystore package, users of the mystore package can see exactly what methods they can call on types that implement the Store interface.

	On the other hand, if we declare the Store interface in a separate package where it is implemented, users of the mystore package would not be able to see the methods of the Store interface unless they imported the implementation package. This could make it more difficult for users of the mystore package to understand how to use types that implement the Store interface.

	There are, however, cases where it may be appropriate to declare an interface in a separate package where it is implemented. For example, if the interface is intended to be a "well-known" interface that will be used by many packages, it may make sense to declare it in a separate package to make it easier for other packages to import and use.

	In the following example, the Store interface is declared here in the main package and specifies two methods: Save and Load. Any type that implements these methods can be used as a Store.
*/

// Store is an interface for storing and retrieving data.
type Store interface {
	// Save saves the given data to the store.
	Save(data []byte) error

	// Load loads the data from the store.
	Load() ([]byte, error)
}

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Interfaces")

	/*
		The provided code demonstrates the use of an interface in Golang named Store, which declares two methods: Save(data []byte) error for storing data, and Load() ([]byte, error) for retrieving data. In the main function, it creates a variable s of type Store and, based on a random number generator, it assigns s a value that implements the Store interface from either the mystore or storecle packages. Afterwards, the code calls a function named processData, which expects a type Store as input. This function first calls the Load() method to get the data and after some processing, it calls the Save() method to save back the processed data to the store.

		This is a slightly modified version of "Golang Interfaces" of the 42Snippets repository.
	*/

	// Create a variable of type Store
	var store Store

	// Set store to a value that implements the Store interface like MyStore from mystore package
	if rand.Intn(2) == 0 { // Generate a random number that is either 0 or 1
		store = &mystore.MyStore{}
	} else {
		store = &storecle.Storecle{}
	}

	// Call a function that expects a Store
	err := processData(store)
	if err != nil {
		// Handle the error...
	}
}

func processData(store Store) error {
	data, err := store.Load()
	if err != nil {
		return err
	}

	// Process the data...
	fmt.Println("Processing the data!")

	return store.Save(data)
}
