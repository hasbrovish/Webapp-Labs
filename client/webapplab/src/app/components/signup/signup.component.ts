import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

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
  constructor(private fb: FormBuilder){}
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
      console.log(this.signupForm.controls)
    }
    else{
      console.log(this.signupForm)
      this.validateAllFormFields(this.signupForm)
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
