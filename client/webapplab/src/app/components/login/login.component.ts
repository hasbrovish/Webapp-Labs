import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent implements OnInit{

  type: string = "password"
  eyeIcon: string = "fa-eye-slash"
  isText: boolean = false
  ngOnInit(): void {
      
  }

  hideShowPass(){
    this.isText = !this.isText
    this.isText ? this.eyeIcon = 'fa-eye': this.eyeIcon = 'fa-eye-slash'
    this.isText ? this.type = 'text' : this.type = 'password'
  }

}
