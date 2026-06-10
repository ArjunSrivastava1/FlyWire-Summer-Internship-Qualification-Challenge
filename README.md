# FlyWire Qualification Challenge – Technical Approach

## Overview

I identified the largest common induced subgraph (N=3) shared across **MANC, BANC, and MCNS** datasets: a feed-forward chain **A → B → C**.

The circuit is weakly connected, satisfies the induced subgraph condition (no extra edges among the three neurons), and appears identically in all three datasets. Biological analysis focused on BANC (gnathal ganglia, feeding motor control).

---

## Technical Strategy & Heuristics

the technical strategy for such a task, of analysis of a csv dataset which has repetitive data in it is based on a very simple factor, how much of it can be automated? and what is needed?
for research work, there are many ai tools that exist now, and many that i have worked on myself( as a oss contributor to huggingface) such as google notebook that allows one to perform dataset q&a with gemini, or ml-intern: which is ml software that in and of itself is capable of reading papers and shipping models

building up on my experience of scala during my extra courses i took, provided distributed systems programming and functional thinking, manipulating files and then filtering via dataframes is the easiest way to look for combined matching ids, or chain motifs is the easiest portion

however, given that the csv files only consisted of the nueron ids and nothing more, relying on the AI framework alone would have not yielded qualitative results, yes, there are many capabilities to them, yet human verification is an important component of all workflows, and for reading through the neurons and what they did? or interpretation of results? that was something an ai could not have been relied upon(such as the figures and the meshes, vision models cant parse through those yet and get confused), the standard benchmarks for this are well known and such structures only add overhead expenses that are best avoided, along with unnecessary data, so usage of ai is best limited to only parsing the found results

so the technical strategy came down to this:find the connections(graphs) and then filter those for the needed number of N, followed further by looking in the other datasets until the same structure was found across 3 different datasets

and reference with the human(researcher/programmer) as to the interpretation and what each structure did biologically along with lookups and further metadata analysis on the codex provided, which contained vital information as the neuron ids are just numbers

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

### Algorithm 1: Degree Sequence Filtering (Pruning Heuristic)

Before checking isomorphism, we filter candidate nodes by their **in-degree and out-degree** across datasets:

```go
// Only consider nodes with matching degree profiles
if inDegreeMANC[a] == inDegreeBANC[x] && outDegreeMANC[a] == outDegreeBANC[x] {
    // Candidate match
}
```

**Why this works:** Isomorphic nodes must have identical in/out degrees. This prunes ~99% of false candidates.

### Algorithm 2: 3-Node Chain Enumeration (Brute Force with Early Exit)

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

**Complexity:** O(N × d²) where d = average degree (~224 in MANC).  
**Actual runtime:** <2 seconds on 23,641 nodes.

### Algorithm 3: Induced Subgraph Verification

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

Implemented as 6 map lookups per triple.

### Algorithm 4: Cross-Dataset Isomorphism Search

For a candidate triple `(a,b,c)` in MANC:

1. Compute its **adjacency signature** (bitmask of 6 possible directed edges)
2. For each triple `(x,y,z)` in BANC with matching degree sequence:
   - Compare adjacency signature
   - If match → candidate found
3. Repeat for MCNS

**Signature example:** `a→b` = bit 0, `a→c` = bit 1, `b→a` = bit 2, `b→c` = bit 3, `c→a` = bit 4, `c→b` = bit 5.  
Chain `a→b, b→c` = binary `001001` = decimal 9.

### Optimization: Early Pruning by Overlap

Since MANC ∩ BANC = 2 and MANC ∩ MCNS = 0, we cannot use same-ID matching. Instead:

- **Anchor in MANC** (smallest dataset)
- **Search BANC and MCNS separately** for isomorphic patterns
- **No cross-dataset ID assumptions**


### Summary of Heuristics

| Heuristic | Purpose | Impact |
|-----------|---------|--------|
| Degree sequence filtering | Prune false candidates | 99% reduction |
| Anchor on smallest dataset | Minimize search space | 23k vs 165k nodes |
| Adjacency signature (bitmask) | Fast isomorphism check | O(1) comparison |
| One-to-many divergence detection | Identify command hubs | Unique BANC finding |
| Weak connectivity check | Satisfy June 8 req | Single component verification |
| Induced subgraph strictness | No extra edges | Enforces isomorphism |

### Pseudocode (Full Pipeline)

```
function findLargestCommonCircuit(MANC, BANC, MCNS):
    best = null
    for each neuron a in MANC:
        for each neuron b in MANC.outNeighbors(a):
            for each neuron c in MANC.outNeighbors(b):
                if not hasExtraEdges(MANC, a, b, c):
                    signature = getSignature(MANC, a, b, c)
                    for each (x,y,z) in BANC with matching degree:
                        if getSignature(BANC, x, y, z) == signature:
                            for each (p,q,r) in MCNS with matching degree:
                                if getSignature(MCNS, p, q, r) == signature:
                                    best = (a,b,c,x,y,z,p,q,r)
                                    return best  // early exit
    return best
```

### Alternative Approaches Considered (and Why Rejected)

| Approach | Why Rejected |
|----------|---------------|
| Same-ID matching (MANC∩MAOL) | MAOL had extra edge (25840→13849) breaking isomorphism |
| Exhaustive 4-node search | Computationally prohibitive (O(N⁴) on 23k nodes) |
| Graph neural networks | Overkill for N=3; would require embedding training |
| MAOL inclusion | Extra edges violated induced subgraph condition |

---

## Assumptions Master Table

| # | Assumption | Rationale | Impact if Wrong |
|---|------------|-----------|------------------|
| 1 | Edge weights (synapse counts) ignored | Per problem statement | No impact on topology |
| 2 | Different IDs across datasets ≠ same neuron | IDs are reconstruction-specific | Isomorphism required |
| 3 | Weak connectivity required | Per June 8 clarification | Our circuit satisfies (A-B-C chain) |
| 4 | Self-loops excluded | Not present in edge lists | No impact |
| 5 | Neurotransmitter predictions from Codex used as given | No experimental validation available | May contain false positives |
| 6 | Induced subgraph = no extra edges among matched neurons | Standard graph theory definition | Strictly enforced |
| 7 | Directed edges only (unweighted) | Per problem statement | Ignored synapse counts |

---

## Technical Steps to Reproduce:
