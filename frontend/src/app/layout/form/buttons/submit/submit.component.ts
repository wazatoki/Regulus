import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-submit',
  templateUrl: './submit.component.html',
  styleUrls: ['./submit.component.css']
})
export class SubmitComponent implements OnInit {

  private buttonLabel: string;

  @Input() set buttonLabelType(type: string) {

    switch (type) {
      case 'insert':
        this.buttonLabel = '新規作成'
        break;
      case 'update':
        this.buttonLabel = '更新'
        break;
      case 'save':
        this.buttonLabel = '保存'
        break;
      default:
        this.buttonLabel = '送信'
        break;
    }
  }

  @Output() clicked: EventEmitter<string> = new EventEmitter();

  constructor() {
    this.buttonLabel = '送信'
  }

  ngOnInit() {
  }

  onClick() {
    this.clicked.emit('submitClicked');
  }

}
