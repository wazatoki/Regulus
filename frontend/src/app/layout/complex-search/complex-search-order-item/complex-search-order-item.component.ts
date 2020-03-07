import { Component, OnInit, Input } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { fieldAttr } from '../complex-search.component'

@Component({
  selector: 'app-complex-search-order-item',
  templateUrl: './complex-search-order-item.component.html',
  styleUrls: ['./complex-search-order-item.component.css']
})
export class ComplexSearchOrderItemComponent implements OnInit {

  readonly orderFieldKeyWords: orderKeyWordAttr[] = [
    {name: 'asc', viewValue:'昇順'},
    {name: 'desc',  viewValue:'降順'},
  ];

  get orderFieldSelected() {
    return this.formGroup.get('orderFieldSelected') as FormControl;
  }

  get orderFieldKeyWordSelected() {
    return this.formGroup.get('orderFieldKeyWordSelected') as FormControl;
  }

  @Input() orderFields: fieldAttr[];
  @Input() formGroup: FormGroup;

  constructor() { }

  ngOnInit() { }

}

interface orderKeyWordAttr{
  name: string,
  viewValue: string,
}
