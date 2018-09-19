package coinGame

type CoinList []Coin

type CoinString string

// attempt to flip coins. If there is an error, don't flip it
func (coinList CoinList) Flip(i int) (CoinList, error) {
	switch i {
	case 0, len(coinList) - 1:
		goto FLIPCOIN
	default:
		if _, err := coinList[i].flipOuter(coinList[i - 1], coinList[i + 1]); err != nil{
			return nil, err
		}
		goto FLIPCOIN
	}
	
	FLIPCOIN:
	
	newList := make(CoinList, 4)
	copy(newList, coinList)
	newList[i] = Coin(newList[i]).flip()
	return newList, nil
}

// loop through and attempt to flip each coin
func (coinList CoinList) GetPermutations() []CoinList {
	var combinations []CoinList
	for i := 0; i < len(coinList); i++ {
		newList, err := coinList.Flip(i)
		if err == nil {
			combinations = append(combinations, newList) 
		}
	}
	return combinations
}

// find out if the move just made is reversible. We can do this
// by looking at the child's permutations and if we find that
// one of the child's permutations is equal to the parent, we
// know it is reversible
func (coinList CoinList) IsReversible(list CoinList) bool {
	permutations := list.GetPermutations()
	for _, permutation := range permutations {
		if coinList.String() == permutation.String() {
			return true
		}
	}
	return false
}

// turn the coin list into a string for map purposes
func (coinList CoinList) String() CoinString {
	var coinString CoinString
	for i := 0; i < len(coinList); i++ {
		coinString += CoinString(coinList[i])
	}
	return coinString
}

func (coinList CoinList) IsSame(other CoinList) bool  {
	coinListString := coinList.String()
	otherString := other.String()

	for i := 0; i < len(coinListString); i++ {
		if coinListString[i] != otherString[i] {
			return false
		}
	}

	return true
}