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
import { NavController, ModalController } from 'ionic-angular';
import { HttpBase } from '../../app/httpbase';
import { MovementModal } from './movement-modal';
import * as _ from "lodash";
var MovementPage = (function () {
    function MovementPage(navCtrl, http, modalCtrl) {
        this.navCtrl = navCtrl;
        this.http = http;
        this.modalCtrl = modalCtrl;
        this.criteria = "";
        this.refreshMovementFromBackend();
    }
    MovementPage.prototype.refreshMovementFromBackend = function () {
        var _this = this;
        this.http.get('movements').subscribe(function (movements) {
            _this.movements = movements;
            _this.allMovements = movements;
        });
    };
    MovementPage.prototype.filter = function (criteria) {
        this.movements = this.allMovements;
        this.movements = _.filter(this.movements, function (movement) {
            return movement.name.indexOf(this.criteria) >= 0;
        }.bind(this));
    };
    MovementPage.prototype.showModal = function (mv, title) {
        var _this = this;
        var movement = this.modalCtrl.create(MovementModal, { movement: mv, 'title': (title ? title : "") }, { showBackdrop: false, enableBackdropDismiss: false });
        movement.onDidDismiss(function (needRefresh) {
            if (needRefresh) {
                _this.refreshMovementFromBackend();
            }
        });
        movement.present();
    };
    MovementPage.prototype.getMovements = function () {
        return this.movements;
    };
    MovementPage.prototype.newMovement = function () {
        this.showModal({}, "添加动作");
    };
    return MovementPage;
}());
MovementPage = __decorate([
    Component({
        selector: 'page-movement',
        templateUrl: 'movement.html',
        providers: [HttpBase]
    }),
    __metadata("design:paramtypes", [NavController, HttpBase, ModalController])
], MovementPage);
export { MovementPage };
//# sourceMappingURL=movement.js.map