import { Component } from '@angular/core';

import { NavController } from 'ionic-angular';

import { HttpBase } from '../../app/httpbase';
import { URLSearchParams } from '@angular/http';
import { WorkoutDetail } from './workout-detail'


@Component({
	selector: 'page-workout',
	templateUrl: 'workout.html',
    providers: [HttpBase]
})
export class WorkoutPage {

    workouts: Array<any>;

	constructor(public navCtrl: NavController, private httpBase: HttpBase) {
        let param = new URLSearchParams();
        param.set('user', window.localStorage.getItem('username'));
        this.workouts = [
            {
                name: 'UppderBody',
                description: 'this is a upperbody training session'
            },
            {
                name: 'LowerBody',
                description: 'this is a lowerbody training session'
            }
        ]
        // this.httpBase.get('workouts', param).subscribe(workouts => this.workouts = workouts)
	}

    goToDetail(workout) {
        this.navCtrl.push(WorkoutDetail, { workout: workout })
    }
}
