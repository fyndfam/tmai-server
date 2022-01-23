import { Prop, Schema, SchemaFactory } from "@nestjs/mongoose";
import { Document } from "mongoose";

export type UserDocument = User & Document;

@Schema({ collection: "users" })
export class User {
  id: string;

  @Prop({ required: true })
  externalUserId: string;

  @Prop({ required: true })
  email: string;

  @Prop()
  username: string;

  @Prop()
  createdAt: Date;

  @Prop()
  updatedAt: Date;
}

export const UserSchema = SchemaFactory.createForClass(User);

UserSchema.set("timestamps", true);
UserSchema.virtual("id").get(function () {
  return this._id.toHexString();
});
UserSchema.set("toObject", { virtuals: true });
UserSchema.method("toJSON", function () {
  const { __v, _id, ...object } = this.toObject();
  object.id = _id;
  return object;
});
