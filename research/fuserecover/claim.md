Claiming Bounty #8 - Transparent FUSE daemon recovery - Challenger deliverable

Intent: HIP-1344 #8, en hanzoai/vfs

Falsifiable Hypothesis:
- H1: Tras crash + restart, open handles se recuperan igual
- H2: Fsyncs confirmados no se pierden, lost = 0
- H3: Sin crash, restart no crea nuevo daemon

Matched Baselines:
- Baseline: sin recovery, crash pierde handles y fsyncs
- Policy: con recovery, session, handles y fsyncs preservados

Deliverable:
- research/fuserecover/fuserecover.go - clean-room, stdlib only
- research/fuserecover/fuserecover_test.go - 4 tests H1/H2/H3/borde
- research/fuserecover/README.md

Evidence Termux Android/arm64:
go test -v ./research/fuserecover/... PASS 4/4
H2 OK: ack=10 lost=0
go vet clean

Honesty Labels:
- No FUSE real, no mount real
- Modelo de sesion simplificado, no guest sandbox
- Figures are simulation outputs, not Hanzo expectations
- Clean-room, no upstream code copied
