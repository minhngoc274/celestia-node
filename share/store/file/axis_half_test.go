package file

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/celestia-node/share/testing/sharetest"
)

func TestExtendAxisHalf(t *testing.T) {
	shares := sharetest.RandShares(t, 16)

	original := AxisHalf{
		Shares:   shares,
		IsParity: false,
	}

	extended, err := original.Extended()
	require.NoError(t, err)

	parity := AxisHalf{
		Shares:   extended[len(shares):],
		IsParity: true,
	}

	parityExtended, err := parity.Extended()
	require.NoError(t, err)

	require.Equal(t, extended, parityExtended)
}
