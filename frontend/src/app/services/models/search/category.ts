import { ComplexSearchItems } from './complex-search-items';

export interface Category {
    name: string;
    viewValue: string;
    searchItems: ComplexSearchItems;
}
