package stepImpl

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/getgauge-contrib/gauge-go/gauge"
	. "github.com/getgauge-contrib/gauge-go/testsuit"
)

var _ = gauge.Step("Comparing <file1> with <file2>", func(file1 string, file2 string) {
	actualCount := calcCntDiffLines(file1, file2)
	expectedCount := 0
	if actualCount != expectedCount {
		T.Fail(fmt.Errorf("Got: %d, Expected: %d", actualCount, expectedCount))
	}
})

func calcCntDiffLines(file1 string, file2 string) int {
	out, err := exec.Command("diff", "-y", "--suppress-common-lines", file1, file2).CombinedOutput()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
		default: //couldnt run diff
			log.Fatal(err)
		}
	}
	cntDiffLines := bytes.Count(out, []byte("\n"))
	return cntDiffLines
}
