import { ComplexSearchItems } from './complex-search-items';
import { createTestArray } from 'src/app/services/models/search/field-attr.spec';
import { ceateTestArray as createTestGroupArray } from 'src/app/services/models/group/staff-group.spec';

export function createTestInstance1(): ComplexSearchItems {
    return {
        displayItemList: createTestArray(),
        searchConditionList: createTestArray(),
        orderConditionList: createTestArray(),
        staffGroups: createTestGroupArray(),
        isShowDisplayItem: true,
        isShowSaveCondition: true,
        isShowOrderCondition: true,
    };
}
export function createTestInstance2(): ComplexSearchItems {
    return {
        displayItemList: createTestArray(),
        searchConditionList: createTestArray(),
        orderConditionList: createTestArray(),
        staffGroups: createTestGroupArray(),
        isShowDisplayItem: false,
        isShowSaveCondition: true,
        isShowOrderCondition: true,
    };
}

export function createTestInstance3(): ComplexSearchItems {
    return {
        displayItemList: createTestArray(),
        searchConditionList: createTestArray(),
        orderConditionList: createTestArray(),
        staffGroups: createTestGroupArray(),
        isShowDisplayItem: false,
        isShowSaveCondition: false,
        isShowOrderCondition: true,
    };
}

export function createTestInstance4(): ComplexSearchItems {
    return {
        displayItemList: createTestArray(),
        searchConditionList: createTestArray(),
        orderConditionList: createTestArray(),
        staffGroups: createTestGroupArray(),
        isShowDisplayItem: false,
        isShowSaveCondition: false,
        isShowOrderCondition: false,
    };
}
