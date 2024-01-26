package outofboundarypaths

type pathCountMemory [][]map[int]int

type MapInfo struct {
	rows    int
	columns int
	pathCountMemory
}

const (
	modNum = 1e9 + 7
)

var dxy = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

var mapInfo MapInfo

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	// init info
	mapInfo = MapInfo{
		rows:            m,
		columns:         n,
		pathCountMemory: make([][]map[int]int, m, m),
	}
	for i := range mapInfo.pathCountMemory {
		mapInfo.pathCountMemory[i] = make([]map[int]int, n, n)
		for j := range mapInfo.pathCountMemory[i] {
			mapInfo.pathCountMemory[i][j] = make(map[int]int)
		}
	}

	totalCount := 0
	for moveCount := 1; moveCount <= maxMove; moveCount++ {
		totalCount += findPathCountToExit(startRow, startColumn, moveCount)
		totalCount %= modNum
	}

	return totalCount
}

func findPathCountToExit(startRow int, startColumn int, moveCount int) int {
	if startRow < 0 || startRow >= mapInfo.rows || startColumn < 0 || startColumn >= mapInfo.columns {
		if moveCount == 0 {
			return 1
		}
		return 0
	}
	if moveCount == 0 {
		return 0
	}

	if storedCount, exists := mapInfo.pathCountMemory[startRow][startColumn][moveCount]; exists {
		return storedCount
	}

	count := 0
	for i := 0; i < 4; i++ {
		adjacentRow, adjacentColumn := startRow+dxy[i][0], startColumn+dxy[i][1]
		count += findPathCountToExit(adjacentRow, adjacentColumn, moveCount-1)
		count %= modNum
	}
	mapInfo.pathCountMemory[startRow][startColumn][moveCount] = count
	return count
}
