import { Component } from '@angular/core';

import { NavController, ModalController } from 'ionic-angular';

import { HttpBase } from '../../app/httpbase';

import { MovementModal } from './movement-modal';
import * as _ from "lodash";

@Component({
	selector: 'page-movement',
	templateUrl: 'movement.html',
	providers: [HttpBase]
})
export class MovementPage {

    private movements: Array<any>;
	private selected: string;
	private criteria: string = "";
	private allMovements: Array<any>;

    constructor(public navCtrl: NavController, public http: HttpBase, public modalCtrl: ModalController) {
        this.refreshMovementFromBackend()
    }

	refreshMovementFromBackend() {
		this.http.get('movements').subscribe((movements) => {
            this.movements = movements;
			this.allMovements = movements;
        })
	}

	filter(criteria) {
		this.movements = this.allMovements;
		this.movements = _.filter(this.movements, function(movement) {

			return movement.name.indexOf(this.criteria) >= 0
		}.bind(this))
	}

	showModal(mv: any, title?: string) {
		let movement = this.modalCtrl.create(MovementModal, { movement: mv, 'title': (title ? title : "") }, { showBackdrop: false, enableBackdropDismiss: false });
		movement.onDidDismiss(needRefresh => {
			if (needRefresh) {
				this.refreshMovementFromBackend();
			}
		})
		movement.present();
	}
	getMovements() {
		return this.movements;
	}

	newMovement() {
		this.showModal({}, "添加动作")
	}

}
