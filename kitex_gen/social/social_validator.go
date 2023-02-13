package social

import "fmt"

// IsValid 校验ToUserId，ActionType
func (p *RelationActionRequest) IsValid() error {
	if p.ToUserId <= 0 {
		return fmt.Errorf("field ToUserId rule failed, current value: %d", p.ToUserId)
	}
	if p.ActionType != 1 && p.ActionType != 2 {
		return fmt.Errorf("field ActionType rule failed, current value: %d", p.ActionType)
	}
	return nil
}
func (p *RelationListRequest) IsValid() error {
	if p.UserId <= 0 {
		return fmt.Errorf("field UserId rule failed, current value: %d", p.UserId)
	}
	return nil
}

// IsValid 校验UserId，MyId
func (p *RelationInfoRequest) IsValid() error {
	if p.UserId <= 0 {
		return fmt.Errorf("field UserId rule failed, current value: %d", p.UserId)
	}
	if p.MyId <= 0 {
		return fmt.Errorf("field MyId rule failed, current value: %d", p.MyId)
	}
	return nil
}

// IsValid 校验ToUserId
func (p *MessageChatRequest) IsValid() error {
	if p.ToUserId <= 0 {
		return fmt.Errorf("field ToUserId rule failed, current value: %d", p.ToUserId)
	}
	return nil
}

// IsValid 校验ToUserId，ActionType，Content
func (p *MessageActionRequest) IsValid() error {
	if p.ToUserId <= 0 {
		return fmt.Errorf("field ToUserId rule failed, current value: %d", p.ToUserId)
	}
	if p.ActionType != 1 {
		return fmt.Errorf("field ActionType rule failed, current value: %d", p.ActionType)
	}
	if len(p.Content) == 0 {
		return fmt.Errorf("field Content rule failed, current value: %d", len(p.Content))
	}
	return nil
}
