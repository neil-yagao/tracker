import { Injectable }     from '@angular/core';
import { Http, Response, Headers, RequestOptions, URLSearchParams } from '@angular/http';

import { Logger } from "angular2-logger/core"

import { Observable }     from 'rxjs/Observable';
import 'rxjs/Rx';

@Injectable()
export class HttpBase {
	private baseUrl = "/";
	private headers = new Headers({ 'Content-Type': 'application/json' });
	private options = new RequestOptions({ headers: this.headers });
	constructor(private http: Http, private _logger: Logger) {
	}


	get(url: string, param?: URLSearchParams): Observable<any> {
		this._logger.debug("get url " + url + " with param " + param)
		if (!param) {
			return this.http.get(this.baseUrl + url)
				.map(this.extractData.bind(this))
				.catch(this.handleError);
		} else {
			return this.http.get(this.baseUrl + url, { search: param })
				.map(this.extractData.bind(this))
				.catch(this.handleError);
		}
	}

	post(url: string, requestBody: Object): Observable<any> {
		return this.http.post(this.baseUrl + url, requestBody, this.options)
			.map(this.extractData.bind(this))
			.catch(this.handleError);
	}

	put(url: string, requestBody: Object): Observable<any> {
		return this.http.put(this.baseUrl + url, requestBody, this.options)
			.map(this.extractData.bind(this))
			.catch(this.handleError);
	}

	private extractData(res: Response) {
		let body = res.json();
		this._logger.debug("response:", body)
		return body.data || {};
	}

	private handleError(error: Response | any): Observable<any> {
		let errMsg: string;

		if (error instanceof Response) {
			const body = error.json() || '';
			const err = body.error || JSON.stringify(body);
			errMsg = `${error.status} - ${error.statusText || ''} ${err}`;
		} else {
			errMsg = error.message ? error.message : error.toString();
		}
		console.error(errMsg);
		return Observable.throw(errMsg);
	}
}
