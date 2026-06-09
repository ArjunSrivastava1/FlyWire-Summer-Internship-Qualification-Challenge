# A Three-Neuron Feedforward Circuit in the Drosophila Gnathal Ganglia

**Dataset:** BANC v626 (female adult brain + nerve cord)  
**Region:** Gnathal ganglia (feeding motor control)

---

## Figures

<img width="640" height="480" alt="Figure_1" src="https://github.com/user-attachments/assets/80a9a3f8-5fc9-4c40-9117-dadfd175bd4f" />

*Figure 1: Induced subgraph of the identified circuit. A (GNG.317) and B (GNG.122) are cholinergic (ACh). C (GNG.1124) is GABAergic. Arrows indicate directed synaptic connectivity. No other edges exist among these three neurons.*

<img width="960" height="540" alt="Untitled presentation" src="https://github.com/user-attachments/assets/b15b1fe7-03f0-4626-b30e-f29156f37357" />

*Figure 2: 3D reconstructions of (A) GNG.317, (B) GNG.122, and (C) GNG.1124 in BANC v626. Neurons are shown in the default Codex viewer with surrounding neuropil for anatomical context. Somata are located in the central brain.*

---

## Observations

We identified a three-neuron feedforward chain (A→B→C) as an identical induced subgraph in BANC, MANC, and MCNS (see `network.csv`). In BANC, the circuit localizes to the gnathal ganglia, a brain region controlling feeding movements.

The induced subgraph contains exactly two directed edges (A→B, B→C) with no other connections among the three neurons. Neurotransmitter predictions from Codex show:

- **A (GNG.317):** Acetylcholine (ACh, 0.75)
- **B (GNG.122):** Acetylcholine (ACh, 0.75)
- **C (GNG.1124):** GABA (0.87)

A and B receive sensory input from labellum taste sensilla. C outputs to descending neuron DNfl032, which projects to motor circuits in the ventral nerve cord.

---

## Interpretation & Hypothesis

The ACh→ACh→GABA pattern is a classic **feedforward inhibitory motif**, where initial excitation is followed by delayed inhibition. In the gnathal ganglia, this circuit may transform gustatory input into precisely timed motor output.

We hypothesize that C (GNG.1124) inhibits descending neuron DNfl032, shaping the timing of proboscis extension during feeding. This mechanism could:
- Prevent runaway excitation
- Create a discrete temporal window for motor commands
- Enable rhythmic, coordinated feeding movements

This architecture is consistent with known feedforward inhibition circuits that enforce temporal fidelity in motor systems (Pouille & Scanziani, 2001).

---

## References

1. Pouille, F., & Scanziani, M. (2001). Enforcement of temporal fidelity in pyramidal cells by somatic feed-forward inhibition. *Science*, 293(5532), 1159-1163.

2. Flood, T.F., et al. (2013). A single pair of interneurons commands the Drosophila feeding motor program. *Nature*, 499(7456), 83-87.

3. Scheffer, L.K., et al. (2020). A connectome and analysis of the adult Drosophila central brain. *eLife*, 9, e57443.
