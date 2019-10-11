import { Component, OnInit } from '@angular/core';

export interface ProductElement {
  name: string;
  jancode: string;
  code: string;
}

@Component({
  selector: 'app-product-master',
  templateUrl: './product-master.component.html',
  styleUrls: ['./product-master.component.css']
})
export class ProductMasterComponent implements OnInit {

  displayedColumns: string[];
  dataSource: ProductElement[];

  constructor() { 
    this.displayedColumns = ['jancode', 'name', 'code'];
    this.dataSource = [];
  }

  ngOnInit() {
  }

}
