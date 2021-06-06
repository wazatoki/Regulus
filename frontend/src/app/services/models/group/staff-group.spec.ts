import { StaffGroup } from './staff-group';

export function createTestInstance1(): StaffGroup {
  return {
    id: 'test-group-id-1',
    name: 'TEST_GROUP_NAME_1,'
  }
}

export function createTestInstance2(): StaffGroup {
  return {
    id: 'test-group-id-2',
    name: 'TEST_GROUP_NAME_2,'
  }
}

export function ceateTestArray(): StaffGroup[] {
  return [
      createTestInstance1(),
      createTestInstance2(),
  ];
}