import { Component, OnInit,Input } from '@angular/core';
import { FormControl, FormArray, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-complex-search-condition-item',
  templateUrl: './complex-search-condition-item.component.html',
  styleUrls: ['./complex-search-condition-item.component.css']
})
export class ComplexSearchConditionItemComponent implements OnInit {

  readonly matchTypesForString: matchTypeAttr[] = [
    {name: 'match', viewValue: '完全一致'},
    {name: 'unmatch', viewValue: '不一致'},
    {name: 'pertialmatch', viewValue: '部分一致'},
  ];
  readonly matchTypesForNumber: matchTypeAttr[] = [
    {name: 'match', viewValue: '完全一致'},
    {name: 'unmatch', viewValue: '不一致'},
    {name: 'gt', viewValue: '>'},
    {name: 'ge', viewValue: '>='},
    {name: 'le', viewValue: '<'},
    {name: 'lt', viewValue: '<='},
  ];
  readonly operators: string[] = ['and','or'];

  get fieldSelected() {
    return this.formGroup.get('fieldSelected') as FormControl;
  }

  get conditionValue() {
    return this.formGroup.get('conditionValue') as FormControl;
  }
  
  get matchTypeSelected() {
    return this.formGroup.get('matchTypeSelected') as FormControl;
  }
  
  get operatorSelected() {
    return this.formGroup.get('operatorSelected') as FormControl;
  }

  matchTypes: matchTypeAttr[];
  @Input() fields: fieldAttr[];
  @Input() formGroup: FormGroup;

  onSelectField(): void {
    if (this.fieldSelected.value.fieldType === "string") {
      this.matchTypes = this.matchTypesForString;
    } else {
      this.matchTypes = this.matchTypesForNumber;
    }
  }

  constructor() {
  }

  ngOnInit() {
    this.matchTypes = this.matchTypesForString;
    this.operatorSelected.setValue('and');
  }

}

interface fieldAttr {
  entityName: string,
  fieldName: string,
  viewValue: string,
  fieldType: string,
}

interface matchTypeAttr {
  name: string,
  viewValue: string,
}
