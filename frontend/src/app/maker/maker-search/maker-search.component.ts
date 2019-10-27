import { Component, OnInit } from '@angular/core';
import { Maker } from '../../services/models/maker/maker';
import { MakerService } from '../../services/api/maker.service';
import { MakerCondition } from '../../services/models/maker/maker-condition';

@Component({
  selector: 'app-maker-search',
  templateUrl: './maker-search.component.html',
  styleUrls: ['./maker-search.component.css']
})
export class MakerSearchComponent implements OnInit {

  constructor(
    private makerService: MakerService,
    private makerCondition: MakerCondition) { }

  ngOnInit() {
  }

  makers: Maker[];

  onSearch(data: string) {

    this.makerCondition.searchStrings = data
    this.makerService.findByCondition(this.makerCondition).subscribe(
      makers => {
        this.makers = makers;
      }
    );
  }
}
