import { Module } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";
import { TerminusModule } from "@nestjs/terminus";
import { HealthCheckController } from "./health-check.controller";

@Module({
  imports: [TerminusModule, MongooseModule.forRoot(process.env.MONGODB_URL)],
  controllers: [HealthCheckController],
})
export class HealthCheckModule {}
