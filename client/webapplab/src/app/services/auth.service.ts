import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private baseUrl : string = "localhost:9090/"
  constructor(private httpclient : HttpClient) { }

  signup(userObj: any){
    return this.httpclient.post<any>(`${this.baseUrl}register`, userObj);
  }

  login(loginObj : any){
    return this.httpclient.post<any>(`${this.baseUrl}login`, loginObj)
  }

}
