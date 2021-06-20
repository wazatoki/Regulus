import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';
import { StaffGroup } from 'src/app/services/models/group/staff-group';
import { ConditionData, splitStrings } from 'src/app/services/models/search/condition-data';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';

@Component({
  selector: 'app-staff-group-search',
  templateUrl: './staff-group-search.component.html',
  styleUrls: ['./staff-group-search.component.css']
})
export class StaffGroupSearchComponent implements OnInit {

  private condition: ConditionData;

  @Output() fetched: EventEmitter<StaffGroup[]> = new EventEmitter();
  
  onSearch(searchStrings: string) {

    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = splitStrings(searchStrings)
    this.search();
  }

  search() {
    this.staffGroupService.findByCondition(this.condition).subscribe(
      (res: StaffGroup[]) => {
        this.fetched.emit(res);
      }
    );
  }

  constructor(
    private staffGroupService: StaffGroupService,
    private complexSearchService: ComplexSearchService,
  ) { 
    this.condition = complexSearchService.initConditionDataObj();
  }

  ngOnInit() {
  }

}
