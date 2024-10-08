package consensus

import "errors"

var (
	// ErrUnknownAncestor is returned when validating a block requires an ancestor
	// that is unknown.
	ErrUnknownAncestor = errors.New("unknown ancestor")

	// ErrUnknownAncestorTD is returned when validating a block requires an ancestor
	// whose total difficulty is unknown.
	ErrUnknownAncestorTD = errors.New("unknown ancestor TD")

	// ErrPrunedAncestor is returned when validating a block requires an ancestor
	// that is known, but the state of which is not available.
	ErrPrunedAncestor = errors.New("pruned ancestor")

	// ErrFutureBlock is returned when a block's timestamp is in the future according
	// to the current node.
	ErrFutureBlock = errors.New("block in the future")

	// ErrInvalidNumber is returned if a block's number doesn't equal its parent's
	// plus one.
	ErrInvalidNumber = errors.New("invalid block number")

	// ErrUnexpectedWithdrawals is returned if a pre-Shanghai block has withdrawals.
	ErrUnexpectedWithdrawals = errors.New("unexpected withdrawals")
)
