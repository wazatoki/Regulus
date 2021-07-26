import { ViewChild, Component, OnInit } from '@angular/core';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { MatDialog } from '@angular/material/dialog';
import { SaveData } from 'src/app/services/models/search/save-data';
import { SelectionModel } from '@angular/cdk/collections';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { ComplexSearchConditionInputFormDialogComponent } from 'src/app/complex-search-condition/complex-search-condition-input-form-dialog/complex-search-condition-input-form-dialog.component';
import { AlertDialogComponent } from 'src/app/layout/dialog/alert-dialog/alert-dialog.component';
import { TRUE } from 'src/app/services/models/enum/boolean';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-complex-search-condition-master',
  templateUrl: './complex-search-condition-master.component.html',
  styleUrls: ['./complex-search-condition-master.component.css']
})
export class ComplexSearchConditionMasterComponent implements OnInit {

  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;

  displayedColumns: string[];
  dataSource: MatTableDataSource<SaveData>;
  selection: SelectionModel<SaveData>;
  conditionData: ConditionData;

  search(condition: ConditionData) {


    this.complexSearchConditionService.findByCondition(condition).subscribe(
      (res: SaveData[] | HttpErrorResponse) => {

        if (res instanceof HttpErrorResponse == true) {

          this.dialog.open(NoticeDialogComponent, {
            data: { contents: 'エラーが発生したため処理が正常に完了しませんでした。' }
          });

        } else {

          this.dataSource.data = (res as SaveData[]) || []
          this.conditionData = condition

        }

      }
    );
  }

  onUpdateClicked(saveData: SaveData): void {

    this.complexSearchConditionService.findAllCategories().subscribe(categories => {
      const dialogRef = this.dialog.open(ComplexSearchConditionInputFormDialogComponent, {
        data: {
          categories: categories,
          saveData: saveData,
        },
      })

      dialogRef.componentInstance.submitted.subscribe(() => {
        this.search(this.conditionData)
      })
    });

  }

  openInputForm(): void {

    this.complexSearchConditionService.findAllCategories().subscribe(categories => {
      const dialogRef = this.dialog.open(ComplexSearchConditionInputFormDialogComponent, {
        data: {
          categories: categories,
          saveData: null,
        },
      })

      dialogRef.componentInstance.submitted.subscribe(() => {
        this.search(this.conditionData)
      })
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

    this.selection.selected.forEach((saveData: SaveData) => {
      data.push(saveData.id);
    });

    this.complexSearchConditionService.delete(data).subscribe((res: SaveData[]) => {
      if (res.length > 0) {
        let str: string;
        str = '以下のデータが削除できませんでした。<br/>';

        res.forEach((s: SaveData) => {
          str += '・' + s.patternName + '<br/>';
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

  onFetchedSearchConditions(data: SaveData[]) {
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
  checkboxLabel(row?: SaveData): string {
    if (!row) {
      return `${this.isAllSelected() ? 'select' : 'deselect'} all`;
    }
    return `${this.selection.isSelected(row) ? 'deselect' : 'select'} row ${row.id}`;
  }

  constructor(
    private complexSearchConditionService: ComplexSearchConditionService,
    private complexSearchService:ComplexSearchService,
    public dialog: MatDialog
  ) {
    const initialSelection = [];
    const allowMultiSelect = true;
    this.displayedColumns = ['select', 'name', 'category', 'owner', 'action'];
    this.dataSource = new MatTableDataSource<SaveData>([]);
    this.selection = new SelectionModel<SaveData>(allowMultiSelect, initialSelection);
    this.conditionData = this.complexSearchService.initConditionDataObj();

  }

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
  }

}
