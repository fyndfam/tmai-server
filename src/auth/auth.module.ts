import { Module } from "@nestjs/common";
import { PassportModule } from "@nestjs/passport";

import { JwtStrategy } from "./jwt.strategy";
import { UserModule } from "../user/user.module";

@Module({
  imports: [PassportModule.register({ defaultStrategy: "jwt" }), UserModule],
  providers: [JwtStrategy],
  exports: [PassportModule],
})
export class AuthModule {}
