package models

type WorkingSet struct {
	Id            int64   `json:"id"`
	Workout       int64   `json:"workout"`
	Movement      int64   `json:"movement"`
	MovementName  string  `json:"movementName"`
	TargetWeight  float64 `json:"targetWeight"`
	AcheiveWeight float64 `json:"acheiveWeight"`
	TargetNumber  int     `json:"targetNumber"`
	AcheiveNumber int64   `json:"acheiveNumber"`
	Sequence      int8    `json:"sequence"`
}

type DividableWorkingSet struct {
	Dividable int8 `json:"dividable"`
	WorkingSet
}

type SessionUpdateInfo struct {
	Movement int64   `json:"movement"`
	WorkSet  []int64 `json:"workset"`
}
