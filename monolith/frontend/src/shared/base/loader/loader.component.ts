import { Component } from '@angular/core';

@Component({
  selector: 'app-loader',
  template: `
    <div class="fade show loader-wrapper">
      <div class="spinner-grow spinner-grow-sm" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
      <span class="m-1">Loading...</span>
    </div>
  `,
  styleUrls: ['./loader.component.scss']
})
export class LoaderComponent { }