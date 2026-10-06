package main

import "fmt"

func BFS(g map[string][]string, s string) ([]string, map[string]int) {
	order, dist, seen := []string{}, map[string]int{s: 0}, map[string]bool{s: true}
	q := []string{s}
	for i := 0; i < len(q); i++ {
		n := q[i]
		order = append(order, n)
		for _, m := range g[n] {
			if !seen[m] {
				seen[m], dist[m] = true, dist[n]+1
				q = append(q, m)
			}
		}
	}
	return order, dist
}

var pass, fail int

func check(name string, ok bool) {
	if ok {
		pass++
		fmt.Println("PASS:", name)
	} else {
		fail++
		fmt.Println("FAIL:", name)
	}
}

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// --- Normal graph ---
	g := map[string][]string{
		"A": {"B", "C"}, "B": {"A", "D", "E"},
		"C": {"A"}, "D": {"B"}, "E": {"B"},
	}
	o, d := BFS(g, "A")
	fmt.Println("Order:", o, "Dist:", d)
	check("normal order", eq(o, []string{"A", "B", "C", "D", "E"}))
	check("normal distances", d["A"] == 0 && d["B"] == 1 && d["C"] == 1 && d["D"] == 2 && d["E"] == 2)

	// --- Changed start ---
	o2, _ := BFS(g, "D")
	check("changed start", eq(o2, []string{"D", "B", "A", "E", "C"}))

	// --- Single node ---
	o3, d3 := BFS(map[string][]string{"A": {}}, "A")
	check("single node", eq(o3, []string{"A"}) && d3["A"] == 0)

	// --- Disconnected ---
	g4 := map[string][]string{"A": {"B"}, "B": {"A"}, "X": {"Y"}, "Y": {"X"}}
	o4, d4 := BFS(g4, "A")
	check("disconnected", len(o4) == 2)
	_, reach := d4["X"]
	check("disconnected X unreachable", !reach)

	// --- Cycle (would loop forever if visited-at-pop bug) ---
	g5 := map[string][]string{"A": {"B", "C"}, "B": {"A", "C"}, "C": {"A", "B"}}
	o5, _ := BFS(g5, "A")
	check("cycle no loop", len(o5) == 3)

	// --- Empty graph ---
	o6, d6 := BFS(map[string][]string{}, "Z")
	check("empty graph", len(o6) == 1 && o6[0] == "Z" && d6["Z"] == 0)

	fmt.Printf("\n%d passed, %d failed\n", pass, fail)

}
