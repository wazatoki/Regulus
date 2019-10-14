import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MakerMasterComponent } from './maker-master.component';
import { LayoutModule } from '../../layout/layout.module';
import {MatTableModule} from '@angular/material/table';

describe('MakerMasterComponent', () => {
  let component: MakerMasterComponent;
  let fixture: ComponentFixture<MakerMasterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MakerMasterComponent ],
      imports: [ 
        LayoutModule,
        MatTableModule,
       ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakerMasterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
