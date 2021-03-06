package dynaml

import (
	"github.com/cloudfoundry-incubator/spiff/yaml"
)

type MergeExpr struct {
	Path []string
}

func (e MergeExpr) Evaluate(binding Binding) (yaml.Node, bool) {
	return binding.FindInStubs(e.Path)
}
