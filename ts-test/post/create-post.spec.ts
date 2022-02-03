import { INestApplication } from "@nestjs/common";
import request from "supertest";
import { expect } from "chai";

import { createPostTestModule } from "./helper";
import { getConnection, givenUser0 } from "../lib/helper";
import { JWT } from "../lib/constants";

describe("POST /posts", () => {
  let app: INestApplication;

  before(async () => {
    app = await createPostTestModule();
    await app.init();
  });

  it("should be able to create post", async () => {
    await givenUser0();
    const postContent = "hello, this is my first post";

    await request(app.getHttpServer())
      .post("/posts")
      .set("Authorization", `bearer ${JWT}`)
      .send({
        content: postContent,
      })
      .expect(201);

    const connection = await getConnection();
    const post = await connection.collection("posts").findOne({});
    expect(post.content).to.equal(postContent);
    expect(post.createdBy).to.equal("test");
  });

  it("should throw 400 if username is not set before creating a post", async () => {
    const postContent = "hello, this is my first post";

    await request(app.getHttpServer())
      .post("/posts")
      .set("Authorization", `bearer ${JWT}`)
      .send({
        content: postContent,
      })
      .expect(400);
  });

  it("should not be able to create post if not authorized", async () => {
    await request(app.getHttpServer())
      .post("/posts")
      .send({
        content: "hello, i just want to create a post",
      })
      .expect(401);
  });
});
