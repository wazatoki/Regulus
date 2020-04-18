import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { FieldAttr } from '../../../services/models/search/field-attr';

@Component({
  selector: 'app-complex-search-condition-item',
  templateUrl: './complex-search-condition-item.component.html',
  styleUrls: ['./complex-search-condition-item.component.css']
})
export class ComplexSearchConditionItemComponent implements OnInit {

  readonly matchTypesForString: matchTypeAttr[] = [
    { name: 'match', viewValue: '完全一致' },
    { name: 'unmatch', viewValue: '不一致' },
    { name: 'pertialmatch', viewValue: '部分一致' },
  ];
  readonly matchTypesForNumber: matchTypeAttr[] = [
    { name: 'match', viewValue: '完全一致 =' },
    { name: 'unmatch', viewValue: '不一致 !=' },
    { name: 'gt', viewValue: '超過 >' },
    { name: 'ge', viewValue: '以上 >=' },
    { name: 'le', viewValue: '未満 <' },
    { name: 'lt', viewValue: '以下 <=' },
  ];
  readonly operators: string[] = ['and', 'or'];

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

  @Input() fields: FieldAttr[] = [];
  @Input() formGroup: FormGroup;
  @Output() onDelete = new EventEmitter();

  onSelectField(): void {
    this.setMatchType();
  }

  constructor() {
  }

  ngOnInit() {
    // matchTypeの初期設定
    if (this.fieldSelected.value !== null && this.fieldSelected.value !== undefined && this.fieldSelected.value !== '') {
      this.setMatchType();

    } else {
      this.matchTypes = this.matchTypesForString;
    }
    // operatorの初期設定
    if (this.operatorSelected.value !== null && this.operatorSelected.value !== undefined && this.operatorSelected.value == '') {
      this.operatorSelected.setValue(this.operators[0]);
    }
  }

  setMatchType(): void {
    const f = this.fields.find((field) => {
      return (field.id === this.fieldSelected.value)
    })

    if (f) {
      switch (f.fieldType) {
        case 'number':
          this.matchTypes = this.matchTypesForNumber;
          break;
        case 'string':
          this.matchTypes = this.matchTypesForString;
          break;

        default:
          this.matchTypes = this.matchTypesForString;
          break;
      }
    } else { // 検索対象フィールドが選択されていないときなど。
      this.matchTypes = this.matchTypesForString;
    }
  }

  deleteClicked() {
    this.onDelete.emit();
  }

}

interface matchTypeAttr {
  name: string,
  viewValue: string,
}
