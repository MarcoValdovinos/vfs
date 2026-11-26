# HIP-1344 Bounty 3 - Grouped checkpoint IO

## Hipotesis
- H1: 1 GiB en chunks 4 MiB = 256 chunks, agrupados de 8 = 32 requests
- H2: Requests por GiB con grouping < sin grouping, reduccion 8x
- H3: Tiempo con grouping < 20 por ciento del tiempo sin grouping

## Baselines
- Baseline: 1 request por chunk = 256 requests por GiB
- Policy: 1 request por grupo de 8 = 32 requests por GiB

## Como correr
go test -v ./research/groupedcheck/...

## Honesty Labels
- No S3 real, no NVMe real
- Overhead model hand-set: 10ms por request
- Figures are simulation outputs, not Hanzo expectations
- Clean-room, no upstream code copied
