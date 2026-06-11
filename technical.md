This md presents a deeper dive for results produced and obtained during the research period


## Master Table: Datasets Overview

| Dataset | Full Name                   | Sex    | Region Covered           | Nodes   | Edges | File Size |
| ------- | --------------------------- | ------ | ------------------------ | ------- | ----- | --------- |
| MANC    | Male Adult Nerve Cord       | Male   | Nerve cord only          | 23,641  | 5.3M  | ~100MB    |
| BANC    | Brain and Nerve Cord        | Female | Brain + nerve cord       | 112,885 | 2.7M  | ~100MB    |
| MCNS    | Male Central Nervous System | Male   | Brain + nerve cord       | 165,820 | 6.2M  | ~100MB    |
| MAOL    | Male Adult Optic Lobe       | Male   | Optic Lobe/Visual System | 51,669  | 6.5M  | ~100MB    |
| FAFB    | Female Adult Fly Brain      | Female | Brain                    | 138,584 | 3.7M  | ~100MB    |

---

## Node Overlap Analysis

| Pair | Overlap Size | Implication |
|------|--------------|-------------|
| MANC ∩ MAOL | 5,293 | Significant overlap: candidate for same-ID matching |
| MANC ∩ BANC | 2 | Negligible: requires isomorphism (structural matching) |
| MANC ∩ MCNS | 0 | No ID overlap: requires isomorphism |

## Comparison Table: Circuit Across Three Datasets

| Feature | MANC | BANC | MCNS |
|---------|------|------|------|
| **Neuron A ID** | 11981 | 720575941612175635 | 45246 |
| **Neuron A Name** | LEGNP_T3.LEGNP_T1.19 | GNG.317 | g23008.2 |
| **Neuron A NT** | Acetylcholine (0.97) | Acetylcholine (0.75) | Acetylcholine (predicted) |
| **Neuron A Cell Type** | DNp43, DNxl051 | CB0467 | SLP372 |
| **Neuron A Flow** | Descending (afferent) | Central brain intrinsic | cb_intrinsic |
| **Neuron B ID** | 13849 | 720575941473151035 | 56994 |
| **Neuron B Name** | OV.158 | GNG.122 | g18423.2 |
| **Neuron B NT** | Glutamate (0.90) | Acetylcholine (0.75) | Acetylcholine (predicted) |
| **Neuron B Cell Type** | AN09B021 | CB0553 | SLP359 |
| **Neuron B Flow** | Ascending (efferent) | Central brain intrinsic | cb_intrinsic |
| **Neuron C ID** | 25840 | 720575941589440140 | 128125 |
| **Neuron C Name** | LEGNP_T1.585 | GNG.1124 | g65572.4 |
| **Neuron C NT** | Acetylcholine (0.95) | GABA (0.87) | Unknown (all 0.0) |
| **Neuron C Cell Type** | AN17A050 | CB0460 | SMP167 |
| **Neuron C Flow** | Ascending (efferent) | Central brain intrinsic | cb_intrinsic |
| **One-to-Many Divergence?** | No | **Yes (3+ targets)** | No |
| **Circuit Class** | Sensorimotor | Local intrinsic | Central intrinsic |

---

## Key Biological Finding: BANC One-to-Many Divergence

The `A → B` core in BANC diverges to **at least three distinct C neurons** satisfying the induced subgraph condition:

| C Neuron | ID | Neurotransmitter | Target | Confidence |
|----------|-----|------------------|--------|------------|
| GNG.1124 | 720575941589440140 | GABA | Descending neuron DNfl032 | High (0.87) |
| GNG.632 | 720575941454380013 | Serotonin | Proboscis muscles | Low (0.29) |
| GNG.363 | 720575941611153069 | GABA | Cervical nerve (CvN) | High (0.92) |

**Interpretation:** The A→B core acts as a **sensory-motor command hub** in the female feeding system, broadcasting taste information to multiple output channels (inhibition of competing programs, direct muscle modulation, neck coordination). This finding is unique to BANC and is detailed further in `science.md`.

---

## Weak Connectivity (Per June 8 Clarification)

The circuit `A → B → C` forms a single connected component when edge directions are ignored:

```
A — B — C  (undirected path of length 2)
```

This satisfies the weak connectivity requirement. Verified by checking that the undirected graph on `{A,B,C}` has exactly 2 edges and is fully connected.

---

## Performance Metrics

| Operation | Time |
|-----------|------|
| Parse all 5 edge lists (Go) | 11 seconds |
| MANC node set construction | 1.7s |
| MANC 3-node chain enumeration | <2 seconds |
| BANC isomorphism search | 2.3 seconds |
| MCNS isomorphism search | 3.5 seconds |
| **Total verification time** | **~11 seconds** |

*Python equivalent would take ~9 minutes just to load all the given 5 datasets.*
Python, to load 1 dataset alone(BANC) produced took the following times:

'took 1m42s' 

As compared to go:

banc_626_edge_list.csv: 2676592 edges, 112885 nodes (1.31s)
fafb_783_edge_list.csv: 3732460 edges, 138584 nodes (1.97s)
manc_1.2.1_edge_list.csv: 5305638 edges, 23641 nodes (1.69s)
maol_1.1_edge_list.csv: 6484936 edges, 51669 nodes (2.27s)
mcns_0.9_edge_list.csv: 6239112 edges, 165820 nodes (3.54s)

'took 18s'
A vast speed up that allowed for iteration, prototyping within the provided time constraint

### About Existing Libraries?

| Library | Why Not Used |
|---------|--------------|
| NetworkX (Python) | Too slow on 100M+ edges |
| igraph (R/Python) | Overhead |
| nauty/Traces | Requires compilation|
| VF2 | General algorithm but slower than our degree-filtered brute force |

---

## Results: Final Matched Circuit (`network.csv`)

| MANC | BANC | MCNS |
|------|------|------|
| 11981 | 720575941612175635 | 45246 |
| 13849 | 720575941473151035 | 56994 |
| 25840 | 720575941589440140 | 128125 |

**Induced subgraph:** `A → B → C` in all three datasets. 
