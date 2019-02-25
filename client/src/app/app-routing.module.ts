import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BucketsComponent } from './components/buckets/buckets.component'
import { DataContentComponent } from './components/data-content/data-content.component'

const routes: Routes = [
  { path: '', component: BucketsComponent },
  { path: 'data/:bucket', component: DataContentComponent }
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}