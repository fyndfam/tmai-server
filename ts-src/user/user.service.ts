import { HttpException, HttpStatus, Injectable, Logger, NotFoundException } from "@nestjs/common";
import { InjectModel } from "@nestjs/mongoose";
import { Model } from "mongoose";
import { User, UserDocument } from "./user.entity";

@Injectable()
export class UserService {
  private readonly logger = new Logger("UserService");

  constructor(
    @InjectModel(User.name)
    private readonly userModel: Model<UserDocument>,
  ) {}

  async getUserByEmail(email: string): Promise<User | null> {
    this.logger.log(`get user with email: ${email}`);

    return this.userModel.findOne({ email });
  }

  async doesUsernameExists(username: string): Promise<boolean> {
    this.logger.log(`checking if username: ${username} exists`);

    const userWithUsername = await this.userModel.findOne({ username });
    return !!userWithUsername;
  }

  async createUser(email: string): Promise<User> {
    this.logger.log(`create user with email: ${email}`);

    const user = new this.userModel({
      email,
    });
    return user.save();
  }

  async setUsername(email: string, username: string): Promise<void> {
    this.logger.log(`set user of email address: ${email} with username: ${username}`);

    const user = await this.getUserByEmail(email);
    if (!user) {
      throw new NotFoundException();
    }

    if (user.username) {
      throw new HttpException(
        {
          status: HttpStatus.FORBIDDEN,
          error: "Can not change username",
        },
        HttpStatus.FORBIDDEN,
      );
    }

    const usernameExists = await this.doesUsernameExists(username);
    if (usernameExists) {
      throw new HttpException(
        {
          status: HttpStatus.BAD_REQUEST,
          error: "Username not available",
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    await this.userModel.updateOne({ email }, { $set: { username } });
  }
}
