// Package implementation for privacy transformation and sensitive-value protection.
package application

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ali/go-0821/privacy-transform-service/internal/domain"
)

func (s *Service) ImportPolicyWorkspace(ctx context.Context, data []byte) error {
	var in Export
	if err := json.Unmarshal(data, &in); err != nil {
		return fmt.Errorf("decode policy workspace: %v", err)
	}
	if err := s.CreatePolicyWorkspace(ctx, in.PolicyWorkspace); err != nil {
		return fmt.Errorf("create policy workspace: %v", err)
	}
	for _, e := range in.ProcessingPurposes {
		if err := s.CreateProcessingPurpose(ctx, e); err != nil {
			return fmt.Errorf("create processing purpose %q: %v", e.ID, err)
		}
	}
	for _, f := range in.TransformRuleSets {
		if err := s.CreateTransformRuleSet(ctx, f); err != nil {
			return fmt.Errorf("create transform rule set %q: %v", f.ID, err)
		}
	}
	return nil
}
func DecodeTransformRuleSet(data []byte) (domain.TransformRuleSet, error) {
	var f domain.TransformRuleSet
	if err := json.Unmarshal(data, &f); err != nil {
		return f, fmt.Errorf("decode transform rule set: %v", err)
	}
	return f, nil
}
