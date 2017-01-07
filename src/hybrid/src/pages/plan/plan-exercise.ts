import { Component, Input} from '@angular/core';

import { WorkoutTemplate } from './workout-template';
import { Exercise } from './exercise';

import { HttpBase } from '../../app/httpbase';

import * as _ from "lodash";

@Component({
	selector: 'plan-exercise',
	templateUrl: 'plan-exercise.html',
	providers: [HttpBase]

})

export class PlanExercisePage {
    @Input() template: WorkoutTemplate;
    activeExercise: Exercise = new Exercise();
    private movements: Array<any>
    constructor(public http: HttpBase) {
        http.get('movements').subscribe((movements) => {
            this.movements = movements
        })
    }

    addExercise() {
        if (this.activeExercise.name) {
            this.template.addExercise(_.cloneDeep(this.activeExercise))
        }
    }
}
