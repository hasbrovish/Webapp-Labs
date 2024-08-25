import { Component } from '@angular/core';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.scss'
})
export class SignupComponent {
  
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
