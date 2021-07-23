import { FieldAttr } from './field-attr';
import { StaffGroup } from '../group/staff-group';

/*

ComplexSearchItems 検索フォームを構築する際に必要な選択肢アイテム

*/
export interface ComplexSearchItems {
    displayItemList: FieldAttr[]
    searchConditionList: FieldAttr[]
    orderConditionList: FieldAttr[]
    isShowDisplayItem: boolean
    isShowOrderCondition: boolean
    isShowSaveCondition: boolean
    staffGroups: StaffGroup[]
}
