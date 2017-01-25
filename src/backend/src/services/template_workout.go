package services

import (
	"models"
	"strings"
)

type templateWorkoutService struct {
}

var TemplateWorkoutService templateWorkoutService

func (this *templateWorkoutService) FindTemplateWorkouts(template string) []*models.Workout {
	return findTemplateWorkouts(template)
}

func findTemplateWorkouts(name string) []*models.Workout {
	if strings.TrimSpace(name) == "" {
		return make([]*models.Workout, 0)
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

func assignWorkoutsToTemplate(template string, workouts []*models.Workout) {
	templateId := findTemplateId(template)
	if templateId == -1 {
		saveTemplate(models.Template{0, template, " "})
	}
	for _, workout := range workouts {
		models.BasicCRUD.Save("template_workout", TemplateWorkout{templateId, workout.Id})
	}
}
