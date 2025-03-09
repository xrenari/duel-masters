package dm08

import (
	"duel-masters/game/civ"
	"duel-masters/game/cnd"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
	"fmt"
)
// KuukaiFinderOfKarma ...
func KuukaiFinderOfKarma(c *match.Card) {

	c.Name = "Kuukai, Finder of Karma"
	c.Power = 10500
	c.Civ = civ.Light
	c.Family = []string{family.MechaThunder}
	c.ManaCost = 5
	c.ManaRequirement = []string{civ.Light}

	c.Use(fx.Creature, fx.Blocker, fx.Evolution, fx.CantAttackPlayers, func(card *match.Card, ctx *match.Context) {

		if event, ok := ctx.Event.(*match.CreatureDestroyed); ok {

			if event.Source == card && event.Blocked {
				card.Tapped = false
			}

		}

	})

}
