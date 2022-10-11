package share

import (
	"context"
	"errors"
	"time"

	"github.com/ipfs/go-blockservice"
	format "github.com/ipfs/go-ipld-format"

	"github.com/celestiaorg/celestia-node/ipld"
)

type FAOption func(*FullAvailability)

// To be used with the construction
// example:
//
// NewFullAvailability(
//
//	bServ,
//	disc,
//	WithFullAvailabilityTimeout(10 * time.Minute),
//
// )
func WithFullAvailabilityTimeout(timeout time.Duration) FAOption {
	return func(fa *FullAvailability) {
		fa.timeout = timeout
	}
}

// Use for default timeout value
func WithFullAvailabilityTimeoutDefault() FAOption {
	return func(fa *FullAvailability) {
		fa.timeout = DefaultAvailabilityTimeout
	}
}

// FullAvailability implements Availability using the full data square
// recovery technique. It is considered "full" because it is required
// to download enough shares to fully reconstruct the data square.
type FullAvailability struct {
	timeout time.Duration

	rtrv *ipld.Retriever
	disc *Discovery

	cancel context.CancelFunc
}

// NewFullAvailability creates a new full Availability.
func NewFullAvailability(
	bServ blockservice.BlockService,
	disc *Discovery,
	options ...FAOption,
) *FullAvailability {
	fa := &FullAvailability{
		rtrv: ipld.NewRetriever(bServ),
		disc: disc,
	}

	// assign options if applicable
	for _, applyOpt := range options {
		applyOpt(fa)
	}

	return fa
}

func (fa *FullAvailability) Start(context.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	fa.cancel = cancel

	go fa.disc.advertise(ctx)
	go fa.disc.ensurePeers(ctx)
	return nil
}

func (fa *FullAvailability) Stop(context.Context) error {
	fa.cancel()
	return nil
}

// SharesAvailable reconstructs the data committed to the given Root by requesting
// enough Shares from the network.
func (fa *FullAvailability) SharesAvailable(ctx context.Context, root *Root) error {
	ctx, cancel := context.WithTimeout(ctx, fa.timeout) // timeout is configurable through functional options passed to the constructor
	defer cancel()
	// we assume the caller of this method has already performed basic validation on the
	// given dah/root. If for some reason this has not happened, the node should panic.
	if err := root.ValidateBasic(); err != nil {
		log.Errorw("Availability validation cannot be performed on a malformed DataAvailabilityHeader",
			"err", err)
		panic(err)
	}

	_, err := fa.rtrv.Retrieve(ctx, root)
	if err != nil {
		log.Errorw("availability validation failed", "root", root.Hash(), "err", err)
		if format.IsNotFound(err) || errors.Is(err, context.DeadlineExceeded) {
			return ErrNotAvailable
		}

		return err
	}
	return err
}

func (fa *FullAvailability) ProbabilityOfAvailability() float64 {
	return 1
}
