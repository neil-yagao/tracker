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
import { NavParams, ViewController, AlertController } from 'ionic-angular';
import { HttpBase } from '../../app/httpbase';
import * as _ from "lodash";
var MovementModal = (function () {
    function MovementModal(params, viewCtrl, alertCtrl, http) {
        this.viewCtrl = viewCtrl;
        this.alertCtrl = alertCtrl;
        this.http = http;
        this.title = "修改动作";
        var title = params.get("title");
        if (title && title.length > 0) {
            this.title = title;
        }
        this.movement = params.get("movement") || {};
        if (this.movement.targetMuscle) {
            this.movement.muscleGroup = _.split(this.movement.targetMuscle, ",");
        }
        else {
            this.movement.muscleGroup = [];
        }
    }
    MovementModal.prototype.combinedMuscle = function (targetArray) {
        var _this = this;
        this.movement.targetMuscle = "";
        _.forEach(targetArray, function (muscle) {
            _this.movement.targetMuscle += muscle + ",";
        });
        this.movement.targetMuscle = _.trimEnd(this.movement.targetMuscle, ",");
        this.movement.muscleGroup = targetArray;
    };
    MovementModal.prototype.saveOrUpdateMovement = function () {
        if (this.movement.id) {
            this.updateMovement();
        }
        else {
            this.createMovement();
        }
    };
    MovementModal.prototype.createMovement = function () {
        var _this = this;
        var confirm = this.alertCtrl.create({
            title: '创建动作',
            message: '是否保存动作',
            buttons: [{
                    text: '确定',
                    handler: function () {
                        _this.http.put('movements', _this.movement)
                            .subscribe(function (respons) {
                            console.info(respons);
                            _this.viewCtrl.dismiss(true);
                        });
                    }
                }, {
                    text: '取消',
                    handler: function () {
                    }
                }]
        });
        confirm.present();
    };
    MovementModal.prototype.updateMovement = function () {
        var _this = this;
        var confirm = this.alertCtrl.create({
            title: '修改动作',
            message: '是否更改动作？',
            buttons: [
                {
                    text: '确定',
                    handler: function () {
                        var movementId = _this.movement.id;
                        if (movementId && movementId != -1) {
                            //update
                            _this.http.post('movement/' + movementId, _this.movement)
                                .subscribe(function (respons) {
                                console.info(respons);
                                _this.viewCtrl.dismiss();
                            });
                        }
                        else {
                            _this.http.put('movements', _this.movement);
                        }
                    }
                }, {
                    text: '取消',
                    handler: function () {
                    }
                }
            ]
        });
        confirm.present();
    };
    MovementModal.prototype.dividableChange = function (newValue) {
        if (newValue) {
            this.movement.dividable = 1;
        }
        else {
            this.movement.dividable = 0;
        }
    };
    MovementModal.prototype.dismiss = function () {
        this.viewCtrl.dismiss();
    };
    return MovementModal;
}());
MovementModal = __decorate([
    Component({
        selector: 'movement-modal',
        templateUrl: 'movement-modal.html',
        providers: [HttpBase]
    }),
    __metadata("design:paramtypes", [NavParams, ViewController,
        AlertController, HttpBase])
], MovementModal);
export { MovementModal };
//# sourceMappingURL=movement-modal.js.map