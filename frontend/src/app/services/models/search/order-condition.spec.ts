import { OrderCondition } from './order-condition';
import { createTestInstance1 as createFieldAttr1,
    createTestInstance2 as createFieldAttr2} from './field-attr.spec';

export function createTestInstance1(): OrderCondition {
    return {
        orderField: createFieldAttr1(),
        orderFieldKeyWord: {value: 'asc'},
    };
}

export function createTestInstance2(): OrderCondition {
    return {
        orderField: createFieldAttr2(),
        orderFieldKeyWord: {value: 'desc'},
    };
}

export function ceateTestArray(): OrderCondition[] {
    return [
        createTestInstance1(),
        createTestInstance2(),
    ];
}
