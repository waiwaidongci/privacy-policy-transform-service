// Package implementation for privacy transformation and sensitive-value protection.
package repository

import (
	"fmt"
	"github.com/ali/go-0821/privacy-transform-service/internal/domain"
)

func validationError(kind error, message string) error {
	return fmt.Errorf("validate resource: %v: %s", kind, message)
}

func validatePolicyWorkspace(p domain.PolicyWorkspace) error {
	if p.ID == "" || p.Name == "" {
		return validationError(domain.ErrInvalid, "invalid workspace")
	}
	return nil
}
func validateProcessingPurpose(e domain.ProcessingPurpose) error {
	if e.ID == "" || e.PolicyWorkspaceID == "" || e.Name == "" {
		return validationError(domain.ErrInvalid, "invalid purpose")
	}
	return nil
}
func validateTransformRuleSet(f domain.TransformRuleSet) error {
	if err := f.Validate(); err != nil {
		return validationError(err, "invalid transform rule set")
	}
	return nil
}
func validateTransformRevision(v domain.TransformRevision, t domain.ValueType) error {
	return v.Validate(t)
}
