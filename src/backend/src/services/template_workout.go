package services

import (
	"models"
	"strings"
)

type templateWorkoutService struct {
}

var TemplateWorkoutService templateWorkoutService

func (this *Workout) FindTemplateWorkouts(template string) []*Workout {
	return findTemplateWorkouts(template)
}

func findTemplateWorkouts(name string) []*Workout {
	if strings.TrimSpace(name) == "" {
		return make([]*Workout, 0)
	}
	rows := models.BasicCRUD.BuildAndQuery(QUERY_TEMPLATE_WORKOUT_SQL, map[string]interface{}{"template": name})
	defer rows.Close()
	return scanWorkoutsResult(rows)
}

//TODO change insert into batch insert

type TemplateWorkout struct {
	Template int64
	Workout  int64
}

func assignWorkoutsToTemplate(template string, workouts []*Workout) {
	templateId := findTemplateId(template)
	if templateId == -1 {
		newTemplate := new(models.Template)
		newTemplate.Name = template
		saveTemplate(*newTemplate)
	}
	for _, workout := range workouts {
		models.BasicCRUD.Save("template_workout", TemplateWorkout{templateId, workout.Id})
	}
}
