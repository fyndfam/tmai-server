import {
  Body,
  Controller,
  Get,
  NotFoundException,
  Param,
  Post,
  Query,
  UsePipes,
  UseGuards,
  HttpException,
  HttpStatus,
} from "@nestjs/common";
import { AuthGuard } from "@nestjs/passport";
import { ApiResponse, ApiTags } from "@nestjs/swagger";
import { Post as PostEntity } from "./post.entity";
import { PostService } from "./post.service";
import { JoiValidationPipe } from "../shared/joi-validation.pipe";
import { CreatePostInput, GetPostsQuery, PostIdParam } from "./post.input";
import { PostDetail } from "./post.output";
import { User as UserEntity } from "../user/user.entity";
import { User } from "../user/user.decorator";

@Controller("/posts")
@ApiTags("posts")
export class PostController {
  constructor(private readonly postService: PostService) {}

  @Post("/")
  @UseGuards(AuthGuard("jwt"))
  @UsePipes(new JoiValidationPipe({ requestBody: CreatePostInput }))
  @ApiResponse({ status: 201, description: "post created" })
  @ApiResponse({ status: 400, description: "bad request" })
  async createPost(@User() user: UserEntity, @Body() body: CreatePostInput): Promise<PostDetail> {
    if (!user.username) {
      throw new HttpException(
        {
          status: HttpStatus.BAD_REQUEST,
          error: "username is not set",
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    return this.postService.createPost(user.username, body);
  }

  @Get("/:postId")
  @UsePipes(new JoiValidationPipe({ param: PostIdParam }))
  @ApiResponse({ status: 200, description: "success" })
  @ApiResponse({ status: 404, description: "post not found" })
  async getPostById(@Param() params: PostIdParam): Promise<PostEntity> {
    const { postId } = params;
    const post = await this.postService.getPostById(postId);

    if (!post) {
      throw new NotFoundException();
    }

    return post;
  }

  @Get("/")
  @UsePipes(new JoiValidationPipe({ query: GetPostsQuery }))
  @ApiResponse({ status: 200, description: "success" })
  async getPosts(@Query() queries: GetPostsQuery): Promise<Array<PostEntity>> {
    const limit = queries.limit ? Number(queries.limit) : 50;
    const offset = queries.limit ? Number(queries.offset) : 0;

    return this.postService.getLatestPosts(limit, offset);
  }
}
