import { Component } from '@angular/core';
import { NavController} from 'ionic-angular';
import { TabsPage } from '../pages/tabs/tabs';


@Component({
	templateUrl: 'login.html'
})

export class LoginPage {

	username: string
	constructor(public nav: NavController) {
		this.nav = nav;
	}

	login() {
		// set the scrollLeft to 0px, and scrollTop to 500px
		// the scroll duration should take 200ms
		window.localStorage.setItem('username', this.username)
		console.debug("login!" + window.localStorage.getItem('username'))
		this.nav.setRoot(TabsPage)
	}
}
