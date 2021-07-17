import { Inject } from '@angular/core';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';
import { StaffGroup } from 'src/app/services/models/group/staff-group';

@Component({
  selector: 'app-staff-group-input-form-dialog',
  templateUrl: './staff-group-input-form-dialog.component.html',
  styleUrls: ['./staff-group-input-form-dialog.component.css']
})
export class StaffGroupInputFormDialogComponent implements OnInit {

  form: FormGroup;
  
  get groupData(): StaffGroup {
    return this.data.groupData;
  }

  get staffGroup() {
    return this.form.get('staffGroup') as FormGroup;
  }

  setSavedDataToForm() {

    this.staffGroup.get('name').setValue(this.groupData.name)
     
  }

  onClearClick(): void {
    this.form = this.initForm()
  }

  onCancelClick(): void {
    this.dialogRef.close();
  }

  createSaveData(): void {

    if (this.data.groupData === null || this.data.groupData === undefined) {
      this.data.groupData = {
        id: "",
        name: "",
      }
    }

    this.data.groupData.name = this.staffGroup.get("name").value

  }

  onSubmit() {

    if (this.form.valid) {

      this.createSaveData();

      if (this.groupData.id) {
        this.staffGroupService.update(this.groupData).subscribe(data => {
          this.dialog.open(NoticeDialogComponent, {
            data: { contents: '検索条件を修正しました。' }
          });
        });


      } else {

        this.staffGroupService.add(this.groupData).subscribe(data => {
          this.dialog.open(NoticeDialogComponent, {
            data: { contents: '検索条件を保存しました。' }
          });
        });

      }

    }

  }

  getNameErrorMessage() {
    return this.staffGroup.get('name').hasError('required') ? 'グループ名称は必須項目です。' : '';
  }

  initForm(): FormGroup {
    return this.fb.group({
      staffGroup: this.fb.group({
        name: this.fb.control('', [Validators.required]),
      }),
    });
  }

  constructor(
    private dialogRef: MatDialogRef<StaffGroupInputFormDialogComponent>,
    @Inject(MAT_DIALOG_DATA) private data: DialogData,
    private fb: FormBuilder,
    private dialog: MatDialog,
    private staffGroupService: StaffGroupService,
  ) {
    this.form = this.initForm();
  }

  ngOnInit() {
    this.dialogRef.updateSize("1100px")
    // groupDataの編集のときは値をフォームに反映する
    if (this.data.groupData !== null && this.data.groupData !== undefined && this.data.groupData.id !== '') {
      this.setSavedDataToForm()
    }
  }

}

export interface DialogData {
  groupData: StaffGroup;
}
