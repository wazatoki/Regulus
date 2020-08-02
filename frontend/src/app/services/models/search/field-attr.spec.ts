import { FieldAttr } from './field-attr';

describe('Category', () => {
    it('should create an instance', () => {
        expect(new FieldAttr()).toBeTruthy();
    });
});

export function createTestInstance1(): FieldAttr {
    return {
        id: 'fieldid1',
        viewValue: 'ID',
        fieldType: 'string',
        optionItems: []
    }
}
export function createTestInstance2(): FieldAttr {
    return {
        id: 'fieldid2',
        viewValue: 'NAME',
        fieldType: 'string',
        optionItems: []
    }
}

export function createTestArray(): FieldAttr[] {
    return [
        createTestInstance1(),
        createTestInstance2(),
    ];
}