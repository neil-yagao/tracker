import { Component, ViewChild, AfterViewInit } from '@angular/core';
import { URLSearchParams } from '@angular/http';
import { NavController, NavParams, Navbar } from 'ionic-angular';

import * as _ from "lodash";
import {Md5} from 'ts-md5/dist/md5';

import { HttpBase } from '../../app/httpbase';

@Component({
	selector: 'workout-detail',
	templateUrl: 'workout-detail.html',
	providers: [HttpBase]
})
export class WorkoutDetail implements AfterViewInit {

    @ViewChild(Navbar) navBar: Navbar;
	private session: any;
	private exercises: Array<any>;
    constructor(public params: NavParams, private http: HttpBase) {
		var workoutId = params.data.workout.id;
		let param = new URLSearchParams();
		param.set('user', String(Md5.hashStr('powerlift')));
		http.get('session/' + workoutId, param).subscribe(session => {
			var movements = _.groupBy(session, 'movementName')
			var beforeAssign = _.map(movements, function(value, key) {
				var sequence = 0;
				var newSets = _.map(value, function(set: any) {
					var eachSide = (set.targetWeight - 20) / 2;
					if (!set.dividable) {
						eachSide = 1 / 0
					}
					if (set.sequence != 0) {
						sequence += 1
					}
					return {
						'id': set.id,
						'sequence': set.sequence == 0 ? "*" : sequence,
						'repeat': set.targetNumber,
						'eachSide': eachSide,
						'totalWeight': set.targetWeight,
						"eachSideShow": false,
						'achieved': -1
					}
				})
				return {
					"name": key,

					"sets": newSets
				}
			})
			console.info(beforeAssign)
			this.exercises = beforeAssign;

		})

    }

	finishOneSet(set) {
		if (set.achieved == undefined || set.achieved <= 0) {
			set.achieved = set.repeat;
		} else {
			set.achieved -= 1
		}
	}


    ngAfterViewInit() {
        this.navBar.setBackButtonText("返回列表");
	}
}
