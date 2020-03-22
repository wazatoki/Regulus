import { FieldAttr } from './field-attr';
import { SearchCondition } from './search-condition';
import { OrderCondition } from './order-condition';
import { Conditional } from '@angular/compiler';

export interface ConditionData {
    searchStrings: string[],
    displayItemList: FieldAttr[],
    searchConditionList: SearchCondition[],
    orderConditionList: OrderCondition[],
}

export function mapCondition(from: ConditionData, to: ConditionData){
    to.displayItemList = from.displayItemList;
    to.orderConditionList = from.orderConditionList;
    to.searchConditionList = from.searchConditionList;
}