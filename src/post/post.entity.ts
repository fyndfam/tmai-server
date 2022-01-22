import { Prop, Schema, SchemaFactory } from "@nestjs/mongoose";
import { Document } from "mongoose";

export type PostDocument = Post & Document;

@Schema({ collection: "posts" })
export class Post {
  id: string;

  @Prop({ required: true })
  content: string;

  @Prop()
  tags: Array<string>;

  @Prop({ default: 0 })
  view: number;

  // username of the creator
  @Prop({ required: true })
  createdBy: string;

  @Prop({ default: false })
  contentEdited: boolean;

  @Prop()
  createdAt: Date;

  @Prop()
  updatedAt: Date;
}

export const PostSchema = SchemaFactory.createForClass(Post);

PostSchema.set("timestamps", true);
PostSchema.virtual("id").get(function () {
  return this._id.toHexString();
});
PostSchema.set("toObject", { virtuals: true });
PostSchema.method("toJSON", function () {
  const { __v, _id, ...object } = this.toObject();
  object.id = _id;
  return object;
});
PostSchema.index({ createdAt: 1 });
