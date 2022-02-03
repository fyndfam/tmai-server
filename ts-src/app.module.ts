import { MiddlewareConsumer, Module } from "@nestjs/common";
import { TerminusModule } from "@nestjs/terminus";
import { MongooseModule } from "@nestjs/mongoose";

import { getMongoConnectionOptions } from "./shared/db";
import { PostModule } from "./post/post.module";
import { RequestLoggingMiddleware } from "./shared/request-logging.middleware";
import { HealthCheckModule } from "./health-check/health-check.module";

@Module({
  imports: [
    MongooseModule.forRoot(process.env.MONGODB_URL, getMongoConnectionOptions()),
    TerminusModule,
    HealthCheckModule,
    PostModule,
  ],
  controllers: [],
  providers: [],
})
export class AppModule {
  configure(consumer: MiddlewareConsumer): void {
    consumer.apply(RequestLoggingMiddleware).forRoutes("*");
  }
}
