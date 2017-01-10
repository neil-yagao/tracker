import { Component, ViewChild } from '@angular/core';

import { NavController, Tabs } from 'ionic-angular';

import { HttpBase } from '../../app/httpbase';
import { WorkoutTemplate } from './workout-template';
import { Exercise } from './exercise';

@Component({
	selector: 'page-plan',
	templateUrl: 'plan.html',
	providers: [HttpBase]
})
export class PlanPage {
	@ViewChild('planSlide') planSlide;
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

	goNextSlide(template: WorkoutTemplate) {
		this.workoutTemplate = template;
		this.planSlide.slideTo(1);
	}

	toMain() {
		this.planSlide.slideTo(0);
		this.workoutTemplate = new WorkoutTemplate();
		var t: Tabs = this.navCtrl.parent;
		t.select(0);
	}

}
