var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
import { Component, EventEmitter, Output } from '@angular/core';
import { AlertController } from 'ionic-angular';
import { WorkoutTemplate } from './workout-template';
import { HttpBase } from '../../app/httpbase';
import * as _ from "lodash";
var PlanBasicPage = (function () {
    function PlanBasicPage(http, alertCtrl) {
        var _this = this;
        this.http = http;
        this.alertCtrl = alertCtrl;
        this.template = new WorkoutTemplate();
        this.muscleGroup = "";
        this.nextStep = new EventEmitter();
        console.info(this.template.startAt);
        this.http.get('templates').subscribe(function (templates) {
            _this.templates = templates;
        });
    }
    PlanBasicPage.prototype.generateSession = function () {
        console.info(this.template);
        this.nextStep.emit(this.template);
    };
    PlanBasicPage.prototype.combinedMuscle = function (targetArray) {
        var _this = this;
        this.template.targetMuscle = "";
        _.forEach(targetArray, function (muscle) {
            _this.template.targetMuscle += muscle + ",";
        });
        this.template.targetMuscle = _.trimEnd(this.template.targetMuscle, ",");
        this.muscleGroup = targetArray;
    };
    PlanBasicPage.prototype.selectFromExistingTemplate = function () {
        var options = {};
        options.title = "已有训练计划";
        options.inputs = [];
        var currentTemplate = window.localStorage.getItem('template');
        _.forEach(this.templates, function (temp) {
            options.inputs.push({
                'type': 'radio',
                'label': temp.name,
                'value': temp.name,
                'checked': temp.name == currentTemplate
            });
        });
        var self = this;
        options.buttons = [
            {
                'text': '取消',
                'role': 'cancel'
            },
            {
                'text': '确定',
                handler: function (data) {
                    self.template.template = data;
                }
            }
        ];
        var templateSelector = this.alertCtrl.create(options);
        templateSelector.present();
    };
    return PlanBasicPage;
}());
__decorate([
    Output(),
    __metadata("design:type", EventEmitter)
], PlanBasicPage.prototype, "nextStep", void 0);
PlanBasicPage = __decorate([
    Component({
        selector: 'plan-basic',
        templateUrl: 'plan-basic.html',
        providers: [HttpBase]
    }),
    __metadata("design:paramtypes", [HttpBase, AlertController])
], PlanBasicPage);
export { PlanBasicPage };
//# sourceMappingURL=plan-basic.js.map