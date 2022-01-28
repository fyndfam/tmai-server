import { ExtractJwt, Strategy } from "passport-jwt";
import { passportJwtSecret } from "jwks-rsa";
import { Injectable } from "@nestjs/common";
import { PassportStrategy } from "@nestjs/passport";

import { UserService } from "../user/user.service";
import { User } from "..//user/user.entity";

const configurations: any = {
  jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
};

if (process.env.NODE_ENV === "development") {
  configurations.algorithm = ["HS256"];
  configurations.secretOrKey = "BSDGR3VVE3EHMTVEYRMTKSUB";
  configurations.audience = "tmaiserver";
  configurations.issuer = "tmaiserver";
} else {
  configurations.algorithms = ["RS256"];
  configurations.secretOrKeyProvider = passportJwtSecret({
    cache: true,
    rateLimit: true,
    jwksRequestsPerMinute: 5,
    jwksUri: process.env.JWKS_URI,
  });
  configurations.audience = process.env.JWT_AUDIENCE;
  configurations.issuer = process.env.JWT_ISSUER;
}

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor(private readonly userService: UserService) {
    super(configurations);
  }

  async validate(payload: any): Promise<User> {
    const { email } = payload;

    const user = await this.userService.getUserByEmail(email);
    if (!user) {
      return this.userService.createUser(email);
    }

    return user;
  }
}
