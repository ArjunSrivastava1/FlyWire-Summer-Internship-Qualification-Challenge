package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	files := []string{ //csv database names as obtained, be sure to make changes here to match actual local filenames
		"banc_626_edge_list.csv",
		"fafb_783_edge_list.csv",
		"manc_1.2.1_edge_list.csv",
		"maol_1.1_edge_list.csv",
		"mcns_0.9_edge_list.csv",
	}

	for _, fname := range files {
		start := time.Now()

		file, err := os.Open(fname)
		if err != nil {
			fmt.Printf("%s: ERROR opening: %v\n", fname, err)
			continue
		}

		nodes := make(map[string]bool)
		edges := 0

		scanner := bufio.NewScanner(file)
		// Increase buffer for long lines
		scanner.Buffer(make([]byte, 1024*1024), 10*1024*1024)

		for scanner.Scan() {
			line := scanner.Text()
			// Skip header if exists
			if strings.HasPrefix(line, "source") || strings.HasPrefix(line, "pre") {
				continue
			}
			parts := strings.Split(line, ",")
			if len(parts) < 2 {
				continue
			}
			src := strings.TrimSpace(parts[0])
			dst := strings.TrimSpace(parts[1])
			nodes[src] = true
			nodes[dst] = true
			edges++
		}

		file.Close()

		elapsed := time.Since(start)
		fmt.Printf("%s: %d edges, %d nodes (%.2fs)\n",
			fname, edges, len(nodes), elapsed.Seconds())
	}
}
