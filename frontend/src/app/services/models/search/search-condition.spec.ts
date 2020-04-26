import { SearchCondition } from './search-condition';
import { createTestInstance1 as createFieldAttr1,
    createTestInstance2 as createFieldAttr2} from './field-attr.spec';

export function createTestInstance1(): SearchCondition {
    return {
        field: createFieldAttr1(),
        conditionValue: 'test value',
        matchType: 'match',
        operator: 'and',
    }
}

export function createTestInstance2(): SearchCondition {
    return {
        field: createFieldAttr2(),
        conditionValue: 'test value2',
        matchType: 'unmatch',
        operator: 'or',
    }
}

export function ceateTestArray(): SearchCondition[] {
    return [
        createTestInstance1(),
        createTestInstance2(),
    ];
}