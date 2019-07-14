import * as queries from './queries';

const agents = [
  {
    id: '90346453',
    name: 'Livingroom',
    data: {},
    type: 'type1-v3.3.0',
    ip: '192.168.1.10',
    isOnline: true,
  },
  {
    id: '43366411',
    name: 'Badroom',
    data: {},
    type: 'jeep-v0.2.0',
    ip: '192.168.1.17',
    isOnline: false,
  },
];

describe('models/agents/queries', () => {
  describe('isOnline', () => {
    it('should return whether device is on/off line', () => {
      const isOnline1 = queries.isOnline(agents[0]);
      expect(isOnline1).toBe(true);

      const isOnline2 = queries.isOnline(agents[1]);
      expect(isOnline2).toBe(false);
    });
  });

  describe('getNoVersionedType', () => {
    it('should return raw device type (no version)', () => {
      const devType1 = queries.getNoVersionedType(agents[0]);
      expect(devType1).toEqual('type1');

      const devType2 = queries.getNoVersionedType(agents[1]);
      expect(devType2).toEqual('jeep');
    });
  });
});
