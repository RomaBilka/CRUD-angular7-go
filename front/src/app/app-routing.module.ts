import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {AddComponent} from "./add/add.component";
import {ListComponent} from "./list/list.component";
import {EditComponent} from "./edit/edit.component";

const routes: Routes = [
  { path: 'add', component: AddComponent },
  { path: 'edit/:id', component: EditComponent },
  {path : '', component: ListComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }



//export const routing = RouterModule.forRoot(routes);