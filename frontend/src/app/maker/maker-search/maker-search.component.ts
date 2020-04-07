import { Component, OnInit, EventEmitter, Output, OnDestroy } from '@angular/core';
import { Maker } from '../../services/models/maker/maker';
import { MakerService } from '../../services/api/maker.service';
import { ComplexSearchService } from '../../services/share/complex-search.service';
import { ConditionData, mapCondition } from '../../services/models/search/condition-data';
import { Subscription }   from 'rxjs';

import { MatDialog } from '@angular/material/dialog';
import { ComplexSearchDialogComponent } from '../../layout/dialog/complex-search-dialog/complex-search-dialog/complex-search-dialog.component';
import { ComplexSearchItems } from '../../services/models/search/complex-search-items';

@Component({
  selector: 'app-maker-search',
  templateUrl: './maker-search.component.html',
  styleUrls: ['./maker-search.component.css']
})
export class MakerSearchComponent implements OnInit, OnDestroy {

  private condition: ConditionData;
  private subscription: Subscription;

  @Output() fetched: EventEmitter<Maker[]> = new EventEmitter();

  constructor(
    private makerService: MakerService,
    private complexSearchService: ComplexSearchService,
    private dialog: MatDialog) {

      this.condition = this.complexSearchService.initConditionDataObj();

      // 検索条件設定ダイアログで検索ボタンをクリックしたら検索条件を変更して検索を実行する。
      this.subscription = this.complexSearchService.complexSearchOrdered$.subscribe(
        (data: ConditionData) => {
          mapCondition(data, this.condition);
          this.search();
        }
      );

    }

  ngOnInit() {}

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  openComplexSearch(){
    this.makerService.findComplexSearchItems().subscribe( (data: ComplexSearchItems) => {
      // data.saveData = this.testData(); // 検証用
      const dialogRef = this.dialog.open(ComplexSearchDialogComponent, {
        data: data,
      });
    });
  }

  testData() {
    return {
      category: '',
      discloseGroups: [],
      isDisclose: true,
      patternName: 'asdfg',
      conditionData: {
        displayItemList : [],
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

  onSearch(searchStrings: string) {

    const splitString = '--sprit--string--';
    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = searchStrings.replace(' ', splitString).replace('　', splitString).split(splitString);
    this.search();
  }

  search() {
    this.makerService.findByCondition(this.condition).subscribe(
      makers => {
        this.fetched.emit(makers);
      }
    );
  }

}
