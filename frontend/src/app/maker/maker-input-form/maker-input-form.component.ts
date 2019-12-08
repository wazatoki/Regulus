import { Component, OnInit, Inject } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Maker } from '../../services/models/maker/maker';
import { NoticeDialogComponent } from '../../layout/dialog/notice-dialog/notice-dialog.component';
import { FormBuilder, FormGroup, AbstractControl, Validators } from '@angular/forms';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MakerService } from '../../services/api/maker.service';

export interface DialogData {
  name: string;
}

@Component({
  selector: 'app-maker-input-form',
  templateUrl: './maker-input-form.component.html',
  styleUrls: ['./maker-input-form.component.css']
})
export class MakerInputFormComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    public makerService: MakerService,
    private dialog: MatDialog,
    public dialogRef: MatDialogRef<MakerInputFormComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData) { }

  ngOnInit() {
  }

  makerForm: FormGroup = this.fb.group({
    name: ['', [Validators.required]],
  });

  get name(): AbstractControl {
    return this.makerForm.get('name');
  }

  onCancelClick() {
    this.dialogRef.close();
  }

  onClearClick(){
    this.makerForm.reset();
  }

  onSaveClick(){
    this.makerService.add(this.makerForm.value).subscribe(
      (res: Maker) => {
        const dialogRef = this.dialog.open(NoticeDialogComponent, {
          width: '250px',
          data: { contents: '製造販売業者情報を保存しました。' }
        });
      },
      (error: HttpErrorResponse) => {},
    );
  }
}
