import { Component, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-clear',
  templateUrl: './clear.component.html',
  styleUrls: ['./clear.component.css']
})
export class ClearComponent implements OnInit {

  @Output() clicked: EventEmitter<string> = new EventEmitter();

  constructor() { }

  ngOnInit() {
  }

  onClick() {
    this.clicked.emit('clearClicked');
  }

}
