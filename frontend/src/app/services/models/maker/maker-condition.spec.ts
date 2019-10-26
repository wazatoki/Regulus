import { TestBed } from '@angular/core/testing';
import { MakerCondition } from './maker-condition';
import { Maker } from './maker'

describe('MakerCondition', () => {

  beforeEach(() => {

    TestBed.configureTestingModule({
      providers: [
        Maker
      ]
    });
  });

  it('should create an instance', () => {
    const makerCondition = TestBed.get(MakerCondition);
    expect(makerCondition).toBeTruthy();
  });

  it('test property id', () => {
    const makerCondition: MakerCondition = TestBed.get(MakerCondition);
    makerCondition.id = 'testid';
    expect(makerCondition.id).toEqual('testid');
  });

  it('test property name', () => {
    const makerCondition: MakerCondition = TestBed.get(MakerCondition);
    makerCondition.name = 'testname';
    expect(makerCondition.name).toEqual('testname');
  });

  it('test toMap with full member', () => {
    const makerCondition: MakerCondition = TestBed.get(MakerCondition);
    makerCondition.id = 'testid';
    makerCondition.name = 'testname';
    const result: Map<string, string> = new Map();
    result.set('id', 'testid');
    result.set('name', 'testname');
    expect(makerCondition.toMap()).toEqual(result);
  });

  it('test toMap with undefind member', () => {
    const makerCondition: MakerCondition = TestBed.get(MakerCondition);
    makerCondition.id = 'testid';
    const result: Map<string, string> = new Map();
    result.set('id', 'testid');
    expect(makerCondition.toMap()).toEqual(result);
  });

});
