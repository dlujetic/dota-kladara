package dotagsi

import "strconv"

// Formating custom decimal type for struct
type KeepZero float64

func (f KeepZero) MarshalJSON() ([]byte, error) {
	if float64(f) == float64(f) {
		return []byte(strconv.FormatFloat(float64(f), 'f', 1, 32)), nil
	}
	return []byte(strconv.FormatFloat(float64(f), 'f', 1, 32)), nil
}

type State struct {
	Auth      *auth
	Provider  *provider
	Map       *dotamap
	Player    *player
	Hero      *hero
	Abilities map[string]*abilit_base
	Items     map[string]*item_base
	// Buildings 	map[string]*
	// Item_base	*item_base
	// Abilit_base	*abilit_base
}

type provider struct {
	Name      string
	Appid     int
	Version   int
	timestamp float32
}

type auth struct {
	Token string
}

type dotamap struct {
	Name                   string
	Matchid                int
	gametime               string
	clocktime              string
	daytime                string
	nightstalker_night     string
	gamestate              string
	win_team               string
	customgamename         string
	ward_purchase_cooldown string
}

type player struct {
	steamid         string
	name            string
	activity        string
	kills           int
	deaths          int
	assists         int
	last_hits       int
	denies          int
	kill_streak     int
	team_name       string
	gold            int
	gold_reliable   int
	gold_unreliable int
	gpm             int
	xpm             int
}

type hero struct {
	id               int
	name             string
	level            int
	alive            bool
	respawn_seconds  int
	buyback_cost     int
	buyback_cooldown int
	health           int
	max_health       int
	health_percent   int
	mana             int
	max_mana         int
	mana_percent     string
	silenced         bool
	stunned          bool
	disarmed         bool
	magicimmune      bool
	hexed            bool
	muted            bool
	// break bool
	has_debuff 		bool
}

type item_base struct {
	name          string
	contains_rune string
	can_cast      bool
	cooldown      string
	passive       string
	charges       string
}

type abilit_base struct {
	name           string
	level          int
	can_cast       bool
	passive        string
	ability_active bool
	cooldown       int
	ultimate       int
}
