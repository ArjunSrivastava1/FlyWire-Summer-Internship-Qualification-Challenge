//this checks only for the three candidate dbs, and contains maol to show why it was dropped and how it works in verification
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkInduced(filename string, nodes []string) {
	f, _ := os.Open(filename)
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

	fmt.Printf("\n%s:\n", filename)
	for _, a := range nodes {
		for _, b := range nodes {
			if a != b {
				if edges[[2]string{a, b}] {
					fmt.Printf("  %s -> %s\n", a, b)
				}
			}
		}
	}
}

func main() {
	manc := []string{"11981", "13849", "25840"} //nodes are hardcoded to ensure explicitness and correctness, but can be passed via modules too
	maol := []string{"11981", "13849", "25840"}
	banc := []string{"720575941684940783", "720575941503813207", "720575941595063766"}

	checkInduced("manc_1.2.1_edge_list.csv", manc)
	checkInduced("maol_1.1_edge_list.csv", maol)
	checkInduced("banc_626_edge_list.csv", banc)
}
