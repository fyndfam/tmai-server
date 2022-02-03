import { Body, Controller, Post, UseGuards, UsePipes } from "@nestjs/common";
import { AuthGuard } from "@nestjs/passport";
import { ApiResponse, ApiTags } from "@nestjs/swagger";
import { User as UserEntity } from "./user.entity";
import { UpdateUsernameInput } from "./user.input";
import { UserService } from "./user.service";
import { User } from "./user.decorator";
import { JoiValidationPipe } from "../shared/joi-validation.pipe";

@Controller("/users")
@ApiTags("users")
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Post("/username")
  @UseGuards(AuthGuard("jwt"))
  @UsePipes(new JoiValidationPipe({ requestBody: UpdateUsernameInput }))
  @ApiResponse({ status: 201, description: "username successfully update" })
  async updateUsername(
    @User() user: UserEntity,
    @Body() body: UpdateUsernameInput,
  ): Promise<{ status: "success" }> {
    const { username } = body;

    await this.userService.setUsername(user.email, username);

    return { status: "success" };
  }
}
