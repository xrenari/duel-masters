package fx

import (
	"duel-masters/game/cnd"
	"duel-masters/game/match"
	"fmt"
	"math/rand"
)

func GiveTapAbilityToAllies(card *match.Card, ctx *match.Context, alliesFilter func(card *match.Card) bool, tapAbility func(card *match.Card, ctx *match.Context)) {
	// This is added for the case where the card is added to the field. There is another creature
	// that doesn't initially have a tap abbility but should receive one. The change doesn't propagate fast
	// enough to the FE and that creature doesn't get tap ability until another action takes places.
	// This is an ugly workaround.
	FindFilter(
		card.Player,
		match.BATTLEZONE,
		alliesFilter,
	).Map(func(x *match.Card) {
		x.AddUniqueSourceCondition(cnd.TapAbility, tapAbility, card.ID)
	})

	ctx.Match.ApplyPersistentEffect(func(ctx2 *match.Context, exit func()) {
		if card.Zone != match.BATTLEZONE {
			Find(
				card.Player,
				match.BATTLEZONE,
			).Map(func(x *match.Card) {
				x.RemoveConditionBySource(card.ID)
			})

			exit()
			return
		}

		FindFilter(
			card.Player,
			match.BATTLEZONE,
			alliesFilter,
		).Map(func(x *match.Card) {
			x.AddUniqueSourceCondition(cnd.TapAbility, tapAbility, card.ID)
		})
	})
}

func FilterShieldTriggers(ctx *match.Context, filter func(*match.Card) bool) {

	if event, ok := ctx.Event.(*match.ShieldTriggerEvent); ok {
		validCards, invalidCards := FilterCardList(event.Cards, filter)
		event.Cards = validCards
		event.UnplayableCards = append(event.UnplayableCards, invalidCards...)
	}

}

func OpponentDiscardsRandomCard(card *match.Card, ctx *match.Context) {

	hand, err := ctx.Match.Opponent(card.Player).Container(match.HAND)

	if err != nil || len(hand) < 1 {
		return
	}

	discardedCard, err := ctx.Match.Opponent(card.Player).MoveCard(hand[rand.Intn(len(hand))].ID, match.HAND, match.GRAVEYARD, card.ID)
	if err == nil {
		ctx.Match.ReportActionInChat(ctx.Match.Opponent(card.Player), fmt.Sprintf("%s was discarded from %s's hand by %s", discardedCard.Name, discardedCard.Player.Username(), card.Name))
	}

}

// To be used as part of a card effect, not for initial shuffle
func ShuffleDeck(card *match.Card, ctx *match.Context, forOpponent bool) {
	if !forOpponent {
		card.Player.ShuffleDeck()
		ctx.Match.ReportActionInChat(card.Player, fmt.Sprintf("%s shuffled their deck", card.Player.Username()))
	} else {
		opponent := ctx.Match.Opponent(card.Player)
		opponent.ShuffleDeck()
		ctx.Match.ReportActionInChat(opponent, fmt.Sprintf("%s deck shuffled by %s effect", opponent.Username(), card.Name))
	}

}

func CantBeBlockedByPowerUpTo(card *match.Card, ctx *match.Context, power int) {
	blockersList := BlockersList(ctx)
	var newBlockersList []*match.Card
	for _, blocker := range *blockersList {
		if ctx.Match.GetPower(blocker, false) > power {
			newBlockersList = append(newBlockersList, blocker)
		}
	}
	*blockersList = newBlockersList
}

func CantBeBlockedByPowerUpTo4000(card *match.Card, ctx *match.Context) {
	CantBeBlockedByPowerUpTo(card, ctx, 4000)
}

func CantBeBlockedByPowerUpTo5000(card *match.Card, ctx *match.Context) {
	CantBeBlockedByPowerUpTo(card, ctx, 5000)
}

func CantBeBlockedByPowerUpTo8000(card *match.Card, ctx *match.Context) {
	CantBeBlockedByPowerUpTo(card, ctx, 8000)
}

func CantBeBlockedByPowerUpTo3000(card *match.Card, ctx *match.Context) {
	CantBeBlockedByPowerUpTo(card, ctx, 3000)
}

func RemoveBlockerFromList(card *match.Card, ctx *match.Context) {
	blockersList := BlockersList(ctx)
	var newBlockersList []*match.Card
	for _, blocker := range *blockersList {
		if blocker.ID != card.ID {
			newBlockersList = append(newBlockersList, blocker)
		}
	}
	*blockersList = newBlockersList
}
