import { INestApplication } from "@nestjs/common";
import { expect } from "chai";
import request from "supertest";
import { givenUser0 } from "../lib/helper";
import { createPostTestModule, givenPosts } from "./helper";

describe("GET /posts", () => {
  let app: INestApplication;

  before(async () => {
    app = await createPostTestModule();
    await app.init();
  });

  it("should be able to get posts", async () => {
    await givenUser0();
    await givenPosts();

    await request(app.getHttpServer())
      .get("/posts")
      .expect(200)
      .then((response) => {
        expect(response.body.length).to.equal(3);
      });
  });
});
