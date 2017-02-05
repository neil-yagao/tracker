import { Component } from '@angular/core';
import { NavController, PopoverController  } from 'ionic-angular';

import {Md5} from 'ts-md5/dist/md5';

import { HttpBase } from '../../app/httpbase';
import { URLSearchParams } from '@angular/http';
import { WorkoutDetail } from './workout-detail';
import { TemplatePopover } from './template-popover';

import * as _ from 'lodash'

@Component({
	selector: 'page-workout',
	templateUrl: 'workout.html',
    providers: [HttpBase]
})
export class WorkoutPage {

    workouts: Array<any>;
	debounce = false;
	constructor(public navCtrl: NavController, public pop:PopoverController ,
		private httpBase: HttpBase) {

	};

    goToDetail(workout) {
        this.navCtrl.push(WorkoutDetail, { workout: workout })
    };

	ionViewWillEnter() {
		let param = new URLSearchParams();
		param.set('user', window.localStorage.getItem('username'));
		param.set('template', window.localStorage.getItem('template'));
		this.httpBase.get('workouts', param).subscribe(
			workouts => this.workouts = _.sortBy(workouts,"perform_date"))
	}

	getNoteColor(workout): any {
		var today:string =  new Date().toISOString();
		if (today.startsWith(workout.perform_date)){
			return {'color':'RoyalBlue'}
		}else if (workout.perform_date < today) {
			return { 'color': 'red' }
		}
		else {
			return {}
		}
	}

	openPopover(event){
		let popover =	this.pop.create(TemplatePopover)
		popover.onDidDismiss(_ => {
			this.ionViewWillEnter()
		})
		popover.present({
			ev:event
		})

	}
}
