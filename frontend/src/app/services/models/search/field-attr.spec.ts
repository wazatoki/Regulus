import { FieldAttr } from './field-attr';

describe('Category', () => {
    it('should create an instance', () => {
        expect(new FieldAttr()).toBeTruthy();
    });
});

export function createTestInstance1(): FieldAttr {
    return {
        id: 'fieldid1',
        entityName: 'testField',
        fieldName: 'id',
        viewValue: 'ID',
        fieldType: 'string',
    }
}
export function createTestInstance2(): FieldAttr {
    return {
        id: 'fieldid2',
        entityName: 'testField',
        fieldName: 'name',
        viewValue: 'NAME',
        fieldType: 'string',
    }
}

export function createTestArray(): FieldAttr[] {
    return [
        createTestInstance1(),
        createTestInstance2(),
    ];
}