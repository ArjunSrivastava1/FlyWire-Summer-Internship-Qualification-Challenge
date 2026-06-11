package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadBANC() map[[2]string]bool {
	f, _ := os.Open("banc_626_edge_list.csv")
	defer f.Close()
	edges := make(map[[2]string]bool)
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 10*1024*1024)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) < 2 {
			continue
		}
		src, dst := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		edges[[2]string{src, dst}] = true
	}
	return edges
}

func main() {
	banc := loadBANC()

	// Build outgoing map
	out := make(map[string][]string)
	for e := range banc {
		out[e[0]] = append(out[e[0]], e[1])
	}

	fmt.Println("Searching BANC for a chain isomorphic to a->b->c...")
	fmt.Println()

	count := 0
	for src, targets := range out {
		for _, mid := range targets {
			for _, tgt := range out[mid] {
				// Check no extra edges among these three
				extra := false
				if banc[[2]string{mid, src}] {
					extra = true
				}
				if banc[[2]string{tgt, mid}] {
					extra = true
				}
				if banc[[2]string{src, tgt}] {
					extra = true
				}
				if banc[[2]string{tgt, src}] {
					extra = true
				}

				if !extra {
					fmt.Printf("Found match #%d:\n", count+1)
					fmt.Printf("  BANC src: %s\n", src)
					fmt.Printf("  BANC mid: %s\n", mid)
					fmt.Printf("  BANC tgt: %s\n", tgt)
					fmt.Println()
					count++
					if count >= 3 {
						fmt.Println("Use any of the above as your BANC triple.")
						return
					}
				}
			}
		}
	}

	if count == 0 {
		fmt.Println("No exact chain found in BANC for this pattern.")
	}
}
