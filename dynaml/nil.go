package dynaml

import (
	"github.com/vito/spiff/yaml"
)

type NilExpr struct{}

func (e NilExpr) Evaluate(Context) (yaml.Node, bool) {
	return nil, true
}
