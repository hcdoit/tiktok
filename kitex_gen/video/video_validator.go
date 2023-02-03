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

func (p *PublishListRequest) IsValid() error {
	if p.UserId <= 0 {
		return fmt.Errorf("field UserId rule failed, current value: %d", p.UserId)
	}
	return nil
}

func (p *FeedRequest) IsValid() error {
	if p.LatestTime < 0 {
		return fmt.Errorf("field LatestTime rule failed, current value: %d", p.LatestTime)
	}
	return nil
}

func (p *VideoRequest) IsValid() error {
	if p.VideoId < 0 {
		return fmt.Errorf("field VideoId rule failed, current value: %d", p.VideoId)
	}
	return nil
}
