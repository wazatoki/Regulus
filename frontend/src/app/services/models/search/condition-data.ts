import { FieldAttr } from './field-attr';
import { SearchCondition } from './search-condition';
import { OrderCondition } from './order-condition';

export interface ConditionData {
    searchStrings: string[],
    displayItemList: FieldAttr[],
    searchConditionList: SearchCondition[],
    orderConditionList: OrderCondition[],
}

export function mapCondition(from: ConditionData, to: ConditionData) {
    to.displayItemList = from.displayItemList;
    to.orderConditionList = from.orderConditionList;
    to.searchConditionList = from.searchConditionList;
}

export function splitStrings(str: string): string[] {

    const splitString = '--sprit--string--';

    // 全角空白半角空白を一旦区切り文字列に置き換えて配列に分割
    return str.replace(' ', splitString).replace('　', splitString).split(splitString);
}