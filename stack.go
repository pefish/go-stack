package go_stack

import (
	"runtime"
	"strconv"
	"strings"
)

type StackClass struct {
}

var Stack = StackClass{}

type Option struct {
	Skip                int  // 跳过前多少行
	Count               int  // 显示多少行，默认全部显示
	FilenameMustInclude string
	FilenameMustExclude string
}

func (this *StackClass) GetStack(opts Option) string {
	result := ``
	pc := make([]uintptr, 64)
	cnt := runtime.Callers(0, pc)
	skip := opts.Skip
	if skip > cnt {
		return "get stack: Skip too big"
	}
	skipCounter := 0
	countCounter := 0
	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i])
		name := fu.Name()
		if !strings.Contains(name, "runtime") && !strings.Contains(name, "go-stack") {
			file, line := fu.FileLine(pc[i] - 1)
			if opts.FilenameMustExclude != `` && strings.Contains(file, opts.FilenameMustExclude) {
				continue
			}
			if opts.FilenameMustInclude != `` && !strings.Contains(file, opts.FilenameMustInclude) {
				continue
			}
			if skipCounter < skip {
				skipCounter++
				continue
			}
			result += file + ` ` + strconv.FormatInt(int64(line), 10) + "\n"
			countCounter++
			if opts.Count != 0 && countCounter >= opts.Count {
				break
			}
		}
	}
	return result
}
