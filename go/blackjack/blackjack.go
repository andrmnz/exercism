package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
  var option string
  dealersHand := ParseCard(dealerCard)
  playersHand := ParseCard(card1) + ParseCard(card2)
	switch {
	case playersHand == 22:
		option = "P"
	case playersHand == 21 && dealersHand != 11 && dealersHand != 10:
		option = "W"
	case playersHand == 21 && (dealersHand == 11 || dealersHand == 10):
		option = "S"
	case playersHand >= 17 && playersHand <= 20:
		option = "S"
	case playersHand >= 12 && playersHand <= 16 && dealersHand < 7:
		option = "S"
	case playersHand >= 12 && playersHand <= 16 && dealersHand >= 7:
		option = "H"
	case playersHand <= 11:
		option = "H"
	}
  return option
}
