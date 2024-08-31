import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import ValidateForm from '../../helpers/validateForm';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent implements OnInit{

  constructor(private fb: FormBuilder, private auth: AuthService, private router: Router){}
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
      this.auth.login(this.loginForm.value).subscribe({
        next:(res=>{
          //alert(res.message)
          this.loginForm.reset()
          this.router.navigate(['dashboard'])
        }),
        error:(err=>{
          alert(err?.error.message)
        })
      })
    }
    else{
      ValidateForm.validateAllFormFields(this.loginForm)
      alert("Your form is invalid")
    }
  }
}
