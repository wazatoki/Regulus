import { from } from "rxjs";
import { SaveData } from '../search/save-data'

export interface Staff {
    id: string
    name: string
    favoriteConditions: SaveData[]
}
