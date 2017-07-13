package cppdep

import (
	"strings"
	"testing"

	"github.com/simulatedsimian/assert"
)

func TestProcessLine(t *testing.T) {
	assert := assert.Make(t)

	assert(processLine(`#include <stdio.h>`)).Equal("")
	assert(processLine(`#include "wibble.h"`)).Equal("wibble.h")

}

var testfile = `#include <stdio.h>
#define wibble 0
  #include "hello.h"
#include "file.h"  

void main (){
	printf ("hello world");
}
`

func TestExtract(t *testing.T) {
	assert := assert.Make(t)

	assert(extractIncludes(strings.NewReader(testfile))).
		Equal([]string{"hello.h", "file.h"}, nil)
}
