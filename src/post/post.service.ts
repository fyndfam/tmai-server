import { ForbiddenException, Injectable, Logger, NotFoundException } from "@nestjs/common";
import { InjectModel } from "@nestjs/mongoose";
import { Model, Types } from "mongoose";
import { Post, PostDocument } from "./post.entity";

const { ObjectId } = Types;

@Injectable()
export class PostService {
  private readonly logger = new Logger("PostService");

  constructor(
    @InjectModel(Post.name)
    private readonly postModel: Model<PostDocument>,
  ) {}

  async getPostById(id: string) {
    this.logger.log(`get post by id: ${id}`);

    return this.postModel.findOne({ _id: new ObjectId(id) });
  }

  async getLatestPosts(limit: number, offset: number) {
    this.logger.log(`get latest post with limit: ${limit} and offset: ${offset}`);

    return this.postModel.find({}).sort({ createdAt: -1 }).skip(offset).limit(limit);
  }

  async createPost(username: string, data: any) {
    this.logger.log(`user ${username} creating a post`);

    const post = new this.postModel({
      content: data.content,
      createdBy: username,
    });

    return post.save();
  }

  async updatePost(username: string, postId: string, content: string) {
    this.logger.log(`update post by user: ${username} for post with id ${postId}`);

    if (!content) {
      return;
    }

    const post = await this.getPostById(postId);
    if (!post) {
      throw new NotFoundException();
    }
    if (post.createdBy !== username) {
      throw new ForbiddenException();
    }

    return this.postModel.findOneAndUpdate(
      {
        _id: new ObjectId(postId),
      },
      {
        $set: {
          content,
        },
      },
      {
        new: true,
      },
    );
  }
}
