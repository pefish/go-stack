package go_stack

import (
	"runtime"
	"strconv"
	"strings"
)

type StackClass struct {

}

var Stack = StackClass{}

func (this *StackClass) GetStack(skip int) string {
	result := ``
	pc := make([]uintptr, 64)
	cnt := runtime.Callers(0, pc)
	if skip > cnt {
		return "get stack: skip too big"
	}
	skipCounter := 0
	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i])
		name := fu.Name()
		if !strings.Contains(name, "runtime") && !strings.Contains(name, "go-stack") {
			if skipCounter < skip {
				skipCounter++
				continue
			}
			file, line := fu.FileLine(pc[i] - 1)
			result += file + ` --- ` + strconv.FormatInt(int64(line), 10) + "\n"
		}
	}
	return result
}
