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
import { NavParams, ViewController } from 'ionic-angular';
import { Observable } from 'rxjs/Observable';
var RestingModal = (function () {
    function RestingModal(params, viewCtrl) {
        var _this = this;
        this.viewCtrl = viewCtrl;
        var source = Observable
            .interval(1000)
            .timeInterval()
            .take(600);
        this.percentage = "p100";
        this.countDown = params.get("timeout");
        this.subscription = source.subscribe(function (x) {
            _this.countDown--;
            var percentage = Math.floor(_this.countDown / params.get("timeout") * 100) + 1;
            if (percentage > 100) {
                percentage = 100;
            }
            else if (percentage <= 0) {
                percentage = 0;
            }
            _this.percentage = "p" + String(percentage);
            console.info(_this.percentage);
        }, function (error) { }, function () {
            _this.dismiss();
        });
    }
    RestingModal.prototype.dismiss = function () {
        this.viewCtrl.dismiss();
        this.subscription.unsubscribe(4);
    };
    return RestingModal;
}());
RestingModal = __decorate([
    Component({
        selector: 'resting-modal',
        templateUrl: 'resting-modal.html'
    }),
    __metadata("design:paramtypes", [NavParams, ViewController])
], RestingModal);
export { RestingModal };
//# sourceMappingURL=resting-modal.js.map