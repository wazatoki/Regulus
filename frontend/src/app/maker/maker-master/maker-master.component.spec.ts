import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MakerMasterComponent } from './maker-master.component';
import { LayoutModule } from '../../layout/layout.module';
import {MatTableModule} from '@angular/material/table';
import { MakerService } from '../../services/api/maker.service';
import { MakerCondition } from '../../services/models/maker/maker-condition'

describe('MakerMasterComponent', () => {
  let component: MakerMasterComponent;
  let fixture: ComponentFixture<MakerMasterComponent>;

  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('MakerService', ['findByCondition']);

    TestBed.configureTestingModule({
      declarations: [ MakerMasterComponent ],
      imports: [ 
        LayoutModule,
        MatTableModule,
       ],
       providers: [
        { provide: MakerService, useValue: spy },
        MakerCondition,
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
