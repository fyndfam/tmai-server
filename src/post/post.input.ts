import * as jf from "joiful";
import { ApiProperty, ApiPropertyOptional } from "@nestjs/swagger";
import { objectId } from "../shared/types";

export class CreatePostInput {
  @jf.string().required()
  @ApiProperty()
  content: string;
}

export class PostIdParam {
  @objectId()
  postId: string;
}

export class GetPostsQuery {
  @jf.number().optional().min(0).max(50)
  @ApiPropertyOptional()
  limit?: number;

  @jf.number().optional().min(0)
  @ApiPropertyOptional()
  offset?: number;
}
