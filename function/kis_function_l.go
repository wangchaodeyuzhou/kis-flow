package function

import (
	"context"
	"fmt"
	"kis-flow/flow"
)

type KisFunctionL struct {
	BaseFunction
}

func (f *KisFunctionL) Call(ctx context.Context, flow *flow.KisFlow) error {
	fmt.Printf("KisFunctionL, flow = %+v\n", flow)

	// TODO 调用具体的Function执行方法

	return nil
}