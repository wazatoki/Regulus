import { ConditionData } from './condition-data'

export interface SaveData {
    id: string,
    patternName: string,
    category: string,
    isDisclose: boolean,
    discloseGroups: string[],
    ownerID: string,
    conditionData: ConditionData,
  }