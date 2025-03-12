import { Injectable, ComponentFactoryResolver, ApplicationRef, Injector, EmbeddedViewRef } from '@angular/core';
import { LoaderComponent } from '../../base/loader/loader.component';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LoaderService {
  private loaderContainer: HTMLElement;
  private componentRef: any;
  private isLoading$ = new BehaviorSubject<boolean>(false); 

  constructor(
    private resolver: ComponentFactoryResolver,
    private appRef: ApplicationRef,
    private injector: Injector
  ) {
    this.loaderContainer = document.createElement('div');
    this.loaderContainer.className = 'loader-container';
    document.body.appendChild(this.loaderContainer);

    this.isLoading$.subscribe((isLoading) => {
      if (isLoading) {
        this.showLoader();
      } else {
        this.hideLoader();
      }
    });
  }

  setLoading(isLoading: boolean) {
    this.isLoading$.next(isLoading);
  }

  private showLoader() {
    if (this.componentRef) return;

    const factory = this.resolver.resolveComponentFactory(LoaderComponent);
    this.componentRef = factory.create(this.injector);

    this.appRef.attachView(this.componentRef.hostView);
    const domElem = (this.componentRef.hostView as EmbeddedViewRef<any>).rootNodes[0] as HTMLElement;
    this.loaderContainer.appendChild(domElem);
  }

  private hideLoader() {
    if (!this.componentRef) return; 

    this.appRef.detachView(this.componentRef.hostView);
    this.componentRef.destroy();
    this.componentRef = null;
  }
}