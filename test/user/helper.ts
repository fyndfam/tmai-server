import { INestApplication } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";
import { Test } from "@nestjs/testing";

import { AuthModule } from "../../src/auth/auth.module";
import { getMongoConnectionOptions } from "../../src/shared/db";
import { UserController } from "../../src/user/user.controller";
import { User, UserSchema } from "../../src/user/user.entity";
import { UserService } from "../../src/user/user.service";

export async function createUserTestModule(): Promise<INestApplication> {
  const moduleRef = await Test.createTestingModule({
    imports: [
      MongooseModule.forRoot(process.env.MONGODB_URL, getMongoConnectionOptions()),
      MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
      AuthModule,
    ],
    providers: [UserService],
    controllers: [UserController],
  }).compile();

  return moduleRef.createNestApplication();
}
