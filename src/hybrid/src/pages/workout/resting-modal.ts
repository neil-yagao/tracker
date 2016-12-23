import { Component} from '@angular/core';
import { NavParams, ViewController } from 'ionic-angular';
import { Observable }     from 'rxjs/Observable';

@Component({
	selector: 'resting-modal',
	templateUrl: 'resting-modal.html'
})
export class RestingModal {
    private countDown: number;
    private percentage: string;
    private subscription: any;
    constructor(params: NavParams, public viewCtrl: ViewController) {
        var source = Observable
			.interval(1000)
            .timeInterval()
            .take(600);
        this.percentage = "p100"
        this.countDown = params.get("timeout");
		this.subscription = source.subscribe((x) => {
            this.countDown--;
            let percentage = Math.floor(this.countDown / params.get("timeout") * 100) + 1;
            if (percentage > 100) {
                percentage = 100
            } else if (percentage <= 0) {
                percentage = 0
            }
            this.percentage = "p" + String(percentage);
            console.info(this.percentage)
        }, (error) => { }, () => {
            this.dismiss()
        });

    }

    dismiss() {
        this.viewCtrl.dismiss();
        this.subscription.unsubscribe(4)
    }

}
