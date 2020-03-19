import { FieldAttr } from './field-attr';

export interface SearchCondition {
    field: FieldAttr,
    conditionValue: string,
    matchType: string,
    operator: string,
}
