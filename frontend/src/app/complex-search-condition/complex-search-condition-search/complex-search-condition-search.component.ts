import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { SaveData } from 'src/app/services/models/search/save-data';
import { ConditionData, splitStrings } from 'src/app/services/models/search/condition-data';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';

@Component({
  selector: 'app-complex-search-condition-search',
  templateUrl: './complex-search-condition-search.component.html',
  styleUrls: ['./complex-search-condition-search.component.css']
})
export class ComplexSearchConditionSearchComponent implements OnInit {

  private condition: ConditionData;

  @Output() fetched: EventEmitter<SaveData[]> = new EventEmitter();

  onSearch(searchStrings: string) {

    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = splitStrings(searchStrings)
    this.search();
  }

  search() {
    this.complexSearchConditionService.findByCondition(this.condition).subscribe(
      conditions => {
        this.fetched.emit(conditions);
      }
    );
  }
  
  constructor( private complexSearchConditionService: ComplexSearchConditionService ) { }

  ngOnInit() {
  }

}
