import { Component, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-cancel',
  templateUrl: './cancel.component.html',
  styleUrls: ['./cancel.component.css']
})
export class CancelComponent implements OnInit {

  @Output() clicked: EventEmitter<string> = new EventEmitter();

  constructor() { }

  ngOnInit() {
  }

  onClick() {
    this.clicked.emit('cancelClicked');
  }

}
