import { SaveData } from './save-data';
import { createTestInstance1 as createStaffInstance1} from 'src/app/services/models/staff/staff.spec';
import { createTestInstance1 as createConditionData, createInitConditionData } from 'src/app/services/models/search/condition-data.spec';
import { ceateTestArrayForMasterMaintenanceTest as createCategoryArrayData } from 'src/app/services/models/search/category.spec';

export function createTestInstance1(): SaveData {
    return {
        id: 'saveID1',
        patternName: 'saveName1',
        category: createCategoryArrayData()[0],
        owner: createStaffInstance1(),
        ownerID: 'ownerID1',
        discloseGroups: [{id: 'test-group-id-1', name: ''}, {id: 'test-group-id-2', name: ''}],
        isDisclose: true,
        conditionData: createConditionData(),
    };
}

export function createTestInstance2(): SaveData {
    return {
        id: 'saveID2',
        patternName: 'saveName2',
        category: createCategoryArrayData()[1],
        owner: createStaffInstance1(),
        ownerID: 'ownerID1',
        discloseGroups: [{id: 'groupID1', name: ''}],
        isDisclose: true,
        conditionData: createConditionData(),
    };
}

export function createTestArray(): SaveData[] {
    return [
        createTestInstance1(),
        createTestInstance2(),
    ];
}

export function createInitSaveData(): SaveData {
    return {
        id: '',
        patternName: '',
        category: null,
        isDisclose: false,
        discloseGroups: [],
        ownerID: '',
        conditionData: createInitConditionData(),
        owner: {
          id: '',
          name: '',
          operatorUsableConditions: [],
        },
      };
}
