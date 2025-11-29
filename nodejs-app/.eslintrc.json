const request = require('supertest');
const server = require('./server');
describe('Server', () => {
  afterAll(() => server.close());
  it('Root path', async () => {
    const res = await request(server).get('/');
    expect(res.text).toEqual('Welcome to root web');
  });
  it('Sub path', async () => {
    const res = await request(server).get('/xyz');
    expect(res.text).toEqual('you access the xyz');
  });
});