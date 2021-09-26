import { Component, OnInit, Output, EventEmitter, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';

import {
  ComplexSearchDialogComponent
} from '../../layout/dialog/complex-search-dialog/complex-search-dialog/complex-search-dialog.component';
import { SaveData } from 'src/app/services/models/search/save-data';
import { ConditionData, mapCondition, splitStrings } from 'src/app/services/models/search/condition-data';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { ComplexSearchItems } from '../../services/models/search/complex-search-items';
import { LoginService } from 'src/app/services/api/login.service';
import {
  OperatorUsableConditionsDialogComponent
} from 'src/app/layout/dialog/operator-usable-conditions-dialog/operator-usable-conditions-dialog.component';

@Component({
  selector: 'app-complex-search-condition-search',
  templateUrl: './complex-search-condition-search.component.html',
  styleUrls: ['./complex-search-condition-search.component.css']
})
export class ComplexSearchConditionSearchComponent implements OnInit, OnDestroy {

  private complexSearchSubscription: Subscription;
  private saveData: SaveData;
  private dialogRef: MatDialogRef<ComplexSearchDialogComponent>;
  private condition: ConditionData;

  public selectedPatternName: string;

  @Output() searchClicked: EventEmitter<ConditionData> = new EventEmitter();

  get isConditionDataAvailable(): boolean {

    return this.saveData.conditionData.searchConditionList.length > 0 ||
      this.saveData.conditionData.displayItemList.length > 0 ||
      this.saveData.conditionData.orderConditionList.length > 0;
  }

  openDialogSelectSearchCondition() {

    const conditions: SaveData[] = this.loginSsevice.currentUserValue.operatorUsableConditions
      .filter((d: SaveData) => d.category.name === 'query-condition');

    const dialogRef = this.conditionSelectDialog.open(OperatorUsableConditionsDialogComponent, {
      data: {
        title: '検索条件',
        operatorUsableConditions: conditions
      }
    });

    dialogRef.afterClosed().subscribe(
      (data: SaveData) => {
        this.selectedPatternName = data.patternName;
        this.saveData = data;
        this.condition = data.conditionData;
        this.searchClicked.emit(data.conditionData);
      }
    );

  }

  clearCondition() {
    this.saveData = this.complexSearchService.initSaveDataObj();
    this.condition = this.saveData.conditionData;
    this.selectedPatternName = '';
  }

  openComplexSearch() {
    this.complexSearchConditionService.findComplexSearchItems().subscribe((data: ComplexSearchItems) => {
      const aData: {
        title: string,
        complexSearchItems: ComplexSearchItems,
        saveData: SaveData,
      } = {
        title: '検索条件設定',
        complexSearchItems: data,
        saveData: this.saveData
      };

      // aData.saveData = this.testData(); // 検証用
      this.dialogRef = this.conditionDialog.open(ComplexSearchDialogComponent, {
        data: aData,
      });
    });
  }

  onSearch(searchStrings: string) {

    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = splitStrings(searchStrings);
    this.searchClicked.emit(this.condition);
  }

  ngOnDestroy() {

    // prevent memory leak when component destroyed
    this.complexSearchSubscription.unsubscribe();
  }

  defineDialogSearch() {
    // 検索条件設定ダイアログで検索ボタンをクリックしたら検索条件を変更して検索を実行する。
    this.complexSearchSubscription = this.complexSearchService.complexSearchOrdered$.subscribe(
      (data: ConditionData) => {
        this.saveData.conditionData = data;
        mapCondition(data, this.condition);
        this.searchClicked.emit(this.condition);
        this.dialogRef.close();
      }
    );
  }

  constructor(
    private complexSearchConditionService: ComplexSearchConditionService,
    private complexSearchService: ComplexSearchService,
    private conditionDialog: MatDialog,
    private conditionSelectDialog: MatDialog,
    private loginSsevice: LoginService
  ) {
    this.clearCondition();
    this.defineDialogSearch();
  }

  ngOnInit() {
  }

}
