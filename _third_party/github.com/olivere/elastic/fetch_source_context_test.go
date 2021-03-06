package elastic

import (
	"encoding/json"
	"testing"
)

func TestFetchSourceContextNoFetchSource(t *testing.T) {
	builder := NewFetchSourceContext(false)
	data, err := json.Marshal(builder.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `false`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestFetchSourceContextNoFetchSourceIgnoreIncludesAndExcludes(t *testing.T) {
	builder := NewFetchSourceContext(false).Include("a", "b").Exclude("c")
	data, err := json.Marshal(builder.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `false`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestFetchSourceContextFetchSource(t *testing.T) {
	builder := NewFetchSourceContext(true)
	data, err := json.Marshal(builder.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"excludes":[],"includes":[]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestFetchSourceContextFetchSourceWithIncludesAndExcludes(t *testing.T) {
	builder := NewFetchSourceContext(true).Include("a", "b").Exclude("c")
	data, err := json.Marshal(builder.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"excludes":["c"],"includes":["a","b"]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
