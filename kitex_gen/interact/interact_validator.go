package interact

import "fmt"

func (p *FavoriteActionRequest) IsValid() error {
	if p.VideoId < 0 {
		return fmt.Errorf("field VideoId rule failed, current value: %d", p.VideoId)
	}
	if p.ActionType != 1 && p.ActionType != 2 {
		return fmt.Errorf("field ActionType rule failed, current value: %d", p.ActionType)
	}
	return nil
}

func (p *FavoriteListRequest) IsValid() error {
	if p.UserId <= 0 {
		return fmt.Errorf("field UserId rule failed, current value: %d", p.UserId)
	}
	return nil
}

func (p *CommentActionRequest) IsValid() error {
	if p.VideoId <= 0 {
		return fmt.Errorf("field VideoId rule failed, current value: %d", p.VideoId)
	}
	if p.ActionType == 1 {
		if len(p.CommentText) == 0 {
			return fmt.Errorf("field CommentText rule failed, current value: %d", len(p.CommentText))
		}
		return nil
	}
	if p.ActionType == 2 {
		if p.CommentId <= 0 {
			return fmt.Errorf("field CommentId rule failed, current value: %d", p.CommentId)
		}
		return nil
	}
	return fmt.Errorf("field ActionType rule failed, current value: %d", p.ActionType)
}

func (p *CommentListRequest) IsValid() error {
	if p.VideoId <= 0 {
		return fmt.Errorf("field VideoId rule failed, current value: %d", p.VideoId)
	}
	return nil
}

func (p *VideoInteractRequest) IsValid() error {
	if p.UserId < 0 {
		return fmt.Errorf("field UserId rule failed, current value: %d", p.UserId)
	}
	if p.VideoId <= 0 {
		return fmt.Errorf("field VideoId rule failed, current value: %d", p.VideoId)
	}
	return nil
}
