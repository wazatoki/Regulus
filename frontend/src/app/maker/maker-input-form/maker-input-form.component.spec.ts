import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MakerInputFormComponent } from './maker-input-form.component';
import { LayoutModule } from '../../layout/layout.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';

describe('MakerInputFormComponent', () => {
  let component: MakerInputFormComponent;
  let fixture: ComponentFixture<MakerInputFormComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MakerInputFormComponent ],
      imports: [
        BrowserAnimationsModule,
        ReactiveFormsModule,
        MatFormFieldModule,
        MatInputModule,
        MatDialogModule,
        MatGridListModule,
        LayoutModule,
      ],
      providers: [
        {
          provide: MatDialogRef,
          useValue: { close: (dialogResult: any) => { } }
        },
        {
          provide: MAT_DIALOG_DATA, useValue: {} 
        },
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakerInputFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
