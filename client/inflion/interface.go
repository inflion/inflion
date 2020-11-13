package inflion

import (
	flowv1 "github.com/inflion/inflion/client/inflion/typed/flow/v1alpha"
)

type Interface interface {
	FlowV1() flowv1.FlowV1alpha1Interface
}