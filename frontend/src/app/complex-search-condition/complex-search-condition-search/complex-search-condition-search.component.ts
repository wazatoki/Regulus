import { Component, OnInit, Output, EventEmitter, Input } from '@angular/core';
import { Subscription } from 'rxjs';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';

import { ComplexSearchDialogComponent } from '../../layout/dialog/complex-search-dialog/complex-search-dialog/complex-search-dialog.component';
import { SaveData } from 'src/app/services/models/search/save-data';
import { ConditionData, mapCondition, splitStrings } from 'src/app/services/models/search/condition-data';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { ComplexSearchItems } from '../../services/models/search/complex-search-items';
import { HttpErrorResponse } from '@angular/common/http';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';

@Component({
  selector: 'app-complex-search-condition-search',
  templateUrl: './complex-search-condition-search.component.html',
  styleUrls: ['./complex-search-condition-search.component.css']
})
export class ComplexSearchConditionSearchComponent implements OnInit {

  private complexSearchSubscription: Subscription;
  private saveData: SaveData;
  private dialogRef: MatDialogRef<ComplexSearchDialogComponent>;

  @Input() condition: ConditionData;

  @Output() searchClicked: EventEmitter<ConditionData> = new EventEmitter();

  get complexConditionButtonColor():string {

    if (this.saveData.conditionData.searchConditionList.length > 0 ||
      this.saveData.conditionData.displayItemList.length > 0 ||
      this.saveData.conditionData.orderConditionList.length > 0) {

      return "accent";

    } else {

      return "";

    }
  }

  openComplexSearch() {
    this.complexSearchConditionService.findComplexSearchItems().subscribe((data: ComplexSearchItems) => {
      const aData: any = data;
      aData.saveData = this.saveData;
      // aData.saveData = this.testData(); // 検証用
      this.dialogRef = this.dialog.open(ComplexSearchDialogComponent, {
        data: aData,
      });
    });
  }

  onSearch(searchStrings: string) {

    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = splitStrings(searchStrings)
    this.searchClicked.emit(this.condition)
  }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.complexSearchSubscription.unsubscribe();
  }

  defineDialogSearch() {
    // 検索条件設定ダイアログで検索ボタンをクリックしたら検索条件を変更して検索を実行する。
    this.complexSearchSubscription = this.complexSearchService.complexSearchOrdered$.subscribe(
      (data: ConditionData) => {
        this.saveData.conditionData = data
        mapCondition(data, this.condition);
        this.searchClicked.emit(this.condition)
        this.dialogRef.close();
      }
    );
  }

  constructor(
    private complexSearchConditionService: ComplexSearchConditionService,
    private complexSearchService: ComplexSearchService,
    private dialog: MatDialog
  ) {
    this.saveData = this.complexSearchService.initSaveDataObj()
    this.defineDialogSearch();
  }

  ngOnInit() {
  }

}
