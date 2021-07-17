import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatDialog, MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';

import { StaffGroupInputFormDialogComponent } from './staff-group-input-form-dialog.component';
import { createTestInstance1 as createGroupData} from 'src/app/services/models/group/staff-group.spec'
import { StaffGroup } from 'src/app/services/models/group/staff-group';
import { LayoutModule } from 'src/app/layout/layout.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule, MatCardModule, MatFormFieldModule, MatGridListModule, MatInputModule } from '@angular/material';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { BrowserDynamicTestingModule } from '@angular/platform-browser-dynamic/testing';
import { Subject } from 'rxjs';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';

describe('StaffGroupInputFormDialogComponent', () => {
  let component: StaffGroupInputFormDialogComponent;
  let fixture: ComponentFixture<StaffGroupInputFormDialogComponent>;
  const matDialogRefSpy = jasmine.createSpyObj('MatDialogRef', ['close', 'updateSize']);
  let staffGroupServiceSpy: jasmine.SpyObj<StaffGroupService>;
  let dialog: MatDialog;
  const groupData: StaffGroup = createGroupData();
  const dialogPassedData = {
    groupData: groupData,
  };

  beforeEach(async(() => {
    staffGroupServiceSpy = jasmine.createSpyObj('StaffGroupService', ['update', 'add']);

    TestBed.configureTestingModule({
      declarations: [ StaffGroupInputFormDialogComponent ],
      imports: [
        LayoutModule,
        FlexLayoutModule,
        FormsModule,
        ReactiveFormsModule,
        MatDialogModule,
        MatCardModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        MatGridListModule,
        NoopAnimationsModule,
      ],
      providers: [
        {
          provide: FormBuilder,
          useValue: new FormBuilder()
        },
        {
          provide: StaffGroupService,
          useValue: staffGroupServiceSpy
        },
        {
          provide: MatDialogRef,
          useValue: matDialogRefSpy,
        },
        {
          provide: MAT_DIALOG_DATA,
          useValue: dialogPassedData,
        },
      ],
    })
    .overrideModule(BrowserDynamicTestingModule, {
      set: {
        entryComponents: [StaffGroupInputFormDialogComponent],
      }
    })
    .compileComponents();
  }));

  beforeEach(() => {

    dialog = TestBed.get(MatDialog);
    fixture = TestBed.createComponent(StaffGroupInputFormDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    staffGroupServiceSpy = TestBed.get(StaffGroupService)
    staffGroupServiceSpy.add.and.returnValue(new Subject<StaffGroup>().asObservable());
    staffGroupServiceSpy.update.and.returnValue(new Subject<StaffGroup>().asObservable());
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should click save button', () => {
   

    (component.staffGroup as FormGroup).get('name').setValue('value1');

    const formDebugElement: DebugElement = fixture.debugElement.query(By.css('form'));
    formDebugElement.triggerEventHandler('submit', null);
    fixture.detectChanges();

    if (component.groupData.id) {
      expect(staffGroupServiceSpy.update).toHaveBeenCalled();
    } else {
      expect(staffGroupServiceSpy.add).toHaveBeenCalled();
    }
  });

  it('should create save data', async () => {

    const nameDe: DebugElement = fixture.debugElement.query(By.css("input.name"));
    const nameEl: HTMLInputElement = nameDe.nativeElement;
    nameEl.value = 'sample name';
    nameEl.dispatchEvent(new Event('input'));
    
    component.createSaveData();
    const expectData = component.groupData;
    expect(expectData.name).toBe('sample name');
  });

  it('should clear form', async () => {

    const nameDe: DebugElement = fixture.debugElement.query(By.css("input.name"));
    const nameEl: HTMLInputElement = nameDe.nativeElement;
    nameEl.value = 'sample name';
    nameEl.dispatchEvent(new Event('input'));

    component.onClearClick();
    fixture.detectChanges();
    
    expect(nameEl.textContent).not.toContain('sample name');
    console.log(nameEl.textContent)
    expect(nameEl.textContent).toEqual('')
  });

});
