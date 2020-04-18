import { SaveData } from './save-data';
import { createTestInstance1 as createStaffInstance1} from 'src/app/services/models/staff/staff.spec';
import { createTestInstance1 as createConditionData } from 'src/app/services/models/search/condition-data.spec';

describe('Category', () => {
    it('should create an instance', () => {
        expect(new SaveData()).toBeTruthy();
    });
});

export function createTestInstance1(): SaveData {
    return {
        id: 'saveID1',
        patternName: 'saveName1',
        category: 'test-category-1',
        owner: createStaffInstance1(),
        ownerID: 'ownerID1',
        discloseGroupIDs: ['test-group-id-1', 'test-group-id-2'],
        isDisclose: true,
        conditionData: createConditionData(),
    }
}

export function createTestInstance2(): SaveData {
    return {
        id: 'saveID2',
        patternName: 'saveName2',
        category: 'test-category-2',
        owner: createStaffInstance1(),
        ownerID: 'ownerID1',
        discloseGroupIDs: ['groupID1'],
        isDisclose: true,
        conditionData: createConditionData(),
    }
}

export function createTestArray(): SaveData[]{
    return [
        createTestInstance1(),
        createTestInstance2(),
    ]
}
