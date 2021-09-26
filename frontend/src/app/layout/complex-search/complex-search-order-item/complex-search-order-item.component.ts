import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { FieldAttr } from '../../../services/models/search/field-attr';

@Component({
  selector: 'app-complex-search-order-item',
  templateUrl: './complex-search-order-item.component.html',
  styleUrls: ['./complex-search-order-item.component.css']
})
export class ComplexSearchOrderItemComponent implements OnInit {

  readonly orderFieldKeyWords: orderKeyWordAttr[] = [
    { name: 'asc', viewValue: '昇順' },
    { name: 'desc', viewValue: '降順' },
  ];

  get orderFieldSelected() {
    return this.formGroup.get('orderFieldSelected') as FormControl;
  }

  get orderFieldKeyWordSelected() {
    return this.formGroup.get('orderFieldKeyWordSelected') as FormControl;
  }

  @Input() fields: FieldAttr[] = [];
  @Input() formGroup: FormGroup;
  @Output() onDelete = new EventEmitter();

  constructor() { }

  ngOnInit() {
    this.orderFieldKeyWordSelected.setValue(this.orderFieldKeyWords[0].name);
  }

  deleteClicked() {
    this.onDelete.emit();
  }

}

interface orderKeyWordAttr {
  name: string,
  viewValue: string,
}
