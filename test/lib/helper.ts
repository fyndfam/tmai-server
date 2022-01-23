import { connect, Connection, Mongoose } from "mongoose";
import { getMongoConnectionOptions } from "../../src/shared/db";

let mongoose: Mongoose;

export async function cleanDB(): Promise<void> {
  if (!mongoose) {
    await createConnection();
  }

  const collections = await mongoose.connection.db.listCollections().toArray();
  for (const collection of collections) {
    await mongoose.connection.db.collection(collection.name).deleteMany({});
  }
}

export async function closeConnection(): Promise<void> {
  if (mongoose) {
    await mongoose.disconnect();
    mongoose = null;
  }
}

export async function getConnection(): Promise<Connection> {
  if (!mongoose) {
    await createConnection();
  }

  return mongoose.connection;
}

async function createConnection(): Promise<void> {
  mongoose = await connect(process.env.MONGODB_URL, getMongoConnectionOptions());
}

export async function givenUser0(): Promise<string> {
  const connection = await getConnection();
  const createdUser = await connection.collection("users").insertOne({
    email: "test@fyndfam.com",
    name: "test test",
    alias: "test",
    avatar: "https://avatar.com/1",
    followingUsers: [],
    followingFam: [],
    followingFamBiz: [],
    notificationPreferences: {
      emailNotification: true,
      appNotification: true,
    },
    createdAt: new Date(),
    updatedAt: new Date(),
  });

  return createdUser.insertedId.toHexString();
}

export async function givenUser(): Promise<string> {
  const connection = await getConnection();
  const createdUser = await connection.collection("users").insertOne({
    email: "test_user@fyndfam.com",
    name: "Test Fynd",
    alias: "testfynd",
    avatar: "https://avatar.com/2",
    about: "About myself",
    country: "SG",
    followingUsers: [],
    followingFam: [],
    followingFamBiz: [],
    createdAt: new Date(),
    updatedAt: new Date(),
  });

  return createdUser.insertedId.toHexString();
}
