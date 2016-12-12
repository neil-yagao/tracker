package test

import (
	"controllers"
	"encoding/json"
	"testing"
)

var PARSE_STRING string = "{\"workout\":\"template name\",\"movements\":[{\"name\":\"Bench Press\",\"weight\":\"85\",\"repeats\":\"6\",\"sets\":\"3\"},{\"name\":\"Barbell Pull\",\"weight\":\"65\",\"repeats\":\"6\",\"sets\":\"3\"},{\"name\":\"Sitting Shoulder Press\",\"weight\":\"25\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Pull Up\",\"weight\":\"76\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Incline Bench Press\",\"weight\":\"21\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Barbell Curl\",\"weight\":\"35\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Dip\",\"weight\":\"76\",\"repeats\":\"10\",\"sets\":\"3\"}],\"startAt\":\"2016-11-16\",\"addition\":\"2.5\",\"weekly\":\"Tuesday\",\"targetMuscle\":\"Upper body\",\"description\":\"Here I input some note for myself\"}"

func TestJsonParseResut(t *testing.T) {
	var result controllers.WorkoutTemplate
	err := json.Unmarshal([]byte(PARSE_STRING), &result)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}
