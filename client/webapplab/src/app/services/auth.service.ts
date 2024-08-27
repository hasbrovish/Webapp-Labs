import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private baseUrl : string = "http://localhost:9090/"
  constructor(private httpclient : HttpClient) { }

  signup(userObj: any){
    console.log(userObj)
    return this.httpclient.post<any>(`${this.baseUrl}register`, userObj);
  }

  login(loginObj : any){
    console.log(loginObj)
    return this.httpclient.post<any>(`${this.baseUrl}login`, loginObj)
  }

}
