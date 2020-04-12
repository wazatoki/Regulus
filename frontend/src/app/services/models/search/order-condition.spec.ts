import { OrderCondition } from './order-condition';
import { createTestInstance1 as createFieldAttr1,
    createTestInstance2 as createFieldAttr2} from './field-attr.spec';

describe('Category', () => {
    it('should create an instance', () => {
        expect(new OrderCondition()).toBeTruthy();
    });
});

export function createTestInstance1(): OrderCondition {
    return {
        orderField: createFieldAttr1(),
        orderFieldKeyWord: 'asc',
    }
}

export function createTestInstance2(): OrderCondition {
    return {
        orderField: createFieldAttr2(),
        orderFieldKeyWord: 'desc',
    }
}

export function ceateTestArray(): OrderCondition[] {
    return [
        createTestInstance1(),
        createTestInstance2(),
    ];
}