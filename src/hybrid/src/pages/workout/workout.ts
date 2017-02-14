import { Component } from '@angular/core';
import { NavController, PopoverController,LoadingController ,ModalController  } from 'ionic-angular';

import {Md5} from 'ts-md5/dist/md5';

import { HttpBase } from '../../app/httpbase';
import { URLSearchParams } from '@angular/http';
import { WorkoutDetail } from './workout-detail';
import { WorkoutTemplate } from './workout-template';
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
		private httpBase: HttpBase,  public modalCtrl:ModalController,
		private loading:LoadingController) {

	};

    goToDetail(workout) {
        this.navCtrl.push(WorkoutDetail, { workout: workout })
    };

	ionViewWillEnter() {
		let loader = this.loading.create({
		      content: "Please wait...",
		      duration: 3000
		    });
		loader.present()
		let param = new URLSearchParams();
		param.set('user', window.localStorage.getItem('username'));
		this.httpBase.get('workouts', param).subscribe(
			workouts => {
				this.workouts = _.sortBy(workouts,"perform_date")
				loader.dismiss()
		})

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
		let popover = this.pop.create(TemplatePopover)
		let template = this.modalCtrl.create(WorkoutTemplate,{}, { showBackdrop: false, enableBackdropDismiss: false });
		template.onDidDismiss( _ => {
			this.ionViewWillEnter()
		})
		popover.onDidDismiss((data) => {
			if(data){
				if(data.action == "add"){
					template.present();
				}
				this.ionViewWillEnter()
			}
		});
		popover.present({
			ev:event
		})

	}
}
