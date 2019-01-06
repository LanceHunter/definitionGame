package definitionGame_test

import (
	"fmt"
	"testing"
)

func TestDefinitionGame(*testing.T) {
	if testing.Short() {
		fmt.Println("Testing in short mode.")
	}
	fmt.Println("This is the test.")
}

func BenchmarkGame(*testing.B) {
	fmt.Println("You're running a benchmark. Good for you")
}
