package services

import "models"
import "encoding/json"

type TemplateWorkout struct {
	Template int64
	Workout  int64
}

type WorkoutDefinition struct {
	Name         string                    `json:"name"`
	Movements    []models.MovementTemplate `json:"movements"`
	StartAt      string                    `json:"startAt"`
	Weekly       string                    `json:"weekly"`
	Addition     string                    `json:"addition"`
	TargetMuscle string                    `json:"targetMuscle"`
	Description  string                    `json:"description"`
	Template     string                    `json:"template"`
}

func (this *WorkoutDefinition) AssignWorkoutsToTemplate() {
	template := findTemplate(this.Template)
	var id int64
	workouts := createAndSaveWorkouts(this)
	if template == nil || template.Id == 0 {
		newTemplate := new(Template)
		newTemplate.Name = this.Template
		content, _ := json.Marshal(this)
		newTemplate.Content = string(content)
		id = saveTemplate(*newTemplate)

	} else {
		id = template.Id
	}
	batchInsert := new(models.BatchSqls)
	for _, workout := range workouts {
		batchInsert.Params = append(batchInsert.Params, map[string]interface{}{
			"template": id,
			"workout":  workout.Id})
	}
	batchInsert.BatchInsert("template_workout")
}

func (this *TemplateWorkout) FindWorkouts() []*Workout {
	return findWorkoutsFromTemplates([]int64{this.Template})
}

func findWorkoutsFromTemplates(templatesId []int64) []*Workout {
	workouts := make([]*Workout, 0)
	rows := models.BasicCRUD.BuildAndQuery("select w.id , w.name, w.target, w.sequence, w.repeating from workout w, template_workout tw where tw.template in :template and tw.workout = w.id and w.sequence = 0",
		map[string]interface{}{"template": templatesId})
	for rows.Next() {
		workout := new(Workout)
		rows.Scan(&workout.Id, &workout.Name, &workout.Target, &workout.Sequence, &workout.Repeating)
		workouts = append(workouts, workout)
	}
	return workouts
}
