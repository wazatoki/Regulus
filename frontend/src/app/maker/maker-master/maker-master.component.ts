import { SelectionModel } from '@angular/cdk/collections';
import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';


import { MakerService } from '../../services/api/maker.service';
import { Maker } from '../../services/models/maker/maker';
import { MakerCondition } from '../../services/models/maker/maker-condition';
import { MakerInputFormComponent } from '../maker-input-form/maker-input-form.component';

@Component({
  selector: 'app-maker-master',
  templateUrl: './maker-master.component.html',
  styleUrls: ['./maker-master.component.css']
})
export class MakerMasterComponent implements OnInit {

  
  displayedColumns: string[];
  dataSource: MatTableDataSource<Maker>;
  selection: SelectionModel<Maker>;


  name: string;

  constructor(
    private makerService: MakerService,
    private makerCondition: MakerCondition,
    private dialog: MatDialog) {
    const initialSelection = [];
    const allowMultiSelect = true;
    this.displayedColumns = ['select', 'name'];
    this.dataSource = new MatTableDataSource<Maker>([]);
    this.selection = new SelectionModel<Maker>(allowMultiSelect, initialSelection);

  }

  ngOnInit() {
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.dataSource.data.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.dataSource.data.forEach(row => this.selection.select(row));
  }

  /** The label for the checkbox on the passed row */
  checkboxLabel(row?: Maker): string {
    if (!row) {
      return `${this.isAllSelected() ? 'select' : 'deselect'} all`;
    }
    return `${this.selection.isSelected(row) ? 'deselect' : 'select'} row ${row.id}`;
  }

  onFetchedMakers(data: Maker[]) {
    this.dataSource = new MatTableDataSource(data);
  }

  openInputForm(): void {
    const dialogRef = this.dialog.open(MakerInputFormComponent, {
      width: '500px',
      data: { name: this.name }
    });

    dialogRef.afterClosed().subscribe(result => {
      this.name = '';
    });
  }
}
