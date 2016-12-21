import { Component, ViewChild, AfterViewInit } from '@angular/core';
import { NavController, NavParams, Navbar } from 'ionic-angular';

@Component({
	selector: 'workout-detail',
	templateUrl: 'workout-detail.html'
})
export class WorkoutDetail implements AfterViewInit {

    @ViewChild(Navbar) navBar: Navbar;
    constructor(public params: NavParams) {
    }

    ngAfterViewInit() {
        this.navBar.setBackButtonText("返回列表");
	}
}
