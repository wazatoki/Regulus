import { FieldAttr } from './field-attr';
import { SearchCondition } from './search-condition';
import { OrderCondition } from './order-condition';

export interface ConditionData {
    searchStrings: string[],
    displayItemList: FieldAttr[],
    searchConditionList: SearchCondition[],
    orderConditionList: OrderCondition[],
}
