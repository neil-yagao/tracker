import { Component } from '@angular/core';

import { NavController } from 'ionic-angular';

<<<<<<< HEAD
@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {

  constructor(public navCtrl: NavController) {

  }
=======
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
>>>>>>> 509d4fa0bc794178ff4e28cc50578dd55784edbb

}
