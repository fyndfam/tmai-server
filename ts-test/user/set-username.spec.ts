import { INestApplication } from "@nestjs/common";
import { expect } from "chai";
import request from "supertest";
import { getConnection, givenUser0 } from "../lib/helper";
import { JWT } from "../lib/constants";
import { createUserTestModule } from "./helper";

describe("POST /users/username", () => {
  let app: INestApplication;

  before(async () => {
    app = await createUserTestModule();
    await app.init();
  });

  it("should not be able to set username again", async () => {
    await givenUser0();

    await request(app.getHttpServer())
      .post("/users/username")
      .set("Authorization", `bearer ${JWT}`)
      .send({ username: "test2" })
      .expect(403);
  });

  it("should be able to set username for the first time", async () => {
    await request(app.getHttpServer())
      .post("/users/username")
      .set("Authorization", `bearer ${JWT}`)
      .send({ username: "test2" })
      .expect(201);

    const connection = await getConnection();
    const user = await connection.collection("users").findOne({});
    expect(user.username).to.equal("test2");
  });
});
