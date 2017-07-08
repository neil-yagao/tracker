var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
import { Component, ViewChild } from '@angular/core';
import { NavController } from 'ionic-angular';
import { HttpBase } from '../../app/httpbase';
import { WorkoutTemplate } from './workout-template';
import { Exercise } from './exercise';
var PlanPage = (function () {
    function PlanPage(navCtrl, http) {
        var _this = this;
        this.navCtrl = navCtrl;
        this.http = http;
        http.get('movements').subscribe(function (movements) {
            _this.movements = movements;
        });
        this.activeExercise = new Exercise();
        this.workoutTemplate = new WorkoutTemplate();
    }
    PlanPage.prototype.goNextSlide = function (template) {
        this.workoutTemplate = template;
        this.planSlide.slideTo(1);
    };
    PlanPage.prototype.toMain = function () {
        this.planSlide.slideTo(0);
        this.workoutTemplate = new WorkoutTemplate();
        var t = this.navCtrl.parent;
        t.select(0);
    };
    return PlanPage;
}());
__decorate([
    ViewChild('planSlide'),
    __metadata("design:type", Object)
], PlanPage.prototype, "planSlide", void 0);
PlanPage = __decorate([
    Component({
        selector: 'page-plan',
        templateUrl: 'plan.html',
        providers: [HttpBase]
    }),
    __metadata("design:paramtypes", [NavController, HttpBase])
], PlanPage);
export { PlanPage };
//# sourceMappingURL=plan.js.map