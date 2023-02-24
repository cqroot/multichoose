package multichoose_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/multichoose"
)

func TestMultiChoose(t *testing.T) {
	mc := multichoose.New(5)
	require.Equal(t, 5, mc.Length())

	for i := 0; i < mc.Length(); i++ {
		require.Equal(t, false, mc.IsSelected(i))
	}

	for i := 0; i < mc.Length(); i++ {
		mc.Select(i)
		require.Equal(t, true, mc.IsSelected(i))
	}

	for i := 0; i < mc.Length(); i++ {
		mc.Deselect(i)
		require.Equal(t, false, mc.IsSelected(i))
	}

	mc.SetLimit(2)

	for i := 0; i < 2; i++ {
		mc.Select(i)
		require.Equal(t, true, mc.IsSelected(i))
	}
	for i := 2; i < mc.Length(); i++ {
		mc.Select(i)
		require.Equal(t, false, mc.IsSelected(i))
	}

	for i := 0; i < mc.Length(); i++ {
		mc.Deselect(i)
		require.Equal(t, false, mc.IsSelected(i))
	}
}

func TestMultiChooseToggle(t *testing.T) {
	mc := multichoose.New(5)

	for i := 0; i < mc.Length(); i++ {
		mc.Toggle(i)
		require.Equal(t, true, mc.IsSelected(i))
		mc.Toggle(i)
		require.Equal(t, false, mc.IsSelected(i))
	}
}
