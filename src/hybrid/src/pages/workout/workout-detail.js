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
import { URLSearchParams } from '@angular/http';
import { NavParams, Navbar, ModalController } from 'ionic-angular';
import * as _ from "lodash";
import { Md5 } from 'ts-md5/dist/md5';
import { HttpBase } from '../../app/httpbase';
import { RestingModal } from './resting-modal';
var WorkoutDetail = (function () {
    function WorkoutDetail(params, http, modalCtrl) {
        var _this = this;
        this.params = params;
        this.http = http;
        this.modalCtrl = modalCtrl;
        var workoutId = params.data.workout.id;
        var param = new URLSearchParams();
        param.set('user', String(Md5.hashStr('powerlift')));
        http.get('session/' + workoutId, param).subscribe(function (session) {
            var movements = _.groupBy(session, 'movementName');
            var beforeAssign = _.map(movements, function (value, key) {
                var sequence = 0;
                var newSets = _.map(value, function (set) {
                    var eachSide = (set.targetWeight - 20) / 2;
                    if (!set.dividable) {
                        eachSide = 1 / 0;
                    }
                    if (set.sequence != 0) {
                        sequence += 1;
                    }
                    return {
                        'id': set.id,
                        'sequence': set.sequence == 0 ? "*" : sequence,
                        'repeat': set.targetNumber,
                        'eachSide': eachSide,
                        'totalWeight': set.targetWeight,
                        "eachSideShow": false,
                        'achieved': -1
                    };
                });
                return {
                    "name": key,
                    "sets": newSets
                };
            });
            console.info(beforeAssign);
            _this.exercises = beforeAssign;
        });
    }
    WorkoutDetail.prototype.finishOneSet = function (set) {
        var _this = this;
        if (set.achieved == undefined || set.achieved <= 0) {
            set.achieved = set.repeat;
        }
        else {
            set.achieved -= 1;
        }
        var self = this;
        if (self.debounce) {
            self.debounce.cancel();
        }
        self.debounce = _.debounce(function () {
            var resting = _this.modalCtrl.create(RestingModal, { timeout: 60 }, { showBackdrop: false, enableBackdropDismiss: false });
            resting.present();
        }, 1000);
        self.debounce();
    };
    WorkoutDetail.prototype.finishSession = function () {
        var sessionData = _.reduce(this.exercises, function (arry, value) {
            return _.concat(arry, _.flatMap(value.sets, function (set) {
                return {
                    'id': set.id,
                    'acheiveNumber': set.achieved,
                    'acheiveWeight': set.totalWeight + 0
                };
            }));
        }, []);
        console.info(sessionData);
        this.http.post('session/' + this.params.data.workout.id, JSON.stringify(sessionData))
            .subscribe(function () { });
    };
    WorkoutDetail.prototype.ngAfterViewInit = function () {
        this.navBar.setBackButtonText("返回列表");
    };
    return WorkoutDetail;
}());
__decorate([
    ViewChild(Navbar),
    __metadata("design:type", Navbar)
], WorkoutDetail.prototype, "navBar", void 0);
WorkoutDetail = __decorate([
    Component({
        selector: 'workout-detail',
        templateUrl: 'workout-detail.html',
        providers: [HttpBase]
    }),
    __metadata("design:paramtypes", [NavParams, HttpBase, ModalController])
], WorkoutDetail);
export { WorkoutDetail };
//# sourceMappingURL=workout-detail.js.map