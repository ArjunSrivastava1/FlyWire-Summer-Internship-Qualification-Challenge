# FlyWire Qualification Challenge – Technical Approach

## Overview

I identified the largest common induced subgraph (N=3) shared across **MANC, BANC, and MCNS** datasets: a feed-forward chain **A → B → C**.

The circuit is weakly connected, satisfies the induced subgraph condition (no extra edges among the three neurons), and appears identically in all three datasets. Biological analysis focused on BANC (gnathal ganglia, feeding motor control).

---

## Technical Strategy & Heuristics

The technical strategy for such a task, of analysis of a csv dataset which has repetitive data in it is based on a very simple factor, how much of it can be automated? and what is needed?
For research work, there are many AI tools that exist now, and many that I have worked on myself( as a oss contributor to huggingface) such as google notebook that allows one to perform dataset Q&A with gemini, or ml-intern: which is ml software that in and of itself is capable of reading papers and shipping models

Building up on my experience of scala, providing distributed systems programming and functional programmatic thinking paradigms: manipulating files and then filtering via dataframes is the easiest way to look for combined matching ids, or chain motifs is the easiest portion

However, given that the csv files only consisted of the Neuron IDs and nothing more, relying on the AI framework alone would have not yielded qualitative results, yes, there are many capabilities to them, yet human verification is an important component of all workflows, and for reading through the neurons and what they did? or interpretation of results? that was something an ai could not have been relied upon(such as the figures and the meshes, vision models cant parse through those yet and get confused), the standard benchmarks for this are well known and such structures only add overhead expenses that are best avoided, along with unnecessary data, so usage of AI is best limited to only parsing the found results

So the technical strategy came down to this:find the connections(graphs) and then filter those for the needed number of N, followed further by looking in the other datasets until the same structure was found across 3 different datasets

And reference with the human(researcher/programmer) as to the interpretation and what each structure did biologically along with lookups and further metadata analysis on the codex provided, which contained vital information as the neuron ids are just numbers

### Search Heuristic (Step-by-Step)

1. Anchor on smallest dataset — MANC (23,641 nodes) to minimize search space
2. Enumerate directed 3-node chains `A → B → C`
3. Filter by degree sequence prune candidates where in/out degrees don't match across datasets
4. Verify induced subgraph condition in MANC:
   - `A → B` exists 
   - `B → C` exists 
   - No edges: `B → A`, `C → B`, `A → C`, `C → A` 
5. Search BANC and MCNS for identical adjacency matrix
6. Export matched structure to `network.csv`

### Why This Works

- Real connectomes are sparse (density ~0.0002)
- 3-node chains are the most common motif in neural circuits
- Induced subgraph verification ensures true isomorphism

## Computational Heuristics & Algorithms

### Graph Representation

- **Adjacency map:** `map[[2]string]bool` O(1) edge existence checks
- **Outgoing neighbors index:** `map[string][]string` O(1) lookup of all targets from a source

### Algorithm 1: 3-Node Chain Enumeration (Brute Force with Early Exit)

```go
for a := range nodes {
    for _, b := range outNeighbors[a] {
        for _, c := range outNeighbors[b] {
            // Check induced subgraph condition
            if !hasExtraEdges(a,b,c) {
                // Candidate found
            }
        }
    }
}
```

Complexity: O(N × d²) where d = average degree (~224 in MANC).  
Actual runtime: <2 seconds on 23,641 nodes.

### Algorithm 2: Pattern Search: 
For a given hardcoded pattern (a,b,c), the script exhaustively searches BANC for any triple (x,y,z) with x→y and y→z and no extra edges

For a triple `(a,b,c)`, verify:

```
Required edges:
  a→b ✓
  b→c ✓

Forbidden edges:
  b→a ✗
  c→b ✗
  a→c ✗
  c→a ✗
```

### Summary of Heuristics

| Heuristic | Purpose | Impact |
|-----------|---------|--------|
| Anchor on smallest dataset | Minimize search space | 23k vs 165k nodes |
| One-to-many divergence detection | Identify command hubs | Unique BANC finding |
| Weak connectivity check | Satisfy new req | Single component verification |
| Induced subgraph strictness | No extra edges | Enforces isomorphism |

---
### Alternative Approaches Considered (and Why Rejected)

| Approach | Why Rejected |
|----------|---------------|
| Same-ID matching | Same IDs that were present in one group, broke isomorphism in others |
| Exhaustive n-node searching | Required expensive operations for pattern matching on the datasets |
| Graph neural networks | would require embedding training |
| Other inclusions | Extra edges violated induced subgraph condition |

The alternative approaches were rejected due to either violating the rules of induced subgraphs, isomorphism, or being computationally expensive, while a higher n could have been aimed for, it would have required analysis of further complex structures that looped back on themselves(manc has 3 n=4 chains in which node 2 and node 4 are same)
---

## Assumptions Master Table

| # | Assumption | Rationale | Impact if Wrong |
|---|------------|-----------|------------------|
| 1 | Edge weights (synapse counts) ignored | Per problem statement | No impact on topology |
| 2 | Different IDs across datasets ≠ same neuron | IDs are reconstruction-specific | Isomorphism required |
| 3 | Weak connectivity required | Per June 8 clarification | circuit satisfies (A-B-C chain) |
| 4 | Self-loops excluded | Not present in edge lists | No impact |
| 5 | Neurotransmitter predictions from Codex used as given | No experimental validation available | May contain false positives |
| 6 | Induced subgraph = no extra edges among matched neurons | Standard graph theory definition | Strictly enforced |
| 7 | Directed edges only (unweighted) | Per problem statement | Ignored synapse counts |

---

## Technical Steps to Reproduce:
The attached code is a folder that contains 3 modules of go code, 1 is to obtain stats over the datasets provided, the other 2 are for finding structures that are connected in a dataset, and the last is to verify the obtained structures for verifications(extra edges, loops etc etc)

As the approach itself required human lookups for metadata, the code is kept to a minimum, further, the files are easily modified to check for structures in other datasets, as such, finding and verification files are added for only 1 dataset, and produce one line results that are easy to understand, parse and are obtained within ~5 secs.

As for steps, clone the repo:
```bash
git clone https://github.com/ArjunSrivastava1/FlyWire-Summer-Internship-Qualification-Challenge
cd FlyWire-Summer-Internship-Qualification-Challenge
```

Run the provided codes within, you can even extend the given codes to change the csv files, or add further nodes easily
```
go run src/stats.go 
go run src/find_banc.go   # verifies the matched triple
```
Have fun!!
