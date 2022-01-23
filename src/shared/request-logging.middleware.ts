import { Injectable, Logger, NestMiddleware } from "@nestjs/common";
import { NextFunction, Request, Response } from "express";

@Injectable()
export class RequestLoggingMiddleware implements NestMiddleware {
  private logger = new Logger("HTTP");

  use(req: Request, res: Response, next: NextFunction): void {
    const { ip, method, originalUrl: url } = req;
    const userAgent = req.get("user-agent") ?? "";
    const forwarded = req.get("X-Forwarded-For") ?? "";

    res.on("close", () => {
      const { statusCode } = res;
      const contentLength = res.get("content-length");

      this.logger.log(
        `${ip} ${userAgent} - ${method} ${url} ${statusCode} ${contentLength} - ${forwarded}`,
      );
    });

    next();
  }
}
