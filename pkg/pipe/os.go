package pipe

import (
	"github.com/21stio/go-record/pkg/store"
	"github.com/21stio/go-record/pkg/types"
	"github.com/21stio/go-record/pkg/utils"
	"os/exec"
	"bytes"
	"github.com/21stio/go-record/pkg/e"
)

type OsPipe struct {
	Pipe
}

func (p OsPipe) Exec(dir s.GetString, command s.GetString, errH e.HandleError) (np StringPipe) {
	np.Ch = make(chan types.Ctx, 1000)
	np.Scope = p.Scope

	utils.Para(p.Scope + "__string.Prefix", nJobs, func() {
		ctx := <-p.Ch

		cmd := exec.Command(command.GetString(ctx))
		//cmd.Stdin = strings.NewReader("some input")
		cmd.Dir = dir.GetString(ctx)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			errH.HandleError(ctx, p.Ch, err)
			return
		}

		ctx.Val.String = out.String()

		np.Ch <- ctx
	})

	return
}