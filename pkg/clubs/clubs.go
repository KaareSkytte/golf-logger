package clubs

import "github.com/google/uuid"

type Bag struct {
	PlayerID uuid.UUID
	Clubs    []Club
}

type Club struct {
	ClubName string
	ClubType string
	Distance int
	InBag    bool
}

var AllPossibleClubs = []Club{
	{"Driver", "Wood", 0, false},
	{"3-wood", "Wood", 0, false},
	{"5-wood", "Wood", 0, false},
	{"Hybrid-3", "Hybrid", 0, false},
	{"Hybrid-4", "Hybrid", 0, false},
	{"Hybrid-5", "Hybrid", 0, false},
	{"3-iron", "Iron", 0, false},
	{"4-iron", "Iron", 0, false},
	{"5-iron", "Iron", 0, false},
	{"6-iron", "Iron", 0, false},
	{"7-iron", "Iron", 0, false},
	{"8-iron", "Iron", 0, false},
	{"9-iron", "Iron", 0, false},
	{"Pitching Wedge", "Wedge", 0, false},
	{"Gap Wedge", "Wedge", 0, false},
	{"Sand Wedge", "Wedge", 0, false},
	{"Lop Wedge", "Wedge", 0, false},
	{"Putter", "Putter", 0, false},
}
