import { INestApplication } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";
import { Test } from "@nestjs/testing";

import { AuthModule } from "../../src/auth/auth.module";
import { PostController } from "../../src/post/post.controller";
import { PostService } from "../../src/post/post.service";
import { Post, PostSchema } from "../../src/post/post.entity";
import { getMongoConnectionOptions } from "../../src/shared/db";

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
