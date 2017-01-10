import { NgModule, ErrorHandler } from '@angular/core';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { FormsModule }   from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';

//third party
import { DEBUG_LOGGER_PROVIDERS } from "angular2-logger/core"

import { MyApp } from './app.component';
import { PlanPage } from '../pages/plan/plan';
import { MovementPage } from '../pages/movement/movement';
import { MovementModal } from '../pages/movement/movement-modal';
import { WorkoutPage } from '../pages/workout/workout';
import { WorkoutDetail } from '../pages/workout/workout-detail';
import { LoginPage} from './login';
import { TabsPage } from '../pages/tabs/tabs';
import { RestingModal } from '../pages/workout/resting-modal';
import { PlanBasicPage } from '../pages/plan/plan-basic';
import { PlanExercisePage } from '../pages/plan/plan-exercise';

@NgModule({
	declarations: [
		MyApp,
		PlanPage,
		MovementPage,
		WorkoutPage,
		TabsPage,
		LoginPage,
		WorkoutDetail,
		RestingModal,
		PlanBasicPage,
		PlanExercisePage,
		MovementModal
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
		PlanPage,
		MovementPage,
		WorkoutPage,
		TabsPage,
		LoginPage,
		WorkoutDetail,
		RestingModal,
		PlanBasicPage,
		PlanExercisePage,
		MovementModal
	],
	providers: [{ provide: ErrorHandler, useClass: IonicErrorHandler },
		DEBUG_LOGGER_PROVIDERS]
})
export class AppModule { }
