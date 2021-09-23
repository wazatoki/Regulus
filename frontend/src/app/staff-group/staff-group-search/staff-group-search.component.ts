import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material';
import { OperatorUsableConditionsDialogComponent } from 'src/app/layout/dialog/operator-usable-conditions-dialog/operator-usable-conditions-dialog.component';
import { LoginService } from 'src/app/services/api/login.service';
import { ConditionData, splitStrings } from 'src/app/services/models/search/condition-data';
import { SaveData } from 'src/app/services/models/search/save-data';

@Component({
  selector: 'app-staff-group-search',
  templateUrl: './staff-group-search.component.html',
  styleUrls: ['./staff-group-search.component.css']
})
export class StaffGroupSearchComponent implements OnInit {

  @Input() condition: ConditionData;

  @Output() searchClicked: EventEmitter<ConditionData> = new EventEmitter();

  public selectedPatternName: string

  onSearch(searchStrings: string) {

    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = splitStrings(searchStrings)
    this.searchClicked.emit(this.condition)
  }

  openDialogSelectSearchCondition() {

    const conditions: SaveData[] = this.loginSsevice.currentUserValue.operatorUsableConditions.filter((d: SaveData) => d.category.name === 'staff-group');

    const dialogRef = this.conditionSelectDialog.open(OperatorUsableConditionsDialogComponent, {
      data: {
        title: '検索条件',
        operatorUsableConditions: conditions
      }
    });

    dialogRef.afterClosed().subscribe(
      (data: SaveData) => {
        this.selectedPatternName = data.patternName
        this.searchClicked.emit(data.conditionData)
      }
    )

  }

  constructor(
    private conditionSelectDialog: MatDialog,
    private loginSsevice: LoginService
  ) { }

  ngOnInit() {
    this.selectedPatternName = '';
   }

}
