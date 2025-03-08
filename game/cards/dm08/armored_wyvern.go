package dm08

import (
	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
)

// RocketdiveSkyterror ...
func RocketdiveSkyterror(c *match.Card) {

	c.Name = "Rocketdive Skyterror"
	c.Power = 5000
	c.Civ = civ.Fire
	c.Family = []string{family.Armored Wyvern}
	c.ManaCost = 4
	c.ManaRequirement = []string{civ.Fire}

	c.Use(fx.Creature, fx.CantAttackPlayers, fx.CantBeAttacked, fx.PowerAttacker1000)
  
}
