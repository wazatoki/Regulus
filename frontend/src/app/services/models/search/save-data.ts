import { ConditionData } from './condition-data'
import { Staff } from '../staff/staff';

export interface SaveData {
    id: string,
    patternName: string,
    category: string,
    isDisclose: boolean,
    discloseGroups: string[],
    ownerID: string,
    conditionData: ConditionData,
    owner: Staff
  }