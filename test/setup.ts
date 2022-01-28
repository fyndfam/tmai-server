import { cleanDB, closeConnection } from "./lib/helper";

after(async () => {
  await closeConnection();
});

beforeEach(async () => {
  await cleanDB();
});
