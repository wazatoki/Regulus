import { Staff } from './staff';

export function createTestInstance1(): Staff {
  return {
    id: 'test-group-id-1',
    name: 'TEST_GROUP_NAME_1,'
  }
}

export function createTestInstance2(): Staff {
  return {
    id: 'test-group-id-2',
    name: 'TEST_GROUP_NAME_2,'
  }
}

export function ceateTestArray(): Staff[] {
  return [
      createTestInstance1(),
      createTestInstance2(),
  ];
}