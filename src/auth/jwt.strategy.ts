import { ExtractJwt, Strategy } from "passport-jwt";
import { passportJwtSecret } from "jwks-rsa";
import { Injectable } from "@nestjs/common";
import { PassportStrategy } from "@nestjs/passport";

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
  async validate(payload: any): Promise<any> {
    // TODO: get user by external user id
    // TODO: if no user is created, create user
  }
}
