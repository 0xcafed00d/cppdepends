package cppdep

import (
	"testing"

	"github.com/simulatedsimian/assert"
)

func TestExtract(t *testing.T) {
	assert := assert.Make(t)

	assert(processLine(`#include <stdio>`)).Equal("")
	assert(processLine(`#include "wibble.h"`)).Equal("wibble.h")

}
