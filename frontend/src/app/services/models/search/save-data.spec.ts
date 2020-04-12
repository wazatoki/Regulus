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
        id: 'id1',
        patternName: 'test-pattern-name',
        category: 'test-category-1',
        owner: createStaffInstance1(),
        ownerID: 'ownerID1',
        discloseGroups: ['groupID1', 'groupID2'],
        isDisclose: true,
        conditionData: createConditionData(),
    }
}
