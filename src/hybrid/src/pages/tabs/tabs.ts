import { Component } from '@angular/core';

import { WorkoutPage } from '../workout/workout';
import { AboutPage } from '../about/about';
import { MovementPage } from '../movement/movement';

@Component({
	templateUrl: 'tabs.html'
})
export class TabsPage {
	// this tells the tabs component which Pages
	// should be each tab's root Page
	tab1Root: any = WorkoutPage;
	tab2Root: any = AboutPage;
	tab3Root: any = MovementPage;

	constructor() {

	}
}