import { Exercise } from './exercise';
export class WorkoutTemplate {
    private _exercises: Array<Exercise> = [];
    constructor() { }

    addExercise(active: Exercise) {
        this._exercises.push(active);
    }

    get exercises(): Array<Exercise> {
        return this._exercises;
    }
}
