package services

import (
	"controllers"
	"encoding/json"
	"testing"
	"time"
)

var TEST_TEMPLATE_RAW string = "{\"workout\":\"template name\",\"movements\":[{\"name\":\"Bench Press\",\"weight\":\"85\",\"repeats\":\"6\",\"sets\":\"3\"},{\"name\":\"Barbell Pull\",\"weight\":\"65\",\"repeats\":\"6\",\"sets\":\"3\"},{\"name\":\"Sitting Shoulder Press\",\"weight\":\"25\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Pull Up\",\"weight\":\"76\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Incline Bench Press\",\"weight\":\"21\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Barbell Curl\",\"weight\":\"35\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Dip\",\"weight\":\"76\",\"repeats\":\"10\",\"sets\":\"3\"}],\"startAt\":\"2016-11-16\",\"addition\":\"2.5\",\"weekly\":\"Tuesday\",\"targetMuscle\":\"Upper body\",\"description\":\"Here I input some note for myself\"}"

func testCreateAndSaveWorkout(t *testing.T) {
	var result controllers.WorkoutTemplate
	err := json.Unmarshal([]byte(TEST_TEMPLATE_RAW), &result)
	workouts := createAndSaveWorkouts(template)
	if len(workouts) != 4 {
		t.Error("generate incorrect number of workouts from template")
	}
	for _, workout := range workouts {
		if workout.Id == nil {
			t.Errorf("unable to get workout id:%#v", workout)
		}
	}
}

func TestTimeParse(t *testing.T) {
	timePoint, err := time.Parse("2006-01-02", "2016-11-21")
	if err != nil {
		t.Error(err)
	}
	t.Log(timePoint.String()[0:10])
}
