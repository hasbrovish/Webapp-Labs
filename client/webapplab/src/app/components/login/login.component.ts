import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent implements OnInit{

  constructor(private fb: FormBuilder){}
  type: string = "password"
  eyeIcon: string = "fa-eye-slash"
  isText: boolean = false
  loginForm!:FormGroup

  ngOnInit(): void {
  
    this.loginForm = this.fb.group({
      email: ['', [Validators.email, Validators.required]],
      password: ['', Validators.required]
    })
  }

  hideShowPass(){
    this.isText = !this.isText
    this.isText ? this.eyeIcon = 'fa-eye': this.eyeIcon = 'fa-eye-slash'
    this.isText ? this.type = 'text' : this.type = 'password'
  }

  onSubmit(){
    if(this.loginForm.valid){
      console.log(this.loginForm)
    }
    else{
      this.validateAllFormFields(this.loginForm)
      alert("Your form is invalid")
    }
  }

  validateAllFormFields(formGroup: FormGroup){
    Object.keys(formGroup.controls).forEach(field => {
      const control = formGroup.get(field)
      if(control instanceof FormControl){
        control.markAsDirty({onlySelf: true})
      }else if(control instanceof FormGroup){
        this.validateAllFormFields(control)
      }
    })
  }
}
