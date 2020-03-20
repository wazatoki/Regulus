import { ConditionData } from './condition-data'

export interface SaveData {
    patternName: string;
    category: string;
    isDisclose: boolean;
    discloseGroups: string[];
    ownerID: string;
    conditionData: ConditionData;
  }