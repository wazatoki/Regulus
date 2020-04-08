import { Component, OnInit, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit {

  searchStrings: string;
  
  @Output() searchClick: EventEmitter<string> = new EventEmitter();

  constructor() {
    this.searchStrings = '';
   }

  ngOnInit() {
  }

  onSearchClick() {
    this.searchClick.emit(this.searchStrings);
  }

}
