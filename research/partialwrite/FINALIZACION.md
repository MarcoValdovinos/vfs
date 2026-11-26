## Paper & License Audit
HIP: HIP-1344 Bounty #1 Partial-write absorption
Paper: No paper externo, modelo basado en spec de HIP-1344
Licencia: Clean-room, stdlib only, no upstream code copied

## Measurements & HIP Update
go test -v ./research/partialwrite/... PASS
Propuesta HIP: partial-write debe absorberse sin perder datos confirmados ni corromper el resto

## Evidence & Safety
Evidence: Termux Android arm64, go test PASS, go vet clean
Safety: Research-only, no VFS real, no cambio a produccion
CI: hanzo/ci falla por fork PR, no por codigo. GitGuardian PASS
