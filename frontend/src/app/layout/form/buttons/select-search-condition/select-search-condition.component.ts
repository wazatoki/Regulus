import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-select-search-condition',
  templateUrl: './select-search-condition.component.html',
  styleUrls: ['./select-search-condition.component.css']
})
export class SelectSearchConditionComponent implements OnInit {

  @Output() clicked: EventEmitter<string> = new EventEmitter();

  @Input() set buttonLabelName(name: string) {
    if( !(name === '' || name === null || name === undefined) ){
      this.buttonLabel = name;
    }
  }

  private buttonLabel: string;

  constructor() {
    this.buttonLabel = '条件選択';
  }

  ngOnInit() {
  }

  onClick() {
    this.clicked.emit('selectSearchConditionClicked');
  }

}
