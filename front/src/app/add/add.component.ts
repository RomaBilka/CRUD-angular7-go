import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {Router} from "@angular/router";
import {ApiService} from "../core/api.service";

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.css']
})
export class AddComponent implements OnInit {

  constructor(private formBuilder: FormBuilder,private router: Router, private apiService: ApiService) { }

  addForm: FormGroup;

  ngOnInit() {
    this.addForm = this.formBuilder.group({
      name: ['', Validators.required],
      description: ['']
    });

  }

  onSubmit() {
    this.apiService.create(this.addForm.value)
      .subscribe( () => {
        this.router.navigate(['/']);
    });
  }
  public checkError(element: string, errorType: string) {
    return this.addForm.get(element).hasError(errorType) &&
        this.addForm.get(element).touched
  }

}
