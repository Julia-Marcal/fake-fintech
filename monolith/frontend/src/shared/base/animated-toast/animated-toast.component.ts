import { Component, Input, Output, signal, EventEmitter } from '@angular/core';
import {
  ButtonDirective,
  ProgressComponent,
  ToastBodyComponent,
  ToasterComponent,
  ToastComponent,
  ToastHeaderComponent
} from '@coreui/angular';

@Component({
  selector: 'app-toast',
  standalone: true,
  templateUrl: './animated-toast.component.html',
  styleUrls: ['./animated-toast.component.scss'],
  imports: [
    ButtonDirective,
    ProgressComponent,
    ToasterComponent,
    ToastComponent,
    ToastHeaderComponent,
    ToastBodyComponent
  ],
  exportAs: 'appToast'
})
export class AnimatedToastComponent {
  @Input() title: string = '';
  @Input() message: string = '';
  @Input() toastClass: string = '';
  @Output() dismissed = new EventEmitter<void>();

  position = 'top-end';
  visible = signal(false);
  percentage = signal(0);

  toggleToast() {
    this.visible.update((value) => !value);
    if (this.visible()) {
      this.percentage.set(100);
    } else {
      this.dismissed.emit();
    }
  }

  onVisibleChange(isVisible: boolean) {
    this.visible.set(isVisible);
    if (!isVisible) {
      this.percentage.set(0);
      this.dismissed.emit();
    }
  }

  onTimerChange(progress: number) {
    this.percentage.set(progress * 25);
  }
}