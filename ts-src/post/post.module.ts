import { Module } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";

import { AuthModule } from "../auth/auth.module";
import { PostController } from "./post.controller";
import { PostService } from "./post.service";
import { Post, PostSchema } from "./post.entity";

@Module({
  imports: [MongooseModule.forFeature([{ name: Post.name, schema: PostSchema }]), AuthModule],
  providers: [PostService],
  controllers: [PostController],
})
export class PostModule {}
