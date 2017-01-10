import { Component} from '@angular/core';
import { NavParams, ViewController, AlertController } from 'ionic-angular';

import { HttpBase } from '../../app/httpbase';

import * as _ from "lodash";

@Component({
	selector: 'movement-modal',
	templateUrl: 'movement-modal.html',
	providers: [HttpBase]
})
export class MovementModal {

	private title: string = "修改动作";
	private movement: any;
    constructor(params: NavParams, public viewCtrl: ViewController,
		public alertCtrl: AlertController, public http: HttpBase) {
		let title: string = params.get("title");
		if (title && title.length > 0) {
			this.title = title;
		}
		this.movement = params.get("movement") || {};
		if (this.movement.targetMuscle) {
			this.movement.muscleGroup = _.split(this.movement.targetMuscle, ",");
		}
		else {
			this.movement.muscleGroup = [];
		}

    }

	combinedMuscle(targetArray: Array<string>) {
		this.movement.targetMuscle = ""
		_.forEach(targetArray, (muscle) => {
			this.movement.targetMuscle += muscle + ","
		});
		this.movement.targetMuscle = _.trimEnd(this.movement.targetMuscle, ",");
		this.movement.muscleGroup = targetArray;
	}

	saveOrUpdateMovement() {
		let confirm = this.alertCtrl.create({
			title: '修改动作',
			message: '是否更改动作？',
			buttons: [
				{
					text: '确定',
					handler: () => {
						var movementId = this.movement.id;
						if (movementId && movementId != -1) {
							//update
							this.http.post('movement/' + movementId, this.movement)
								.subscribe((respons) => {
									console.info(respons);
									this.viewCtrl.dismiss();
								})
						} else {
							this.http.put('movements', this.movement)
						}
					}
				}, {
					text: '取消',
					handler: () => {

					}
				}
			]

		})

		confirm.present()

	}

	dividableChange(newValue) {
		if (newValue) {
			this.movement.dividable = 1;
		} else {
			this.movement.dividable = 0;
		}

	}
    dismiss() {
        this.viewCtrl.dismiss();
    }

}
