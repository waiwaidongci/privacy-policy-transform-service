package bravo002

import (
	"reflect"
	"testing"

	"github.com/ali/go-0821/privacy-transform-service/internal/domain"
)

func rules() []domain.Rule {
	return []domain.Rule{
		{ID: "late", Priority: 20, Tags: map[string]string{"region": "jp"}, Value: "late"},
		{ID: "early", Priority: 10, Tags: map[string]string{"region": "jp"}, Value: "early"},
	}
}

func TestEvaluationKeepsCallerOrder(t *testing.T) {
	in := rules()
	want := []string{in[0].ID, in[1].ID}
	_, _, err := domain.Evaluate(domain.TransformRuleSet{DefaultValue: "default", Rules: in}, nil, domain.EvaluationContext{Tags: map[string]string{"region": "jp"}})
	if err != nil {
		t.Fatal(err)
	}
	if got := []string{in[0].ID, in[1].ID}; !reflect.DeepEqual(got, want) {
		t.Fatalf("evaluation reordered caller rules: %v", got)
	}
}

func TestCloneRulesOwnsTags(t *testing.T) {
	in := rules()
	out := domain.CloneRules(in)
	out[0].Tags["region"] = "eu"
	if in[0].Tags["region"] != "jp" {
		t.Fatalf("cloned tag rewrote source: %#v", in[0].Tags)
	}
}

func TestFilteredRuleSetsOwnRules(t *testing.T) {
	in := []domain.TransformRuleSet{{ID: "one", Key: "email", Status: "draft", Rules: rules()}}
	out := domain.FilterTransformRuleSets(in, domain.TransformRuleSetFilter{Status: "draft"})
	out[0].Rules[0].Tags["region"] = "us"
	out[0].Rules = append(out[0].Rules, domain.Rule{ID: "extra"})
	if len(in[0].Rules) != 2 || in[0].Rules[0].Tags["region"] != "jp" {
		t.Fatalf("filtered result rewrote source: %#v", in[0].Rules)
	}
}

func TestRevisionRulesAppendIsolation(t *testing.T) {
	in := rules()
	revision := &domain.TransformRevision{Number: 7, Value: "default", Rules: in}
	_, _, err := domain.Evaluate(domain.TransformRuleSet{}, revision, domain.EvaluationContext{Tags: map[string]string{"region": "jp"}})
	if err != nil {
		t.Fatal(err)
	}
	revision.Rules = append(revision.Rules, domain.Rule{ID: "third", Priority: 30})
	if in[0].ID != "late" || in[1].ID != "early" {
		t.Fatalf("revision evaluation changed shared backing array: %#v", in)
	}
}
