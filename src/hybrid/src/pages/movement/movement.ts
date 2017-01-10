import { Component } from '@angular/core';

import { NavController, ModalController } from 'ionic-angular';

import { HttpBase } from '../../app/httpbase';

import { MovementModal } from './movement-modal';


@Component({
	selector: 'page-movement',
	templateUrl: 'movement.html',
	providers: [HttpBase]
})
export class MovementPage {

    private movements: Array<any>;
	private selected: string;

    constructor(public navCtrl: NavController, public http: HttpBase, public modalCtrl: ModalController) {
        http.get('movements').subscribe((movements) => {
            this.movements = movements
        })
    }

	showModal(mv: any, title?: string) {
		let movement = this.modalCtrl.create(MovementModal, { movement: mv, 'title': (title ? "" : title) }, { showBackdrop: false, enableBackdropDismiss: false });
		movement.present();
	}
	getMovements() {
		return this.movements;
	}

	newMovement() {
		this.showModal({}, "添加动作")
	}

}
