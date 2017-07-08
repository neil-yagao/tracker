var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
import { NgModule, ErrorHandler } from '@angular/core';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { FormsModule } from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';
//third party
import { DEBUG_LOGGER_PROVIDERS } from "angular2-logger/core";
import { MyApp } from './app.component';
import { PlanPage } from '../pages/plan/plan';
import { MovementPage } from '../pages/movement/movement';
import { MovementModal } from '../pages/movement/movement-modal';
import { WorkoutPage } from '../pages/workout/workout';
import { WorkoutDetail } from '../pages/workout/workout-detail';
import { TemplatePopover } from '../pages/workout/template-popover';
import { LoginPage } from './login';
import { TabsPage } from '../pages/tabs/tabs';
import { RestingModal } from '../pages/workout/resting-modal';
import { PlanBasicPage } from '../pages/plan/plan-basic';
import { PlanExercisePage } from '../pages/plan/plan-exercise';
var AppModule = (function () {
    function AppModule() {
    }
    return AppModule;
}());
AppModule = __decorate([
    NgModule({
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
            TemplatePopover
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
            TemplatePopover
        ],
        providers: [{ provide: ErrorHandler, useClass: IonicErrorHandler },
            DEBUG_LOGGER_PROVIDERS]
    })
], AppModule);
export { AppModule };
//# sourceMappingURL=app.module.js.map