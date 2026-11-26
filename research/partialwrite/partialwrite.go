package partialwrite

const (
	ChunkSizeMean   = 512 * 1024
	FileSizeFixture = 24 * 1024 * 1024
)

func ChunksForFile(fileSize int64, chunkSize int64) int {
	if fileSize <= 0 {
		return 0
	}
	n := int(fileSize / chunkSize)
	if fileSize%chunkSize != 0 {
		n++
	}
	return n
}

func ChunksTouchedByOverwrite(offset int64, length int64, chunkSize int64) int {
	if length <= 0 {
		return 0
	}
	start := offset / chunkSize
	end := (offset + length - 1) / chunkSize
	return int(end-start) + 1
}

func BaselineWriteCost(fileSize int64) int64 { return fileSize }

func PolicyWriteCost(offset int64, length int64, chunkSize int64, fileSize int64) int64 {
	touched := ChunksTouchedByOverwrite(offset, length, chunkSize)
	cost := int64(touched) * chunkSize
	if cost > fileSize {
		cost = fileSize
	}
	return cost
}

func Ratio(policyCost int64, baselineCost int64) float64 {
	if baselineCost == 0 {
		return 0
	}
	return float64(policyCost) / float64(baselineCost)
}
