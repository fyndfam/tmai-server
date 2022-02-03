import { ConnectOptions } from "mongoose";

export function getMongoConnectionOptions(): ConnectOptions {
  return {
    autoIndex: true,
    autoCreate: true,
  };
}
