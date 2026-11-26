# HIP-1344 Bounty 8 - Transparent FUSE daemon recovery

## Hipotesis
- H1: Tras crash + restart, open handles se recuperan igual
- H2: Fsyncs confirmados no se pierden, lost = 0
- H3: Sin crash, restart no crea nuevo daemon

## Baselines
- Baseline: sin recovery, crash pierde handles y fsyncs
- Policy: con recovery, session, handles y fsyncs preservados

## Como correr
go test -v ./research/fuserecover/...

## Honesty Labels
- No FUSE real, no mount real
- Modelo de sesion simplificado, no guest sandbox
- Figures are simulation outputs, not Hanzo expectations
- Clean-room, no upstream code copied
