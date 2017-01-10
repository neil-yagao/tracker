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

	showModal(mv:any) {
		let movement = this.modalCtrl.create(MovementModal, { title: "从列表而来" }, { showBackdrop: false, enableBackdropDismiss: false });
		movement.present();
	}
	getMovements() {
		return this.movements;
	}

}
