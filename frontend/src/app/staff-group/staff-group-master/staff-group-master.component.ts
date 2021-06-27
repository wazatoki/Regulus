import { SelectionModel } from '@angular/cdk/collections';
import { ViewChild } from '@angular/core';
import { Component, OnInit } from '@angular/core';
import { MatPaginator } from '@angular/material';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';
import { StaffGroup } from 'src/app/services/models/group/staff-group';
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

  onUpdateClicked(groupData: StaffGroup): void {
    this.dialog.open(StaffGroupInputFormDialogComponent, {
      data: {
        groupData: groupData,
      },
    })
  }

  openInputForm(): void {
    this.dialog.open(StaffGroupInputFormDialogComponent, {
      data: {
        groupData: null,
      },
    })
  }

  deleteItems(): void {
    if (this.selection.selected.length === 0) {
      const dialogRef = this.dialog.open(NoticeDialogComponent, {
        data: { contents: '削除対象が選択されていません。' }
      });
    } else {
      const data: string[] = [];
      this.selection.selected.forEach((staffGroup: StaffGroup) => {
        data.push(staffGroup.id);
      });
      this.staffGroupService.delete(data).subscribe((res: StaffGroup[]) => {
        if (res.length > 0) {
          let str: string;
          str = '以下のdataが削除できませんでした。<br/>';

          res.forEach((s: StaffGroup) => {
            str += '・' + s.name + '<br/>';
          });

          this.dialog.open(NoticeDialogComponent, {
            data: { contents: str }
          });
        }
      });
    }
  }

  onFetchedSearchConditions(data: StaffGroup[]) {
    this.dataSource.data = data || []
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

  constructor(private staffGroupService: StaffGroupService, public dialog: MatDialog
  ) {
    const initialSelection = [];
    const allowMultiSelect = true;
    this.displayedColumns = ['select', 'name'];
    this.dataSource = new MatTableDataSource<StaffGroup>([]);
    this.selection = new SelectionModel<StaffGroup>(allowMultiSelect, initialSelection);

  }

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
  }

}
