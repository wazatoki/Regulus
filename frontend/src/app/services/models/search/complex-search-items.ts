import { FieldAttr } from './field-attr';
import { Group } from '../group/group';
import { SaveData } from './save-data';

export interface ComplexSearchItems {
    displayItemList: FieldAttr[],
    searchConditionList: FieldAttr[],
    orderConditionList: FieldAttr[],
    isShowDisplayItem: boolean,
    isShowOrderCondition: boolean,
    isShowSaveCondition: boolean,
    groupList: Group[],
    saveData: SaveData,
}
