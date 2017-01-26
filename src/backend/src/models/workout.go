package models

type Workout struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Target      string `json:"target"`
	PerformDate string `json:"perform_date"`
	Description string `json:"description"`
}

type WorkoutTemplate struct {
	Name         string             `json:"name"`
	Movements    []MovementTemplate `json:"movements"`
	StartAt      string             `json:"startAt"`
	Weekly       string             `json:"weekly"`
	Addition     string             `json:"addition"`
	TargetMuscle string             `json:"targetMuscle"`
	Description  string             `json:"description"`
	Template     string             `json:"template"`
}

type Movement struct {
	Id              int64  `json:"id"`
	TargetMuscle    string `json:"targetMuscle"`
	SecondaryMuscle string `json:"secondarMuscle"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Dividable       int8   `json:"dividable"`
}

type MovementTemplate struct {
	Name        string  `json:"name"`
	Repeats     int     `json:"repeats"`
	Weight      float64 `json:"weight"`
	Sets        int     `json:"sets"`
	NeedWarmSet int8    `json:"needWarmSet"`
}

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

type Template struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
