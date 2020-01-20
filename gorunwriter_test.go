package gorunewriter_test

import (
	"encoding/csv"
	"fmt"
	"github.com/marrbor/gorunewriter"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"os"
	"testing"
)

const (
	NoBreakSpace = "\u00A0"
	WaveDash     = "\u301C"
)

func TestRuneWriter_Write(t *testing.T) {
	file, err := os.Create("out.csv")
	assert.NoError(t, err)
	defer file.Close()

	writer := csv.NewWriter(&gorunewriter.RuneWriter{Writer: transform.NewWriter(file, japanese.ShiftJIS.NewEncoder())})
	writer.UseCRLF = true

	header := []string{
		"header1",
		"header2",
	}
	body := []string{
		fmt.Sprintf("あ%sい", NoBreakSpace),
		fmt.Sprintf("十時 %s 十二時", WaveDash),
	}

	err = writer.Write(header)
	assert.NoError(t, err)

	err = writer.Write(body)
	assert.NoError(t, err)

	writer.Flush()
	err = writer.Error()
	assert.NoError(t, err)
}
