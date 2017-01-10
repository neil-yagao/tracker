import { Component} from '@angular/core';
import { NavParams, ViewController } from 'ionic-angular';

@Component({
	selector: 'movement-modal',
	templateUrl: 'movement-modal.html'
})
export class MovementModal {

	private title:string = "修改动作";
    constructor(params: NavParams, public viewCtrl: ViewController) {
		if(params.get("title")){
			this.title = params.get("title");
		}

    }

    dismiss() {
        this.viewCtrl.dismiss();
    }

}
