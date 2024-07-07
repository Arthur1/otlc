package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func NoDiff(t *testing.T, want, got any, opts []cmp.Option, msgAndArgs ...any) {
	assert.Empty(t, cmp.Diff(got, want, opts...), msgAndArgs...)
}
