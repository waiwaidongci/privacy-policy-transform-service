// Package implementation for privacy transformation and sensitive-value protection.
package repository

import (
	"github.com/ali/go-0821/privacy-transform-service/internal/domain"
	"strings"
	"sync"
)

type TransformRuleSetIndex struct {
	mu                sync.RWMutex
	byPolicyWorkspace map[string][]string
	byKey             map[string]string
}

func NewTransformRuleSetIndex() *TransformRuleSetIndex {
	return &TransformRuleSetIndex{byPolicyWorkspace: map[string][]string{}, byKey: map[string]string{}}
}
func (i *TransformRuleSetIndex) Add(f domain.TransformRuleSet) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.byPolicyWorkspace[f.PolicyWorkspaceID] = append(i.byPolicyWorkspace[f.PolicyWorkspaceID], f.ID)
	i.byKey[strings.ToLower(f.Key)] = f.ID
}
func (i *TransformRuleSetIndex) Remove(f domain.TransformRuleSet) {
	i.mu.Lock()
	defer i.mu.Unlock()
	ids := i.byPolicyWorkspace[f.PolicyWorkspaceID]
	out := make([]string, 0, len(ids))
	for _, id := range ids {
		if id != f.ID {
			out = append(out, id)
		}
	}
	i.byPolicyWorkspace[f.PolicyWorkspaceID] = out
	delete(i.byKey, strings.ToLower(f.Key))
}
func (i *TransformRuleSetIndex) FindByKey(key string) string {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.byKey[strings.ToLower(key)]
}
