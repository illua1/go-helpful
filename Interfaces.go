package helpful

import (
	"golang.org/x/exp/constraints"
)

// Main helpful math type interface.
type Values interface {
	constraints.Float | constraints.Integer | byte
}
