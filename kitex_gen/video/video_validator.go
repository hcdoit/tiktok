package video

import (
	"fmt"
)

func (p *PublishActionRequest) IsValid() error {
	if len(p.Data) < int(1) {
		return fmt.Errorf("field Data len rule failed, current value: %d", len(p.Data))
	}
	if len(p.Title) < int(1) {
		return fmt.Errorf("field Title len rule failed, current value: %d", len(p.Title))
	}
	return nil
}
