import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css']
})
export class UpdateComponent implements OnInit {

  @Output() clicked: EventEmitter<string> = new EventEmitter();

  @Input() set buttonLabelName(name: string) {
    if (!(name === '' || name === null || name === undefined)) {
      this.buttonLabel = name;
    }
  }

  private buttonLabel: string;

  constructor() {
    this.buttonLabel = '修正';
  }

  ngOnInit() {
  }

  onClick(event: any) {
    event.preventDefault();
    this.clicked.emit('updateClicked');
  }

}
