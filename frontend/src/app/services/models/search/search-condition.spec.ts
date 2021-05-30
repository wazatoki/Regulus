import { SearchCondition } from './search-condition';
import { createTestInstance1 as createFieldAttr1,
    createTestInstance2 as createFieldAttr2} from './field-attr.spec';

export function createTestInstance1(): SearchCondition {
    return {
        searchField: createFieldAttr1(),
        conditionValue: 'test value',
        matchType: {value: 'match'},
        operator: {value: 'and'},
    }
}

export function createTestInstance2(): SearchCondition {
    return {
        searchField: createFieldAttr2(),
        conditionValue: 'test value2',
        matchType: {value: 'unmatch'},
        operator: {value: 'or'},
    }
}

export function ceateTestArray(): SearchCondition[] {
    return [
        createTestInstance1(),
        createTestInstance2(),
    ];
}