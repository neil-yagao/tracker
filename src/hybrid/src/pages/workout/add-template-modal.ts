import { Component} from '@angular/core';
import { NavParams, ViewController } from 'ionic-angular';

import { HttpBase } from '../../app/httpbase';

@Component({
	selector: 'add-template-modal',
	templateUrl: 'add-template-modal.html',
    providers: [HttpBase]
})
export class AddTemplateModal {

    constructor(params: NavParams, public viewCtrl: ViewController, public http: HttpBase) {

    }

}