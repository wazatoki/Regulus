import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { SaveData } from 'src/app/services/models/search/save-data';

@Component({
  selector: 'app-operator-usable-conditions',
  templateUrl: './operator-usable-conditions.component.html',
  styleUrls: ['./operator-usable-conditions.component.css']
})
export class OperatorUsableConditionsComponent implements OnInit {

  @Input() usableConditions: SaveData[];

  @Output() selectedCondition: EventEmitter<SaveData> = new EventEmitter();

  public get conditions(): SaveData[] {
    return this.usableConditions;
  }

  onClick(condition: SaveData) {
    this.selectedCondition.emit(condition);
  }

  constructor() { }

  ngOnInit() {
  }

}
