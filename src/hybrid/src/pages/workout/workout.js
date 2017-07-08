var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
import { Component } from '@angular/core';
import { NavController, PopoverController } from 'ionic-angular';
import { HttpBase } from '../../app/httpbase';
import { URLSearchParams } from '@angular/http';
import { WorkoutDetail } from './workout-detail';
import { TemplatePopover } from './template-popover';
import * as _ from 'lodash';
var WorkoutPage = (function () {
    function WorkoutPage(navCtrl, pop, httpBase) {
        this.navCtrl = navCtrl;
        this.pop = pop;
        this.httpBase = httpBase;
        this.debounce = false;
    }
    ;
    WorkoutPage.prototype.goToDetail = function (workout) {
        this.navCtrl.push(WorkoutDetail, { workout: workout });
    };
    ;
    WorkoutPage.prototype.ionViewWillEnter = function () {
        var _this = this;
        var param = new URLSearchParams();
        param.set('user', window.localStorage.getItem('username'));
        param.set('template', window.localStorage.getItem('template'));
        this.httpBase.get('workouts', param).subscribe(function (workouts) { return _this.workouts = _.sortBy(workouts, "perform_date"); });
    };
    WorkoutPage.prototype.getNoteColor = function (workout) {
        if (workout.perform_date < new Date().toISOString()) {
            return { 'color': 'red' };
        }
        else {
            return {};
        }
    };
    WorkoutPage.prototype.openPopover = function (event) {
        var popover = this.pop.create(TemplatePopover);
        popover.present({
            ev: event
        });
    };
    return WorkoutPage;
}());
WorkoutPage = __decorate([
    Component({
        selector: 'page-workout',
        templateUrl: 'workout.html',
        providers: [HttpBase]
    }),
    __metadata("design:paramtypes", [NavController, PopoverController,
        HttpBase])
], WorkoutPage);
export { WorkoutPage };
//# sourceMappingURL=workout.js.map