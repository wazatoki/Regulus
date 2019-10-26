import { TestBed } from '@angular/core/testing';
import { MakerCondition } from './maker-condition';
import { Maker } from './maker'

describe('MakerCondition', () => {

  beforeEach(() => {

    TestBed.configureTestingModule({
      providers:[
        Maker
      ]
    });
  });
  it('should create an instance', () => {
    const makerCondition = TestBed.get(MakerCondition)
    expect(makerCondition).toBeTruthy();
  });
});
