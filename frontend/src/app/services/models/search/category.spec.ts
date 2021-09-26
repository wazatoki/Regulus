import { Category } from './category';
import {
    createTestInstance1,
    createTestInstance2,
    createTestInstance3,
} from 'src/app/services/models/search/complex-search-items.spec';

export function ceateTestArrayForMasterMaintenanceTest(): Category[] {
    return [
        {
            name: 'testCategory1',
            viewValue: 'TEST_CATEGORY_1',
            searchItems: createTestInstance1(),
        },
        {
            name: 'testCategory2',
            viewValue: 'TEST_CATEGORY_2',
            searchItems: createTestInstance2(),
        },
        {
            name: 'testCategory3',
            viewValue: 'TEST_CATEGORY_3',
            searchItems: createTestInstance3(),
        },
    ];
}
