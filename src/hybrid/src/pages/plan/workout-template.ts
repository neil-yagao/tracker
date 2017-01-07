import { Exercise } from './exercise';
export class WorkoutTemplate {
    exercises: Array<Exercise> = [];
    name: string;
    description: string;
    startAt: string;
	weekly: string;
	addition: string;
	targetMuscle: string;

    constructor() {
        this.startAt = new Date().toString();
    }
    addExercise(active: Exercise) {
        this.exercises.push(active);
    }
}
