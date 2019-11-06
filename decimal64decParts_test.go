package decimal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPartsInf(t *testing.T) {
	a := Infinity64.getParts()
	require.True(t, a.isInf())

	b := NegInfinity64.getParts()
	require.True(t, b.isInf())
}

func TestIsNaN(t *testing.T) {
	require := require.New(t)
	a := Zero64.getParts()
	require.Equal(false, a.isNaN())

	b := SNaN64.getParts()
	require.Equal(true, b.isSNaN())

	c := QNaN64.getParts()
	require.Equal(true, c.isQNaN())
}

func TestPartsSubnormal(t *testing.T) {
	require := require.New(t)

	d := MustParseDecimal64("0.1E-383")
	subnormal64Parts := d.getParts()
	require.Equal(true, subnormal64Parts.isSubnormal())

	e := NewDecimal64FromInt64(42)
	fortyTwoParts := e.getParts()
	require.Equal(false, fortyTwoParts.isSubnormal())

}
