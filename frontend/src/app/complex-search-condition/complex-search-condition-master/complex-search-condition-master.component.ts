import { Component, OnInit } from '@angular/core';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { MatDialog } from '@angular/material/dialog';
import { SaveData } from 'src/app/services/models/search/save-data';
import { SelectionModel } from '@angular/cdk/collections';
import { MatTableDataSource } from '@angular/material/table';

@Component({
  selector: 'app-complex-search-condition-master',
  templateUrl: './complex-search-condition-master.component.html',
  styleUrls: ['./complex-search-condition-master.component.css']
})
export class ComplexSearchConditionMasterComponent implements OnInit {

  displayedColumns: string[];
  dataSource: MatTableDataSource<SaveData>;
  selection: SelectionModel<SaveData>;

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
    checkboxLabel(row?: SaveData): string {
      if (!row) {
        return `${this.isAllSelected() ? 'select' : 'deselect'} all`;
      }
      return `${this.selection.isSelected(row) ? 'deselect' : 'select'} row ${row.id}`;
    }

  constructor( private complexSearchConditionService: ComplexSearchConditionService,
    public dialog: MatDialog
    ) { 
      const initialSelection = [];
      const allowMultiSelect = true;
      this.displayedColumns = ['select', 'name', 'category', 'owner'];
      this.dataSource = new MatTableDataSource<SaveData>([]);
      this.selection = new SelectionModel<SaveData>(allowMultiSelect, initialSelection);

    }

  ngOnInit() {
  }

}
