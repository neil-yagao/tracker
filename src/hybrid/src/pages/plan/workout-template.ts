import { Exercise } from './exercise';
export class WorkoutTemplate {
    movements: Array < Exercise > = [];
    name: string;
    description: string;
    weekly: string;
    addition: string;
    targetMuscle: string;
    template: string;

    constructor() {}
    addExercise(active: Exercise) {
        this.movements.push(active);
    }
}
