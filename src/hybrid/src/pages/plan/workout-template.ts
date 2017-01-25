import { Exercise } from './exercise';
export class WorkoutTemplate {
    movements: Array<Exercise> = [];
    name: string;
    description: string;
    startAt: string;
	weekly: string;
	addition: string;
	targetMuscle: string;
    template:string;

    constructor() {
        this.startAt = new Date().toString();
    }
    addExercise(active: Exercise) {
        this.movements.push(active);
    }
}
