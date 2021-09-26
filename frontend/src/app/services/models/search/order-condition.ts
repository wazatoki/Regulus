import { FieldAttr } from './field-attr';

/*

OrderCondition 並び順設定のためのパーツ

*/
export interface OrderCondition {
    orderField: FieldAttr;
    orderFieldKeyWord: {value: string};
}
