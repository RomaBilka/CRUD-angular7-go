import { Component, OnInit, ViewChild} from '@angular/core';
import {Router} from "@angular/router";
import {Hotdog} from "../model/hotdog.model";
import {ApiService} from "../core/api.service";
import {MatTableDataSource} from '@angular/material';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})

export class ListComponent implements OnInit {

  hotdogs: any;

  constructor(private router: Router, private apiService: ApiService) { }

  displayedColumns: string[] = ['id', 'name', 'description','edit','delete'];
  dataSource = [];
 
  private  extractData(res: any){
    this.dataSource = res;
  }
  
  ngOnInit() {
    this.apiService.get()
      .subscribe( data => {
          this.extractData(data)
      });
  }

  public deleteHotdog(id: number): void {
    this.apiService.delete(id).subscribe( () => {
      this.ngOnInit()
    });
  };


  public addHotdog(): void {
    this.router.navigate(['add']);
  };

}
