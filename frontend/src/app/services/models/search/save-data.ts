import { ConditionData } from './condition-data'
import { Staff } from '../staff/staff';

/*

SaveData 検索条件を保存するときのデータ構造

*/
export class SaveData {
    id: string
    patternName: string
    category: string
    isDisclose: boolean
    discloseGroups: string[]
    ownerID: string
    conditionData: ConditionData
    owner: Staff
  }