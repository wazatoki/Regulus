import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-delete',
  templateUrl: './delete.component.html',
  styleUrls: ['./delete.component.css']
})
export class DeleteComponent implements OnInit {

  @Output() clicked: EventEmitter<string> = new EventEmitter();

  @Input() set buttonLabelName(name: string) {
    if (!(name === '' || name === null || name === undefined)) {
      this.buttonLabel = name;
    }
  }

  buttonLabel: string;

  constructor() {
    this.buttonLabel = '削除';
  }

  ngOnInit() {
  }

  onClick() {
    this.clicked.emit('deleteClicked');
  }

}
