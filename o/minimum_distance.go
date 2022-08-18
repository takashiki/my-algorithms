package o

import (
	"fmt"
	"math"
)

type (
	Block         map[string]bool
	BlockDistance map[string]int
)

const (
	MaxDistance = math.MaxInt32
)

func FindTheBlock(blocks []Block, needs []string) int {
	var blockCount = len(blocks)
	var blockDistanceVector = make([]BlockDistance, blockCount)

	for i, block := range blocks {
		blockDistanceVector[i] = make(BlockDistance)
		for _, need := range needs {
			blockDistanceVector[i][need] = MaxDistance

			if block[need] {
				blockDistanceVector[i][need] = 0
				continue
			}

			if i < 1 {
				continue
			}

			if blockDistanceVector[i-1][need] < MaxDistance {
				blockDistanceVector[i][need] = min(blockDistanceVector[i-1][need]+1, blockDistanceVector[i][need])
			}
		}
	}

	fmt.Printf("%#v\n", blockDistanceVector)

	for i := blockCount - 1; i >= 0; i-- {
		block := blocks[i]
		for _, need := range needs {
			if block[need] {
				blockDistanceVector[i][need] = 0
				continue
			}

			if i >= blockCount-1 {
				continue
			}

			if blockDistanceVector[i+1][need] < MaxDistance {
				blockDistanceVector[i][need] = min(blockDistanceVector[i+1][need]+1, blockDistanceVector[i][need])
			}
		}
	}

	fmt.Printf("%#v\n", blockDistanceVector)

	minimumDistance, minimumIndex := MaxDistance, 0
	for i, blockDistance := range blockDistanceVector {
		totalDistance := 0
		for _, distance := range blockDistance {
			totalDistance += distance
		}
		if totalDistance < minimumDistance {
			minimumDistance = totalDistance
			minimumIndex = i
		}
	}

	return minimumIndex
}
