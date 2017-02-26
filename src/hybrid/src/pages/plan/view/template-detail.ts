import { Component } from '@angular/core';

import { NavController, NavParams /*, PopoverController */ } from 'ionic-angular';

import { HttpBase } from '../../../app/httpbase';
import { PlanPage } from '../create/plan';
import { SessionMovements} from './session-movements';

@Component({
    selector: 'template-detail',
    templateUrl: 'template-detail.html'
})
export class TemplateDetail {
    workouts: Array < any > ;
    template: any;
    constructor(public navCtrl: NavController, public params: NavParams,
        public http: HttpBase /*, public popCtrl: PopoverController*/ ) {
        this.template = params.data.template;
        this.freshWorkoutList()
    }

    freshWorkoutList() {
        let id = this.template.id;
        this.http.get('templates/' + id + '/workouts').subscribe((workouts) => {
            this.workouts = workouts
        })
    }

    goToDetail(workout) {
        this.navCtrl.push(SessionMovements, {'workout': workout, 'workouts': this.workouts})
    }

    openPopover(event) {
        this.navCtrl.push(PlanPage);

    }

    ionViewWillEnter() {
        this.freshWorkoutList()
    }
}
