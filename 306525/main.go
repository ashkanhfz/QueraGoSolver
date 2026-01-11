package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024), 1024*1024)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return
	}

	i := 0
	n, _ := strconv.Atoi(strings.TrimSpace(lines[i]))
	i++
	scores := make(map[string]float64)

	for k := 0; k < n; k++ {
		if i >= len(lines) {
			break
		}
		parts := strings.Fields(lines[i])
		i++
		if len(parts) < 2 {
			continue
		}
		S, _ := strconv.Atoi(parts[0])
		P, _ := strconv.Atoi(parts[1])

		var line string
		if i < len(lines) {
			line = strings.TrimSpace(lines[i])
			i++;
		}

		var names []string
		if line != "" {
			for _, nm := range strings.Split(line, ",") {
				nm = strings.TrimSpace(nm)
				if nm != "" {
					names = append(names, nm)
				}
			}
		}

		if P <= 0 {
			continue
		}
		if P == 1 {
			if len(names) > 0 {
				scores[names[0]] += float64(S)
			}
			continue
		}
		den := float64(P - 1)
		for idx, name := range names {
			r := idx + 1
			scores[name] += float64(S) * float64(P-r) / den
		}
	}

	type pair struct {
		name  string
		score float64
	}
	pairs := make([]pair, 0, len(scores))
	for name, sc := range scores {
		pairs = append(pairs, pair{name, sc})
	}
	sort.Slice(pairs, func(a, b int) bool {
		if pairs[a].score == pairs[b].score {
			return pairs[a].name < pairs[b].name
		}
		return pairs[a].score > pairs[b].score
	})

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, p := range pairs {
		fmt.Fprintln(w, p.name)
	}
}
