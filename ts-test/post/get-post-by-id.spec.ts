import { INestApplication } from "@nestjs/common";
import { expect } from "chai";
import request from "supertest";
import { givenUser0 } from "../lib/helper";
import { createPostTestModule, givenPost } from "./helper";

describe("GET /posts/post_id", () => {
  let app: INestApplication;

  before(async () => {
    app = await createPostTestModule();
    await app.init();
  });

  it("should be able to get posts", async () => {
    await givenUser0();
    const postId = await givenPost();

    await request(app.getHttpServer())
      .get(`/posts/${postId}`)
      .expect(200)
      .then((response) => {
        expect(response.body.content).to.equal("This is a sample post");
        expect(response.body.createdBy).to.equal("test");
      });
  });
});
