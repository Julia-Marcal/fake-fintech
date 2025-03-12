import { Component } from '@angular/core';

@Component({
  selector: 'app-loader',
  template: `<c-spinner color="primary" variant="grow"></c-spinner>`,
  styleUrls: ['./loader.component.scss']
})
export class LoaderComponent {}