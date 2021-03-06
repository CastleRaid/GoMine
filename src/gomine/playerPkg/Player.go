package playerPkg

import "gomine/entityPkg"

type player struct {
	*entityPkg.Entity
	playerName string
	displayName string
}

func NewPlayer() player {
	return player{}
}

func (player *player) getName() string {
	return player.playerName
}

func (player *player) getDisplayName() string {
	return player.displayName
}

func (player *player) setDisplayName(name string) {
	player.displayName = name
}
