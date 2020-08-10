import { Component, OnInit, Input, Output, EventEmitter, ElementRef } from '@angular/core';

@Component({
  selector: 'app-submit',
  templateUrl: './submit.component.html',
  styleUrls: ['./submit.component.css']
})
export class SubmitComponent implements OnInit {

  private buttonLabel: string;

  onBlur(){
    this.focusout.emit('');
  }

  @Input() set focus(a: boolean) {
    if (a) {
      this.el.nativeElement.querySelector('button').focus()
    }
  }

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
  @Output() focusout: EventEmitter<string> = new EventEmitter();


  constructor(private el: ElementRef) {
    this.buttonLabel = '送信'
  }

  ngOnInit() {
  }

  onClick() {
    this.clicked.emit('submitClicked');
  }

}
