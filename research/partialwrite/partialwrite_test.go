package partialwrite

import "testing"

func TestH1_SingleOverwriteTouchesOneChunk(t *testing.T) {
	fileSize := int64(FileSizeFixture)
	offset := int64(1024 * 1024)
	length := int64(100)
	touched := ChunksTouchedByOverwrite(offset, length, ChunkSizeMean)
	if touched != 1 {
		t.Fatalf("H1 fallo: esperaba 1, got %d", touched)
	}
	baseline := BaselineWriteCost(fileSize)
	policy := PolicyWriteCost(offset, length, ChunkSizeMean, fileSize)
	if policy >= baseline {
		t.Fatalf("H1 fallo: policy %d debe ser < baseline %d", policy, baseline)
	}
}

func TestH2_MultipleOverwrites(t *testing.T) {
	offsets := []int64{0, 2 * ChunkSizeMean, 5 * ChunkSizeMean, 10 * ChunkSizeMean, 20 * ChunkSizeMean}
	total := 0
	for _, off := range offsets {
		total += ChunksTouchedByOverwrite(off, 100, ChunkSizeMean)
	}
	if total != 5 {
		t.Fatalf("H2 fallo: esperaba 5, got %d", total)
	}
}

func TestH3_CostLessThan10Percent(t *testing.T) {
	fileSize := int64(FileSizeFixture)
	offset := int64(1024 * 1024)
	policy := PolicyWriteCost(offset, 100, ChunkSizeMean, fileSize)
	baseline := BaselineWriteCost(fileSize)
	r := Ratio(policy, baseline)
	if r >= 0.10 {
		t.Fatalf("H3 fallo: ratio %.4f debe ser < 0.10", r)
	}
	t.Logf("H3 OK: policy=%d baseline=%d ratio=%.4f", policy, baseline, r)
}

func TestEdge_CrossesBoundaryTouchesTwo(t *testing.T) {
	offset := int64(ChunkSizeMean - 50)
	touched := ChunksTouchedByOverwrite(offset, 100, ChunkSizeMean)
	if touched != 2 {
		t.Fatalf("borde fallo: esperaba 2, got %d", touched)
	}
}
