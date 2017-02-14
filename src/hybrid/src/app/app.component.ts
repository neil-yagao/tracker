import { Component } from '@angular/core';
import { Platform,Config } from 'ionic-angular';
import { StatusBar, Splashscreen } from 'ionic-native';
import { TabsPage } from '../pages/tabs/tabs';
import { LoginPage } from './login';

@Component({
	templateUrl: 'app.html'
})
export class MyApp {
	rootPage = /*TabsPage;//*/LoginPage;

	constructor(platform: Platform, private config:Config) {
		platform.ready().then(() => {
			// Okay, so the platform is ready and our plugins are available.
			// Here you can do any higher level native things you might need.
			StatusBar.styleDefault();
			Splashscreen.hide();
			this.config.set("monthShortNames",["1月","2月","3月","4月","5月","6月",
			"7月","8月","9月","10月","11月","12月"])
		});
	}
}
