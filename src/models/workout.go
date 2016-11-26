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
}

type Movement struct {
	Id              int64
	TargetMuscle    string
	SecondaryMuscle string
	Name            string
	Description     string
}

type MovementTemplate struct {
	Name    string `json:"name"`
	Repeats string `json:"repeats"`
	Weight  string `json:"weight"`
	Sets    string `json:"sets"`
}

type WorkingSet struct {
	Id            int64   `json:"id"`
	Workout       int64   `json:"workout"`
	Movement      int64   `json:"movement"`
	MovementName  string  `json:"movementName"`
	TargetWeight  float64 `json:"targetWeight"`
	TargetNumber  int     `json:"targetNumber"`
	AcheiveNumber int64   `json:"acheiveNumber"`
	Sequence      int8    `json:"sequence"`
}
