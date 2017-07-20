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
import { URLSearchParams } from '@angular/http';
import { NavController, AlertController } from 'ionic-angular';
import { TabsPage } from '../pages/tabs/tabs';
import { Md5 } from 'ts-md5/dist/md5';
import { HttpBase } from './httpbase';
var LoginPage = (function () {
    function LoginPage(nav, http, alert) {
        this.nav = nav;
        this.http = http;
        this.alert = alert;
        this.nav = nav;
    }
    LoginPage.prototype.login = function () {
        var _this = this;
        var usr = String(Md5.hashStr(this.username));
        window.localStorage.setItem('username', usr);
        window.localStorage.setItem('template', this.template);
        var param = new URLSearchParams();
        param.set('user', usr);
        this.http.post('login', {
            'username': this.username,
            'userIdentity': usr,
            'template': this.template
        }, param).subscribe(function (response) {
            if (response.success) {
                console.debug("login!" + window.localStorage.getItem('username'));
                _this.nav.setRoot(TabsPage);
            }
            else {
                _this.alert.create({
                    'title': '错误',
                    'subTitle': response.reason
                }).present();
            }
        });
    };
    return LoginPage;
}());
LoginPage = __decorate([
    Component({
        templateUrl: 'login.html',
        providers: [HttpBase]
    }),
    __metadata("design:paramtypes", [NavController, HttpBase,
        AlertController])
], LoginPage);
export { LoginPage };
//# sourceMappingURL=login.js.map