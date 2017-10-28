package processors

import (
	"charisma-beat/integration"
)

type FakeProcessor struct {
}

func (FakeProcessor) Process(event integration.Event) error {

	return nil
}
