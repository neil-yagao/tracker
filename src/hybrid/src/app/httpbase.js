var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { Logger } from "angular2-logger/core";
import { Observable } from 'rxjs/Observable';
import 'rxjs/Rx';
var HttpBase = (function () {
    function HttpBase(http, _logger) {
        this.http = http;
        this._logger = _logger;
        this.baseUrl = "http://localhost:8080/";
        this.headers = new Headers({ 'Content-Type': 'application/json', 'charset': 'UTF-8' });
        this.options = new RequestOptions({ headers: this.headers });
    }
    HttpBase.prototype.get = function (url, param) {
        this._logger.debug("get url " + url + " with param " + param);
        if (!param) {
            return this.http.get(this.baseUrl + url)
                .map(this.extractData.bind(this))
                .catch(this.handleError);
        }
        else {
            return this.http.get(this.baseUrl + url, { search: param })
                .map(this.extractData.bind(this))
                .catch(this.handleError);
        }
    };
    HttpBase.prototype.post = function (url, requestBody, param) {
        if (param) {
            this.options.search = param;
        }
        this._logger.debug("post to url " + url + " with request " + JSON.stringify(requestBody));
        return this.http.post(this.baseUrl + url, JSON.stringify(requestBody), this.options)
            .map(this.extractData.bind(this))
            .catch(this.handleError);
    };
    HttpBase.prototype.put = function (url, requestBody, param) {
        if (param) {
            this.options.search = param;
        }
        this._logger.debug("post to url " + url + " with request " + JSON.stringify(requestBody));
        return this.http.put(this.baseUrl + url, JSON.stringify(requestBody), this.options)
            .map(this.extractData.bind(this))
            .catch(this.handleError);
    };
    HttpBase.prototype.extractData = function (res) {
        var body = res.json();
        this._logger.debug("response:", body);
        return body.data || body || {};
    };
    HttpBase.prototype.handleError = function (error) {
        var errMsg;
        if (error instanceof Response) {
            var body = error.json() || '';
            var err = body.error || JSON.stringify(body);
            errMsg = error.status + " - " + (error.statusText || '') + " " + err;
        }
        else {
            errMsg = error.message ? error.message : error.toString();
        }
        console.error(errMsg);
        return Observable.throw(errMsg);
    };
    return HttpBase;
}());
HttpBase = __decorate([
    Injectable(),
    __metadata("design:paramtypes", [Http, Logger])
], HttpBase);
export { HttpBase };
//# sourceMappingURL=httpbase.js.map