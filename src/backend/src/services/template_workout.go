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
	if template == nil {
		newTemplate := new(Template)
		newTemplate.Name = this.Template
		content, _ := json.Marshal(this)
		newTemplate.Content = string(content)
		id = saveTemplate(*newTemplate)

	} else {
		id = template.Id
	}
	for _, workout := range workouts {
		models.BasicCRUD.Save("template_workout", TemplateWorkout{id, workout.Id})
	}
}

func findWorkoutsFromTemplates(templatesId []int64) []*Workout {
	workouts := make([]*Workout, 0)
	rows := models.BasicCRUD.BuildAndQuery("select w.id , w.name, w.target, w.sequence, w.repeat from workout w, template_workout tw where tw.template in :template and tw.workout = w.id",
		map[string]interface{}{"template": templatesId})
	for rows.Next() {
		workout := new(Workout)
		rows.Scan(&workout.Id, &workout.Name, &workout.Target, &workout.Sequence, &workout.Repeat)
		workouts = append(workouts, workout)
	}
	return workouts
}
