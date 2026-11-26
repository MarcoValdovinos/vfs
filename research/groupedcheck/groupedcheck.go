package groupedcheck

const (
	ChunkSize      = 4 * 1024 * 1024
	GroupSize      = 8
	CheckpointSize = 1 * 1024 * 1024 * 1024
	OverheadMs     = 10
)

func NumChunks(checkpointSize int64, chunkSize int64) int {
	n := int(checkpointSize / chunkSize)
	if checkpointSize%chunkSize != 0 {
		n++
	}
	return n
}

func RequestsWithoutGrouping(numChunks int) int { return numChunks }

func RequestsWithGrouping(numChunks int, groupSize int) int {
	if groupSize <= 0 {
		return numChunks
	}
	n := numChunks / groupSize
	if numChunks%groupSize != 0 {
		n++
	}
	return n
}

func TotalTimeMs(numRequests int, overheadMs int) int { return numRequests * overheadMs }
