import { Component } from '@angular/core';
import { URLSearchParams } from '@angular/http';
import { NavController, NavParams ,AlertController,LoadingController/*, PopoverController */ } from 'ionic-angular';

import { HttpBase } from '../../../app/httpbase';
import { PlanPage } from '../create/plan';

import * as _ from "lodash";

@Component({
    selector: 'session-movements',
    templateUrl: 'session-movements.html'
})
export class SessionMovements {
    movements: Array < any > ;
    workout: any;
    templateWorkouts :Array<any>;
    constructor(public navCtrl: NavController, public params: NavParams,
        public http: HttpBase ,public alertCtrl: AlertController,
        private loading: LoadingController/*, public popCtrl: PopoverController*/ ) {
    	this.workout = this.params.data.workout;
    	this.templateWorkouts = this.params.data.workouts;
    	let id = this.workout.id;
    	this.http.get('workout/' + id + '/movements').subscribe( (movements) => {
    		this.movements = movements
    	} )
    }

    removeMovement(movement){
    	let prompt = this.alertCtrl.create({
		      title: '确认删除',
		      message: "是否要将该动作从训练课中删除？",
		      buttons: [
		        {   
		       		text: '取消'
		        },
		        {
		          text: '确定',
		          handler: data => {
		            console.log(movement);
                    let loader = this.loading.create();
                    loader.present()
                    var param: URLSearchParams = new URLSearchParams();
                    param.set('workouts', _.flatMap(this.templateWorkouts, 'id'));
                    param.set('movement', movement.id);
		            this.http.delete('/workout/movements', param).subscribe( _ => {
                        this.movements = _.remove(this.movements, function(m){
                            m.id = movement.id
                        })
                        loader.dismiss()
                    })

		          }
		        }
		      ]
		    });
    	prompt.present();
  	}
}
