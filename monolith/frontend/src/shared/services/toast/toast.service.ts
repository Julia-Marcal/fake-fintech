import { Injectable, ComponentFactoryResolver, ApplicationRef, Injector, EmbeddedViewRef } from '@angular/core';
import { AnimatedToastComponent } from '../../base/animated-toast/animated-toast.component';

@Injectable({
  providedIn: 'root'
})
export class ToastService {
  private toastContainer: HTMLElement;

  constructor(
    private resolver: ComponentFactoryResolver,
    private appRef: ApplicationRef,
    private injector: Injector
  ) {
    this.toastContainer = document.createElement('div');
    this.toastContainer.className = 'toast-container';
    document.body.appendChild(this.toastContainer);
  }

  showToast(options: { title: string; message: string; duration?: number; position?: string }) {
    const factory = this.resolver.resolveComponentFactory(AnimatedToastComponent);
    const componentRef = factory.create(this.injector);
    const instance = componentRef.instance;

    instance.title = options.title;
    instance.message = options.message;
    instance.position = options.position || 'top-end'; 

    instance.dismissed.subscribe(() => {
      this.removeToast(componentRef);
    });

    this.appRef.attachView(componentRef.hostView);
    const domElem = (componentRef.hostView as EmbeddedViewRef<any>).rootNodes[0] as HTMLElement;
    this.toastContainer.appendChild(domElem);

    if (options.duration) {
      setTimeout(() => {
        instance.toggleToast(); 
      }, options.duration);
    }

    instance.toggleToast();
  }

  private removeToast(componentRef: any) {
    this.appRef.detachView(componentRef.hostView);
    componentRef.destroy();
  }
}