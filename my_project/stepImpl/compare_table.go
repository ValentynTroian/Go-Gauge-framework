package stepImpl

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/getgauge-contrib/gauge-go/gauge"
	m "github.com/getgauge-contrib/gauge-go/models"
	. "github.com/getgauge-contrib/gauge-go/testsuit"
)

var _ = gauge.Step("Comparing files <table>", func(tbl *m.Table) {
	for _, row := range tbl.Rows {
		file1 := row.Cells[0]
		file2 := row.Cells[1]
		actualCount := calcCntDiffLinesTable(file1, file2)
		expectedCount := 0
		if actualCount != expectedCount {
			T.Fail(fmt.Errorf("Got: %d, Expected: %d", actualCount, expectedCount))
		}
	}
})

func calcCntDiffLinesTable(file1 string, file2 string) int {
	out, err := exec.Command("diff", "-y", "--suppress-common-lines", file1, file2).CombinedOutput()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
		default: //couldnt run diff
			log.Fatal(err)
		}
	}
	cntDiffLines := bytes.Count(out, []byte("\n"))

	output, err := exec.Command("pwd").CombinedOutput()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!1")
	fmt.Println(string(output))
	return cntDiffLines
}
