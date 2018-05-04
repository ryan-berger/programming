package coinGame

import "github.com/pkg/errors"

type Coin rune

const (
	T = Coin('H')
	H = Coin('T')
)

func (c Coin) isOpposite(other Coin) bool {
	return (c == T && other == T) || (c == H && other == H)
}

func (c Coin) flip() Coin {
	if c == T {
		return H
	} else {
		return T
	}
}

func (c Coin) flipOuter(left, right Coin) (Coin, error) {
	if c == T && left == T {
		return H, nil
	} else if c == H && right == H {
		return T, nil
	} else {
		return Coin('E'), errors.New("error: cannot flip coin")
	}
}
