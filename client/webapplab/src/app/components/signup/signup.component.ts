import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import ValidateForm from '../../helpers/validateForm';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.scss'
})
export class SignupComponent {

  type: string = "password"
  eyeIcon: string = "fa-eye-slash"
  isText: boolean = false
  signupForm!: FormGroup
  constructor(private fb: FormBuilder, private auth : AuthService){}
  ngOnInit(): void {
      this.signupForm = this.fb.group({
        firstname : ['', Validators.required],
        lastname : ['', Validators.required],
        gender : ['', Validators.required],
        phone: ['', [Validators.required, Validators.minLength(10), Validators.pattern('[0-9]{10}')]],
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
    if(this.signupForm.valid){
      console.log(this.signupForm.value)
      this.auth.signup(this.signupForm.value).subscribe({
        next:(res=>{
          alert(res.message)
        }),
        error:(err=>{
          alert(err?.error.message)
        })
      })

    }
    else{
      console.log(this.signupForm)
      ValidateForm.validateAllFormFields(this.signupForm)
      alert("Your form is invalid")
    }
  }


}
