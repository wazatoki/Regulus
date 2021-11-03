import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { SaveData } from 'src/app/services/models/search/save-data';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';

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

  onDrop(event: CdkDragDrop<SaveData[]>) {
    moveItemInArray(this.conditions, event.previousIndex, event.currentIndex);
    const data: string[] = [];
    this.conditions.forEach( (c: SaveData) => {
      data.push(c.id)
    })
    this.complexSearchConditionService.updateFavoriteConditions(data)
  }

  onClick(condition: SaveData) {
    this.selectedCondition.emit(condition);
  }

  constructor(
    private complexSearchConditionService: ComplexSearchConditionService
  ) { }

  ngOnInit() {
  }

}
