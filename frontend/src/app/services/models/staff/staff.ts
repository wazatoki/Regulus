import { SaveData } from '../search/save-data';

export interface Staff {
    id: string;
    name: string;
    operatorUsableConditions: SaveData[];
}
