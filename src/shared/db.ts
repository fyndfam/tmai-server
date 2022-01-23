import { ConnectOptions } from "mongoose";

export function getMongoConnectionOptions(): ConnectOptions {
  return {
    useFindAndModify: false,
    useNewUrlParser: true,
    useCreateIndex: true,
  };
}
