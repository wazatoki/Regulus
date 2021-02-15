import { ConditionData } from './condition-data'
import { Category } from './category'
import { Staff } from '../staff/staff';
import { Group } from '../group/group';

/*

SaveData 検索条件を保存するときのデータ構造

*/
export interface SaveData {
    id: string
    patternName: string
    category: Category
    isDisclose: boolean
    discloseGroups: Group[]
    ownerID: string
    owner: Staff
    conditionData: ConditionData
  }