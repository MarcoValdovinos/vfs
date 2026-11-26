package groupedcheck

import "testing"

func TestH1_GroupingReducesRequests(t *testing.T) {
	chunks := NumChunks(CheckpointSize, ChunkSize)
	if chunks != 256 {
		t.Fatalf("H1 fallo: esperaba 256 chunks, got %d", chunks)
	}
	grouped := RequestsWithGrouping(chunks, GroupSize)
	if grouped != 32 {
		t.Fatalf("H1 fallo: esperaba 32 requests, got %d", grouped)
	}
}

func TestH2_RequestsPerGiB(t *testing.T) {
	chunks := NumChunks(CheckpointSize, ChunkSize)
	base := RequestsWithoutGrouping(chunks)
	grouped := RequestsWithGrouping(chunks, GroupSize)
	if grouped >= base {
		t.Fatalf("H2 fallo: grouped %d debe ser < base %d", grouped, base)
	}
	t.Logf("H2 OK: base=%d grouped=%d reduccion=8x", base, grouped)
}

func TestH3_TimeLessThan20Percent(t *testing.T) {
	chunks := NumChunks(CheckpointSize, ChunkSize)
	baseTime := TotalTimeMs(RequestsWithoutGrouping(chunks), OverheadMs)
	groupedTime := TotalTimeMs(RequestsWithGrouping(chunks, GroupSize), OverheadMs)
	ratio := float64(groupedTime) / float64(baseTime)
	if ratio >= 0.20 {
		t.Fatalf("H3 fallo: ratio %.4f debe ser < 0.20", ratio)
	}
	t.Logf("H3 OK: baseTime=%d groupedTime=%d ratio=%.4f", baseTime, groupedTime, ratio)
}

func TestEdge_PartialGroup(t *testing.T) {
	grouped := RequestsWithGrouping(10, 8)
	if grouped != 2 {
		t.Fatalf("borde fallo: esperaba 2, got %d", grouped)
	}
}
