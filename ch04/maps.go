package ch04

import (
	"fmt"
	"sort"
)

// Maps prints the contents of a map
func Maps() {

	// a map
	ages := make(map[string]int)

	ages["alice"] = 31
	ages["charlie"] = 34

	// a map literal
	a := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	// accessing a map element
	fmt.Println(a["charlie"])

	// deleting a map element
	delete(a, "alice")

	/*
	 * All of these operations are safe even if the element isn’t in the map;
	 * a map lookup using a key that isn’t present returns the zero value for
	 * its type, so, for instance, the following works even when "bob" is not
	 * yet a key in the map because the value of ages["bob"] will be 0.
	 */

	/*
	 * To enumerate all the key/ value pairs in the map, we use a range-based for
	 * loop similar to those we saw for slices. Successive iterations of the loop
	 * cause the name and age variables to be set to the next key/ value pair:
	 */

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	/*
	 * The order of map iteration is unspecified, and different implementations might use
	 * a different hash function, leading to a different ordering. In practice, the order
	 * is random, varying from one execution to the next. This is intentional; making the
	 * sequence vary helps force programs to be robust across implementations. To
	 * enumerate the key/value pairs in order, we must sort the keys explicitly, for
	 * instance, using the Strings function from the sort package if the keys are strings.
	 */

	var names []string

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	/*
	 * Subscripting a map in this context yields two values; the second is a boolean that reports
	 * whether the element was present. The boolean variable is often called ok, especially if it
	 * is immediately used in an if condition.
	 */

	if age, ok := ages["bob"]; !ok {
		/*...*/
		fmt.Println(age)
	}
}

/*
 * The value type of a map can itself be a composite type, such as a map or slice. In the
 * following code, the key type of graph is string and the value type is map[string]bool,
 * representing a set of strings. Conceptually, graph maps a string to a set of related
 * strings, its successors in a directed graph.
 */

var graph = make(map[string]map[string]bool)

func addEgde(from, to string) {
	edges := graph[from]

	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}