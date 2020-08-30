import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Subscription } from 'rxjs';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';

import { ComplexSearchDialogComponent } from '../../layout/dialog/complex-search-dialog/complex-search-dialog/complex-search-dialog.component';
import { SaveData } from 'src/app/services/models/search/save-data';
import { ConditionData, mapCondition, splitStrings } from 'src/app/services/models/search/condition-data';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { ComplexSearchItems } from '../../services/models/search/complex-search-items';

@Component({
  selector: 'app-complex-search-condition-search',
  templateUrl: './complex-search-condition-search.component.html',
  styleUrls: ['./complex-search-condition-search.component.css']
})
export class ComplexSearchConditionSearchComponent implements OnInit {

  private condition: ConditionData;
  private complexSearchSubscription: Subscription;
  private saveData: SaveData;
  private dialogRef: MatDialogRef<ComplexSearchDialogComponent>;

  @Output() fetched: EventEmitter<SaveData[]> = new EventEmitter();

  openComplexSearch() {
    this.complexSearchConditionService.findComplexSearchItems().subscribe((data: ComplexSearchItems) => {
      const aData: any = data;
      aData.saveData = this.saveData;
      this.dialogRef = this.dialog.open(ComplexSearchDialogComponent, {
        data: aData.searchItems,
      });
    });
  }

  onSearch(searchStrings: string) {

    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = splitStrings(searchStrings)
    this.search();
  }

  search() {
    this.complexSearchConditionService.findByCondition(this.condition).subscribe(
      (res: SaveData[]) => {
        this.fetched.emit(res);
      }
    );
  }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.complexSearchSubscription.unsubscribe();
  }

  defineDialogSearch() {
    // 検索条件設定ダイアログで検索ボタンをクリックしたら検索条件を変更して検索を実行する。
    this.complexSearchSubscription = this.complexSearchService.complexSearchOrdered$.subscribe(
      (data: ConditionData) => {
        mapCondition(data, this.condition);
        this.search();
        this.dialogRef.close();
      }
    );
  }
  
  constructor(
    private complexSearchConditionService: ComplexSearchConditionService,
    private complexSearchService: ComplexSearchService,
    private dialog: MatDialog
    ) {
      this.condition = complexSearchService.initConditionDataObj();
      this.defineDialogSearch();
    }

  ngOnInit() {
  }

}
