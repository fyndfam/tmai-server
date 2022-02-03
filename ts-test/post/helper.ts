import { INestApplication } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";
import { Test } from "@nestjs/testing";

import { AuthModule } from "../../src/auth/auth.module";
import { PostController } from "../../src/post/post.controller";
import { PostService } from "../../src/post/post.service";
import { Post, PostSchema } from "../../src/post/post.entity";
import { getMongoConnectionOptions } from "../../src/shared/db";
import { getConnection } from "../lib/helper";

export async function createPostTestModule(): Promise<INestApplication> {
  const moduleRef = await Test.createTestingModule({
    imports: [
      MongooseModule.forRoot(process.env.MONGODB_URL, getMongoConnectionOptions()),
      MongooseModule.forFeature([{ name: Post.name, schema: PostSchema }]),
      AuthModule,
    ],
    providers: [PostService],
    controllers: [PostController],
  }).compile();

  return moduleRef.createNestApplication();
}

export async function givenPost() {
  const connection = await getConnection();
  const insertPostResult = await connection.collection("posts").insertOne({
    content: "This is a sample post",
    createdBy: "test",
    createdAt: "2022-01-05T01:00:00.000Z",
    updatedAt: "2022-01-05T01:00:00.000Z",
  });

  return insertPostResult.insertedId;
}

export async function givenPosts() {
  const connection = await getConnection();
  await connection.collection("posts").insertMany([
    {
      content: "Here is the first post",
      createdBy: "test",
      createdAt: "2022-01-05T01:00:00.000Z",
      updatedAt: "2022-01-05T01:00:00.000Z",
    },
    {
      content: "Second post",
      createdBy: "test",
      createdAt: "2022-01-06T01:00:00.000Z",
      updatedAt: "2022-01-06T01:00:00.000Z",
    },
    {
      content: "Things become more interesting",
      createdBy: "test",
      createdAt: "2022-01-05T08:00:00.000Z",
      updatedAt: "2022-01-05T08:00:00.000Z",
    },
  ]);
}
