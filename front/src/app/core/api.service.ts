import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {Hotdog} from "../model/hotdog.model";
import {Observable} from "rxjs/index";
import {ApiResponse} from "../model/api.response";

@Injectable()

export class ApiService {

  constructor(private http: HttpClient) { }
  baseUrl: string = 'http://localhost:3000/';


  get() : Observable<ApiResponse> {
    return this.http.get<ApiResponse>(this.baseUrl);
  }

  getById(id: number): Observable<ApiResponse> {
    return this.http.get<ApiResponse>(this.baseUrl + id);
  }

  create(hotdog: Hotdog): Observable<ApiResponse> {
    return this.http.post<ApiResponse>(this.baseUrl, hotdog);
  }

  update(hotdog: Hotdog): Observable<ApiResponse> {
    return this.http.put<ApiResponse>(this.baseUrl + hotdog.id, hotdog);
  }

  delete(id: number): Observable<ApiResponse> {
    return this.http.delete<ApiResponse>(this.baseUrl + id);
  }
}