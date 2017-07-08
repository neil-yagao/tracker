var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
import { Component, Input, EventEmitter, Output } from '@angular/core';
import { URLSearchParams } from '@angular/http';
import { WorkoutTemplate } from './workout-template';
import { Exercise } from './exercise';
import { HttpBase } from '../../app/httpbase';
import * as _ from "lodash";
var PlanExercisePage = (function () {
    function PlanExercisePage(http) {
        var _this = this;
        this.http = http;
        this.toMain = new EventEmitter();
        this.activeExercise = new Exercise();
        http.get('movements').subscribe(function (movements) {
            _this.movements = movements;
        });
    }
    PlanExercisePage.prototype.addExercise = function () {
        if (this.activeExercise.name) {
            var tempExercise = new Exercise();
            //directly clone will cause type incorrect.
            //force type change
            tempExercise.name = this.activeExercise.name;
            tempExercise.rep = Number(this.activeExercise.rep);
            tempExercise.weight = Number(this.activeExercise.weight);
            tempExercise.sets = Number(this.activeExercise.sets);
            this.template.addExercise(tempExercise);
            this.activeExercise = new Exercise();
        }
    };
    PlanExercisePage.prototype.deleteExercise = function (exercise) {
        _.remove(this.template.movements, function (e) {
            return e.name == exercise.name;
        });
    };
    PlanExercisePage.prototype.createSession = function () {
        var _this = this;
        var param = new URLSearchParams();
        param.set('user', window.localStorage.getItem('username'));
        this.template.startAt = this.template.startAt.slice(0, 10);
        this.http.put('workouts', this.template, param).subscribe(function (res) {
            if (res.success) {
                _this.toMain.emit(true);
            }
        });
    };
    return PlanExercisePage;
}());
__decorate([
    Input(),
    __metadata("design:type", WorkoutTemplate)
], PlanExercisePage.prototype, "template", void 0);
__decorate([
    Output(),
    __metadata("design:type", EventEmitter)
], PlanExercisePage.prototype, "toMain", void 0);
PlanExercisePage = __decorate([
    Component({
        selector: 'plan-exercise',
        templateUrl: 'plan-exercise.html',
        providers: [HttpBase]
    }),
    __metadata("design:paramtypes", [HttpBase])
], PlanExercisePage);
export { PlanExercisePage };
//# sourceMappingURL=plan-exercise.js.map