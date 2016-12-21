import { NgModule, ErrorHandler } from '@angular/core';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { FormsModule }   from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';

import { MyApp } from './app.component';
import { AboutPage } from '../pages/about/about';
import { MovementPage } from '../pages/movement/movement';
import { WorkoutPage } from '../pages/workout/workout';
import { WorkoutDetail } from '../pages/workout/workout-detail';
import { LoginPage} from './login';
import { HttpBase } from './httpbase';
import { TabsPage } from '../pages/tabs/tabs';

@NgModule({
	declarations: [
		MyApp,
		AboutPage,
		MovementPage,
		WorkoutPage,
		TabsPage,
		LoginPage,
		WorkoutDetail
	],
	imports: [
		FormsModule,
		IonicModule.forRoot(MyApp),
		HttpModule,
		JsonpModule

	],
	bootstrap: [IonicApp],
	entryComponents: [
		MyApp,
		AboutPage,
		MovementPage,
		WorkoutPage,
		TabsPage,
		LoginPage,
		WorkoutDetail
	],
	providers: [{ provide: ErrorHandler, useClass: IonicErrorHandler }]
})
export class AppModule { }
