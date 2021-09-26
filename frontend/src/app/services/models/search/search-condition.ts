import { FieldAttr } from './field-attr';

/*

SearchCondition 検索条件のパーツ

*/
export interface SearchCondition {
    searchField: FieldAttr;
    conditionValue: string;
    matchType: {value: string};
    operator: {value: string};
}
