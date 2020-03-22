import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { Maker } from '../../services/models/maker/maker';
import { MakerService } from '../../services/api/maker.service';
import { ComplexSearchService } from '../../services/share/complex-search.service';
import { ConditionData } from '../../services/models/search/condition-data';
import { MakerCondition } from '../../services/models/maker/maker-condition';

@Component({
  selector: 'app-maker-search',
  templateUrl: './maker-search.component.html',
  styleUrls: ['./maker-search.component.css']
})
export class MakerSearchComponent implements OnInit {

  private condition: ConditionData;

  constructor(
    private makerService: MakerService,
    private complexSearchService: ComplexSearchService) { }

  ngOnInit() {
    this.condition = this.complexSearchService.initConditionDataObj();
  }

  @Output() fetched: EventEmitter<Maker[]> = new EventEmitter();

  onSearch(searchStrings: string) {

    const splitString = '--sprit--string--';
    // 全角文字を一旦区切り文字列に置き換えて配列に分割
    this.condition.searchStrings = searchStrings.replace('　', splitString).split(splitString);
    this.makerService.findByCondition(this.condition).subscribe(
      makers => {
        this.fetched.emit(makers);
      }
    );
  }
}
