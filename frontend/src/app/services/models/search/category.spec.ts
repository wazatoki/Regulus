import { Category } from './category';

describe('Category', () => {
    it('should create an instance', () => {
        expect(new Category()).toBeTruthy();
    });
});

export function ceateTestArrayForMasterMaintenanceTest(): Category[] {
    return [
        {
            name: 'testCategory1',
            viewValue: 'TEST_CATEGORY_1',
            isShowDisplayItem: true,
            isShowOrderCondition: true,
            isShowSaveCondition: true,
        },
        {
            name: 'testCategory2',
            viewValue: 'TEST_CATEGORY_2',
            isShowDisplayItem: false,
            isShowOrderCondition: true,
            isShowSaveCondition: true,
        },
        {
            name: 'testCategory3',
            viewValue: 'TEST_CATEGORY_3',
            isShowDisplayItem: false,
            isShowOrderCondition: false,
            isShowSaveCondition: true,
        },
    ];
}