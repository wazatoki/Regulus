import { Component, OnInit, Input, Output, EventEmitter, ElementRef } from '@angular/core';

@Component({
  selector: 'app-submit',
  templateUrl: './submit.component.html',
  styleUrls: ['./submit.component.css']
})
export class SubmitComponent implements OnInit {

  private buttonLabel: string;

  @Input() set focus(a: boolean) {
    if (a) {
      this.el.nativeElement.querySelector('button').focus();
    }
  }

  @Input() set buttonLabelType(type: string) {

    switch (type) {
      case 'insert':
        this.buttonLabel = '新規作成';
        break;
      case 'update':
        this.buttonLabel = '更新';
        break;
      case 'save':
        this.buttonLabel = '保存';
        break;
      default:
        this.buttonLabel = '送信';
        break;
    }
  }

  @Output() clicked: EventEmitter<string> = new EventEmitter();
  @Output() focusouted: EventEmitter<string> = new EventEmitter();

  onBlur() {
    this.focusouted.emit('');
  }

  onClick() {
    this.clicked.emit('submitClicked');
  }

  constructor(private el: ElementRef) {
    this.buttonLabel = '送信';
  }

  ngOnInit() {
  }
}
