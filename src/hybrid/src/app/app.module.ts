import { NgModule, ErrorHandler } from '@angular/core';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { FormsModule } from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';

//third party
import { DEBUG_LOGGER_PROVIDERS } from "angular2-logger/core"

import { MyApp } from './app.component';
import { PlanPage } from '../pages/plan/create/plan';
import { MovementPage } from '../pages/movement/movement';
import { MovementModal } from '../pages/movement/movement-modal';
import { WorkoutPage } from '../pages/workout/workout';
import { WorkoutTemplate } from '../pages/workout/workout-template';
import { WorkoutDetail } from '../pages/workout/workout-detail';
import { TemplatePopover } from '../pages/workout/template-popover';
import { LoginPage } from './login';
import { TabsPage } from '../pages/tabs/tabs';
import { RestingModal } from '../pages/workout/resting-modal';
import { PlanBasicPage } from '../pages/plan/create/plan-basic';
import { PlanExercisePage } from '../pages/plan/create/plan-exercise';
import { PlanList } from '../pages/plan/view/plan-list';
import { TemplateDetail } from '../pages/plan/view/template-detail';
import { SessionMovements } from '../pages/plan/view/session-movements';

import { HttpBase } from './httpbase'

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
        MovementModal,
        TemplatePopover,
        WorkoutTemplate,
        PlanList,
        TemplateDetail,
        SessionMovements
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
        MovementModal,
        TemplatePopover,
        WorkoutTemplate,
        PlanList,
        TemplateDetail,
        SessionMovements
    ],
    providers: [
        HttpBase,
        { provide: ErrorHandler, useClass: IonicErrorHandler },
        DEBUG_LOGGER_PROVIDERS
    ]
})
export class AppModule {}
