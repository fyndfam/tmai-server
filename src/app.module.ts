import { MiddlewareConsumer, Module } from "@nestjs/common";
import { TerminusModule } from "@nestjs/terminus";
import { MongooseModule } from "@nestjs/mongoose";

import { getMongoConnectionOptions } from "./shared/db";
import { PostModule } from "./post/post.module";
import { RequestLoggingMiddleware } from "./shared/request-logging.middleware";

@Module({
  imports: [
    MongooseModule.forRoot(process.env.MONGODB_URL, getMongoConnectionOptions()),
    TerminusModule,
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
