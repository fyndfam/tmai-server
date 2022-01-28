import { ForbiddenException, Injectable, Logger, NotFoundException } from "@nestjs/common";
import { InjectModel } from "@nestjs/mongoose";
import { Model, Types } from "mongoose";
import { Post, PostDocument } from "./post.entity";
import { PostDetail } from "./post.output";

const { ObjectId } = Types;

@Injectable()
export class PostService {
  private readonly logger = new Logger("PostService");

  constructor(
    @InjectModel(Post.name)
    private readonly postModel: Model<PostDocument>,
  ) {}

  async getPostById(id: string): Promise<Post | null> {
    this.logger.log(`get post by id: ${id}`);

    return this.postModel.findOne({ _id: new ObjectId(id) });
  }

  async getLatestPosts(limit: number, offset: number): Promise<Array<Post>> {
    this.logger.log(`get latest post with limit: ${limit} and offset: ${offset}`);

    return this.postModel.find({}).sort({ createdAt: -1 }).skip(offset).limit(limit);
  }

  async createPost(username: string, data: any): Promise<PostDetail> {
    this.logger.log(`user ${username} creating a post`);

    const post = new this.postModel({
      content: data.content,
      createdBy: username,
    });

    const createdPost = await post.save();
    return {
      id: createdPost._id,
      content: createdPost.content,
      contentEdited: false,
      createdBy: createdPost.createdBy,
      createdAt: createdPost.createdAt.toISOString(),
      updatedAt: createdPost.updatedAt.toISOString(),
    };
  }

  async updatePost(username: string, postId: string, content: string): Promise<Post> {
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
          contentEdited: true,
        },
      },
      {
        new: true,
      },
    );
  }
}
