import { Component, OnInit, EventEmitter, Output, OnDestroy, Input } from '@angular/core';
import { Maker } from '../../services/models/maker/maker';
import { MakerService } from '../../services/api/maker.service';
import { ComplexSearchService } from '../../services/share/complex-search.service';
import { ConditionData, mapCondition, splitStrings } from '../../services/models/search/condition-data';
import { Subscription } from 'rxjs';

import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { ComplexSearchDialogComponent } from '../../layout/dialog/complex-search-dialog/complex-search-dialog/complex-search-dialog.component';
import { ComplexSearchItems } from '../../services/models/search/complex-search-items';
import { SaveData } from 'src/app/services/models/search/save-data';

@Component({
  selector: 'app-maker-search',
  templateUrl: './maker-search.component.html',
  styleUrls: ['./maker-search.component.css']
})
export class MakerSearchComponent implements OnInit, OnDestroy {

  private condition: ConditionData;
  private complexSearchSubscription: Subscription;
  private saveData: SaveData;
  private dialogRef: MatDialogRef<ComplexSearchDialogComponent>;

  @Input() set savedData(v: SaveData) {// 値の入力があり有意な値であった場合は検索実行
    
    if ( v ) {
      
      v.category.name = 'maker'
      this.saveData = v;

      if( v.id ){
        mapCondition(v.conditionData, this.condition);
        this.search();
      }

    }
  };

  @Output() fetched: EventEmitter<Maker[]> = new EventEmitter();

  constructor(
    private makerService: MakerService,
    private complexSearchService: ComplexSearchService,
    private dialog: MatDialog) {

    this.condition = this.complexSearchService.initConditionDataObj();

    this.defineDialogSearch();
  }

  ngOnInit() { }

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

  openComplexSearch() {
    this.makerService.findComplexSearchItems().subscribe((data: ComplexSearchItems) => {
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
    this.search();
  }

  search() {
    this.makerService.findByCondition(this.condition).subscribe(
      makers => {
        this.fetched.emit(makers);
      }
    );
  }

  testData() {
    return {
      category: '',
      discloseGroups: [],
      isDisclose: true,
      patternName: 'asdfg',
      conditionData: {
        displayItemList: [],
        orderConditionList: [],
        searchConditionList: [
          {
            conditionValue: 'abcd',
            field: {
              id: 'maker.name',
              entityName: 'maker',
              fieldName: 'name',
              fieldType: 'string',
              viewValue: 'メーカー名称',
            },
            matchType: 'unmatch',
            operator: 'or'
          },
        ],
        searchStrings: [],
      },
      id: 'savedataid',
      ownerID: '',
    };
  }

}
