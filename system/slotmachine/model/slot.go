package model

type Card struct {
	Name       string
	Symbol     string
	IsScatter  bool
	IsWild     bool
	Payrate    []int
	Multiplyer int
}

type Reel struct {
	Cards []Card
}

type Machine struct {
	Reels []Reel
}

type StartConfig struct {
	NumOfReels int
	WindowSize int

	BaseCredit     int
	CentBase       int
	SpinMultiplyer int
	Bet            float64
}

type AuditResult struct {
	Stats        map[string]HiveStats
	Multiplyer   int
	ScatterCount int
}

type ReelAuditStats struct {
	Stats           map[string]HiveStats
	Multiplyer      int
	ContainsWild    bool
	ContainsScatter bool
}

type HiveStats struct {
	Name       string
	PayLevel   int
	Count      int
	Multiplyer int
}
