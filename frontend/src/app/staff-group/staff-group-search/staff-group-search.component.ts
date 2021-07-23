import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ConditionData, splitStrings } from 'src/app/services/models/search/condition-data';

@Component({
  selector: 'app-staff-group-search',
  templateUrl: './staff-group-search.component.html',
  styleUrls: ['./staff-group-search.component.css']
})
export class StaffGroupSearchComponent implements OnInit {

  @Input() condition: ConditionData;

  @Output() searchClicked: EventEmitter<ConditionData> = new EventEmitter();
  
  onSearch(searchStrings: string) {

    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = splitStrings(searchStrings)
    this.searchClicked.emit(this.condition)
  }

  constructor() { }

  ngOnInit() { }

}
