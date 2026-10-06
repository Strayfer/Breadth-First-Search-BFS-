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
	// --- Normal graph with a shortcut ---
	g := map[string][]string{
		"Andeevka":       {"Toretsk", "Kosntantinovka"},
		"Toretsk":        {"Andeevka", "Pokrovsk"},
		"Kosntantinovka": {"Andeevka", "Mirnograd"},
		"Pokrovsk":       {"Toretsk", "Mirnograd"},
		"Mirnograd":      {"Pokrovsk", "Kosntantinovka"},
	}
	o, d := BFS(g, "Andeevka")
	fmt.Println("Order:", o, "Dist:", d)

	// FIXED: Pokrovsk before Mirnograd
	check("normal order", eq(o, []string{"Andeevka", "Toretsk", "Kosntantinovka", "Pokrovsk", "Mirnograd"}))
	check("normal distances", d["Andeevka"] == 0 && d["Toretsk"] == 1 && d["Kosntantinovka"] == 1 && d["Pokrovsk"] == 2 && d["Mirnograd"] == 2)

	// --- Changed start ---
	o2, _ := BFS(g, "Pokrovsk")
	check("changed start", eq(o2, []string{"Pokrovsk", "Toretsk", "Mirnograd", "Andeevka", "Kosntantinovka"}))

	// --- Single node ---
	o3, d3 := BFS(map[string][]string{"Andeevka": {}}, "Andeevka")
	check("single node", eq(o3, []string{"Andeevka"}) && d3["Andeevka"] == 0)

	// --- Disconnected ---
	g4 := map[string][]string{"Andeevka": {"Toretsk"}, "Toretsk": {"Andeevka"}, "X": {"Y"}, "Y": {"X"}}
	o4, d4 := BFS(g4, "Andeevka")
	check("disconnected", len(o4) == 2)
	_, reach := d4["X"]
	check("disconnected X unreachable", !reach)

	// --- Cycle (would loop forever if visited-at-pop bug) ---
	g5 := map[string][]string{"Andeevka": {"Toretsk", "Kosntantinovka"}, "Toretsk": {"Andeevka", "Kosntantinovka"}, "Kosntantinovka": {"Andeevka", "Toretsk"}}
	o5, _ := BFS(g5, "Andeevka")
	check("cycle no loop", len(o5) == 3)

	// --- Empty graph ---
	o6, d6 := BFS(map[string][]string{}, "Z")
	check("empty graph", len(o6) == 1 && o6[0] == "Z" && d6["Z"] == 0)

	fmt.Printf("\n%d passed, %d failed\n", pass, fail)
}
