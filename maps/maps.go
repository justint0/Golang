/*
Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages).

To create an empty map, use the builtin make: make(map[key-type]val-type)
*/

package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	// Set key/value pairs using name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13 

	fmt.Println("map:", m)

	// Get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn't exist, the zero value of the value type is returned
	v3 :=  m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	/*
	The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}