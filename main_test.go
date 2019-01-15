package v

import (
	"testing"
	"fmt"
)

func TestPrint(t *testing.T) {
	res := Print1to20()
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

// depends test
func testPrint1(t *testing.T) {
	res := Print1to20()
	if res != 211 {
		t.Errorf("Wrong result of Print1to20")
	}
}
func testSkip(t *testing.T) {
	fmt.Println("TestSkip")
	t.SkipNow()
	fmt.Println("TestSkip")
}
// use t.Run to run subtests, to make sure they are run in fix orders
func TestDependsTest(t *testing.T) {
	// t.Run("a1", testPrint1)
	t.Run("a2", testSkip)
}

// enter func
// init before test cases run
func TestMain(m *testing.M) {
	fmt.Println("test main first")
	m.Run()
}


// bench test
func BenchmarkAll(b *testing.B) {
	fmt.Println("testBench")
	for i := 0; i < b.N; i++ {
		println(b.N)
		Print1to20()
	}
}