import { TRUE, FALSE } from './boolean';

describe('Boolean', () => {
  it('should to string true', () => {
    expect(TRUE).toEqual('true');
  });

  it('should to string false', () => {
    expect(FALSE).toEqual('false');
  });
});
