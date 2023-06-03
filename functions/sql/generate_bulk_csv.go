package sql

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type RowGenerator interface {
	generate(id int) string
}

func GenerateCSVParallel(batchIndex, batchParallelCount, numRows int, r RowGenerator, filenamePrefix string) {
	var wg sync.WaitGroup
	wg.Add(batchParallelCount)

	for i := 0; i < batchParallelCount; i++ {
		go func(fileIndex int) {
			defer wg.Done()

			file, err := os.Create(fmt.Sprintf("%s-%d.csv", filenamePrefix, batchParallelCount*batchIndex+fileIndex))
			if err != nil {
				panic(err)
			}
			defer file.Close()

			writer := bufio.NewWriter(file)
			defer writer.Flush()

			for j := 1; j <= numRows; j++ {
				row := r.generate((batchParallelCount*batchIndex+fileIndex)*numRows + j)
				writer.WriteString(row)
			}
		}(i)
	}

	wg.Wait()
}

type RowGeneratorFunc func(id int) string

// generate is the implementation of the generate method for RowGeneratorFunc
func (f RowGeneratorFunc) generate(id int) string {
	return f(id)
}
