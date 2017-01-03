import { Component } from '@angular/core';

import { NavController } from 'ionic-angular';

import { HttpBase } from '../../app/httpbase';

@Component({
	selector: 'page-movement',
	templateUrl: 'movement.html',
	providers: [HttpBase]
})
export class MovementPage {

    private movements: Array<any>
    constructor(public navCtrl: NavController, public http: HttpBase) {
        http.get('movements').subscribe((movements) => {
            this.movements = movements
        })
    }

	getMovements() {
		return this.movements;
	}

}
