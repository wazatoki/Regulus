import { ConditionData } from './condition-data';
import {
    createTestInstance1 as createFieldAttr1,
    createTestInstance2 as createFieldAttr2,
    createTestArray as createFieldAttrArray
} from './field-attr.spec';
import { ceateTestArray as createSearchConditionArray } from './search-condition.spec';
import { ceateTestArray as createOrderConditionArray } from './order-condition.spec';

describe('Category', () => {
    it('should create an instance', () => {
        expect(new ConditionData()).toBeTruthy();
    });
});

export function createTestInstance1(): ConditionData {
    return {
        displayItemList: createFieldAttrArray(),
        searchConditionList: createSearchConditionArray(),
        orderConditionList: createOrderConditionArray(),
        searchStrings: ['name']
    };
}
