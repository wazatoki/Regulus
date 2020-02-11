import { Component, OnInit,Input } from '@angular/core';

@Component({
  selector: 'app-complex-search-condition-item',
  templateUrl: './complex-search-condition-item.component.html',
  styleUrls: ['./complex-search-condition-item.component.css']
})
export class ComplexSearchConditionItemComponent implements OnInit {

  readonly matchTypesForString: string[] = ['match', 'unmatch', 'pertialmatch'];
  readonly matchTypesForNumber: string[] = ['match', 'unmatch', 'gt', 'ge', 'le', 'lt'];
  readonly operators: string[] = ['and','or'];

  fieldSelected: fieldAttr;
  conditionValue: string;
  matchTypeSelected: string;
  operatorSelected: string;

  matchTypes: string[];
  fields: fieldAttr[];

  onSelectField(): void {
    if (this.fieldSelected.fieldType === "string") {
      this.matchTypes = this.matchTypesForString;
    } else {
      this.matchTypes = this.matchTypesForNumber;
    }
  }

  constructor() {
    this.matchTypes = this.matchTypesForString;
    this.operatorSelected = 'and'
  }

  ngOnInit() {
  }

}

interface fieldAttr {
  entityName: string,
  fieldName: string,
  viewValue: string,
  fieldType: string,
}
