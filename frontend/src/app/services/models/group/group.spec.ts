import { Group } from './group';

export function createTestInstance1(): Group {
  return {
    id: 'test-group-id-1',
    name: 'TEST_GROUP_NAME_1,'
  }
}

export function createTestInstance2(): Group {
  return {
    id: 'test-group-id-2',
    name: 'TEST_GROUP_NAME_2,'
  }
}

export function ceateTestArray(): Group[] {
  return [
      createTestInstance1(),
      createTestInstance2(),
  ];
}