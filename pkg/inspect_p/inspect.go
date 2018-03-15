package inspect_p

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/21stio/go-record/pkg/utils"
	"github.com/21stio/go-record/pkg/types"
	"fmt"
)

var nJobs = 10

func Print(in chan types.Ctx) (out chan types.Ctx) {
	out = make(chan types.Ctx, 1000)

	utils.Para("Print", nJobs, func() {
		ii := <-in
		spew.Dump(ii)
		out <- ii
	})

	return out
}

func Count(name string, steps int, in chan types.Ctx) (out chan types.Ctx) {
	out = make(chan types.Ctx, 1000)
	c := 0

	utils.Para("inspect.Counter", nJobs, func() {
		ii := <-in
		c++
		if c % steps == 0 {
			println(fmt.Sprintf("counter: %v %v", name, c))
		}

		out <- ii
	})

	return out
}