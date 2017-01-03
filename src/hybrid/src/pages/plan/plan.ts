import { Component } from '@angular/core';

import { NavController } from 'ionic-angular';

import { MovementPage } from '../movement/movement';
import { HttpBase } from '../../app/httpbase';
import { WorkoutTemplate } from './workout-template';
import { Exercise } from './exercise';

@Component({
	selector: 'page-plan',
	templateUrl: 'plan.html',
	providers: [HttpBase]
})
export class PlanPage {

    private movements: Array<any>;
    private workoutTemplate: WorkoutTemplate;
    private activeExercise: Exercise;
	constructor(public navCtrl: NavController, public http: HttpBase) {
        http.get('movements').subscribe((movements) => {
            this.movements = movements
        })
        this.activeExercise = new Exercise();
        this.workoutTemplate = new WorkoutTemplate();

	}

}
