import { FieldAttr } from './field-attr';

/*

SearchCondition 検索条件のパーツ

*/
export interface SearchCondition {
    field: FieldAttr
    conditionValue: string
    matchType: string
    operator: string
}
