# HIP-1344 Bounty 1 - Partial-Write Absorption

## Hipotesis pre-declarada
- H1: overwrite 100 bytes en 24 MiB toca 1 chunk
- H2: N overwrites tocan <= N chunks
- H3: costo con absorcion < 10 por ciento del costo sin absorcion

## Baselines
- Baseline: sin absorcion, reescribe todo = 24 MiB
- Policy: con absorcion, solo chunks tocados = 512 KiB por chunk

## Como correr
go test -v ./research/partialwrite/...
go vet ./research/partialwrite/...

## Honesty Labels
- No real hardware, no object store real
- Cost model hand-set: chunk 512 KiB, file 24 MiB
- Figures are simulation outputs, not Hanzo expectations
- Clean-room: no upstream code copied
