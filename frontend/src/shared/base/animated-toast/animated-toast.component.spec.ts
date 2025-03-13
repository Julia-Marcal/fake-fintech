import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnimatedToastComponent } from './animated-toast.component';

describe('AnimatedToastComponent', () => {
  let component: AnimatedToastComponent;
  let fixture: ComponentFixture<AnimatedToastComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AnimatedToastComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AnimatedToastComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
