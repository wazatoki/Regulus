import { SelectionModel } from '@angular/cdk/collections';
import { HttpErrorResponse } from '@angular/common/http';
import { ViewChild } from '@angular/core';
import { Component, OnInit } from '@angular/core';
import { MatPaginator } from '@angular/material';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { AlertDialogComponent } from 'src/app/layout/dialog/alert-dialog/alert-dialog.component';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';
import { TRUE } from 'src/app/services/models/enum/boolean';
import { StaffGroup } from 'src/app/services/models/group/staff-group';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { StaffGroupInputFormDialogComponent } from '../staff-group-input-form-dialog/staff-group-input-form-dialog.component';

@Component({
  selector: 'app-staff-group-master',
  templateUrl: './staff-group-master.component.html',
  styleUrls: ['./staff-group-master.component.css']
})
export class StaffGroupMasterComponent implements OnInit {

  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;

  displayedColumns: string[];
  dataSource: MatTableDataSource<StaffGroup>;
  selection: SelectionModel<StaffGroup>;
  conditionData: ConditionData;

  search(condition: ConditionData) {

    this.conditionData = condition;

    this.staffGroupService.findByCondition(condition).subscribe(
      (res: StaffGroup[] | HttpErrorResponse) => {
        if (res instanceof HttpErrorResponse === true) {

          this.dialog.open(NoticeDialogComponent, {
            data: { contents: 'エラーが発生したため処理が正常に完了しませんでした。' }
          });

        } else {

          this.dataSource.data = (res as StaffGroup[]) || [];

        }

      }
    );
  }

  onUpdateClicked(groupData: StaffGroup): void {
    const dialogRef = this.dialog.open(StaffGroupInputFormDialogComponent, {
      data: {
        groupData,
      },
    });

    dialogRef.componentInstance.submitted.subscribe(() => {
      this.search(this.conditionData);
    });
  }

  openInputForm(): void {
    const dialogRef = this.dialog.open(StaffGroupInputFormDialogComponent, {
      data: {
        groupData: null,
      },
    });

    dialogRef.componentInstance.submitted.subscribe(() => {
      this.search(this.conditionData);
    });
  }

  deleteItems(): void {
    if (this.selection.selected.length === 0) {
      const dialogRef = this.dialog.open(NoticeDialogComponent, {
        data: { contents: '削除対象が選択されていません。' }
      });
    } else {

      // 削除前の確認
      const dialogref = this.dialog.open(AlertDialogComponent, {
        data: {
          title: '確認',
          contents: '選択したデータを削除しますか？',
        }
      });

      dialogref.afterClosed().subscribe(result => {
        if (result === TRUE) {

          this.execDeleteItems();

        }
      });
    }
  }

  execDeleteItems(): void {

    const data: string[] = [];

    this.selection.selected.forEach((staffGroup: StaffGroup) => {
      data.push(staffGroup.id);
    });

    this.staffGroupService.delete(data).subscribe((res: StaffGroup[]) => {
      if (res.length > 0) {
        let str: string;
        str = '以下のデータが削除できませんでした。<br/>';

        res.forEach((s: StaffGroup) => {
          str += '・' + s.name + '<br/>';
        });

        this.dialog.open(NoticeDialogComponent, {
          data: { contents: str }
        });
      } else {
        let str: string;
        str = '選択したデータを削除しました。';

        this.dialog.open(NoticeDialogComponent, {
          data: { contents: str }
        });
      }
    });
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.dataSource.data.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.dataSource.data.forEach(row => this.selection.select(row));
  }

  /** The label for the checkbox on the passed row */
  checkboxLabel(row?: StaffGroup): string {
    if (!row) {
      return `${this.isAllSelected() ? 'select' : 'deselect'} all`;
    }
    return `${this.selection.isSelected(row) ? 'deselect' : 'select'} row ${row.id}`;
  }

  constructor(
    private staffGroupService: StaffGroupService,
    private complexSearchService: ComplexSearchService,
    public dialog: MatDialog
    ) {
    const initialSelection = [];
    const allowMultiSelect = true;
    this.displayedColumns = ['select', 'name', 'action'];
    this.dataSource = new MatTableDataSource<StaffGroup>([]);
    this.selection = new SelectionModel<StaffGroup>(allowMultiSelect, initialSelection);
    this.conditionData = this.complexSearchService.initConditionDataObj();

  }

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
  }

}
