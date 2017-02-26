package services

import "models"

type WorkoutMovement struct {
	Workout  int64
	Movement int64
	models.MovementTemplate
}

const QUERY_WORKOUT_MOVEMENT = "select m.id, w.id, m.name, wm.repeats, wm.target_weight, wm.sets, wm.need_warm_set " +
	"from movement m , workout w , workout_movement wm where w.id = :workout and m.id = wm.movement and w.id = wm.workout"

func getWorkoutMovements(workoutId int64) []WorkoutMovement {
	movements := make([]WorkoutMovement, 0)
	rows := models.BasicCRUD.BuildAndQuery(QUERY_WORKOUT_MOVEMENT,
		map[string]interface{}{"workout": workoutId})
	for rows.Next() {
		var movement WorkoutMovement
		rows.Scan(&movement.Movement, &movement.Workout, &movement.Name,
			&movement.Repeats, &movement.Weight, &movement.Sets, &movement.NeedWarmSet)
		movements = append(movements, movement)
	}
	return movements
}

func (this *WorkoutMovement) GetWorkoutMovements() []WorkoutMovement {
	return getWorkoutMovements(this.Workout)
}
