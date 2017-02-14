import { Component} from '@angular/core';
import { URLSearchParams } from '@angular/http';
import { NavController, AlertController } from 'ionic-angular';
import { TabsPage } from '../pages/tabs/tabs';
import {Md5} from 'ts-md5/dist/md5';
import { HttpBase } from './httpbase';

@Component({
	templateUrl: 'login.html',
	providers: [HttpBase]
})

export class LoginPage {

	username: string;
	template: string; constructor(public nav: NavController, public http: HttpBase,
								public alert: AlertController) {
		this.nav = nav;
	}

	login() {

		let usr: string = String(Md5.hashStr(this.username));
		window.localStorage.setItem('username', usr);
		let param = new URLSearchParams();
		param.set('user', usr);
		this.http.post('login', {
			'username': this.username,
			'userIdentity': usr
		}, param).subscribe((response) => {
			if (response.success) {
				console.debug("login!" + window.localStorage.getItem('username'))
				this.nav.setRoot(TabsPage)
			}
			else {
				this.alert.create({
					'title': '错误',
					'subTitle': response.reason
				}).present()
			}
		});

	}

}
