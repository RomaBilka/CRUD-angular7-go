import { Component, OnInit } from '@angular/core';
import {Router, ActivatedRoute, Params} from "@angular/router";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {first} from "rxjs/operators";
import {ApiService} from "../core/api.service";

@Component({
  selector: 'app-edit',
  templateUrl: './edit.component.html',
  styleUrls: ['./edit.component.css']
})
export class EditComponent implements OnInit {


  editForm: FormGroup;
  data: any;
  constructor(private formBuilder: FormBuilder,private router: Router, private activatedRoute: ActivatedRoute, private apiService: ApiService) { }

  ngOnInit() {
    this.activatedRoute.params.forEach((params: Params) => {
      let id = params["id"];
      
      this.editForm = this.formBuilder.group({
        id: [''],
        name: ['', Validators.required],
        description: [''],
      });
      this.apiService.getById(id)
        .subscribe( data => {
          this.editForm.setValue(data);
          this.data = data;
        });
    });
  }

  public checkError(element: string, errorType: string) {
    return this.editForm.get(element).hasError(errorType) &&
        this.editForm.get(element).touched
  }


  onSubmit() {
    this.apiService.update(this.editForm.value)
      .pipe(first())
      .subscribe(
        () => {
          this.router.navigate(['/']);
        }
      );
  }

}
