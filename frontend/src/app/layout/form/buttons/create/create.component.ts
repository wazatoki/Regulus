import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {

  @Output() clicked: EventEmitter<string> = new EventEmitter();

  @Input() set buttonLabelName(name: string) {
    if (!(name === '' || name === null || name === undefined)) {
      this.buttonLabel = name;
    }
  }

  private buttonLabel: string;

  constructor() {
    this.buttonLabel = '新規作成';
  }

  ngOnInit() {
  }

  onClick() {
    this.clicked.emit('createClicked');
  }

}
