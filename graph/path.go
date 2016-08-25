package graph

import "math"

// Function for finding the shortest path between two vertexes in the graph. It
// uses algorithm by Edsger W. Dijkstra. Returns array of vertexes which
// represent the shortest path and distance.
func Path(g [][]float64, s, d int) ([]int, float64) {

	l := len(g)

	path := make(map[int]int)
	dist := make(map[int]float64)
	q := make([]bool, l)

	dist[s] = 0.0

	for l > 0 {

		cur := -1
		min := math.NaN()

		for i := 0; i < len(q); i++ {

			if dst, ok := dist[i]; !q[i] && ok && (math.IsNaN(min) || dst < min) {
				cur = i
				min = dst
			}

		}

		if cur == -1 {
			// path does not exist
			return []int{}, 0.0
		}

		if cur == d {
			// destination reached
			break
		}

		// mark vertex as processed
		q[cur] = true
		l--

		for i := 0; i < len(g[cur]); i++ {

			if g[cur][i] <= 0.0 {
				continue // no edge between vertexes
			}

			a := dist[cur] + g[cur][i]

			if dst, ok := dist[i]; !ok || a < dst {
				path[i] = cur
				dist[i] = a
			}

		}

	}

	// build result path

	r := []int{}
	i := d

	for v, ok := path[i]; ok; v, ok = path[i] {
		r = append([]int{i}, r...)
		i = v
	}

	// add source to the beginning of the path
	r = append([]int{s}, r...)

	return r, dist[d]

}
