import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { SaveData } from 'src/app/services/models/search/save-data';

@Component({
  selector: 'app-favorite-conditions',
  templateUrl: './favorite-conditions.component.html',
  styleUrls: ['./favorite-conditions.component.css']
})
export class FavoriteConditionsComponent implements OnInit {

  @Input() favoriteConditions: SaveData[]
  @Output() selectedCondition: EventEmitter<ConditionData> = new EventEmitter();

  onClick(condition: SaveData) {
    this.selectedCondition.emit(condition.conditionData)
  }

  constructor() { }

  ngOnInit() {
  }

}
