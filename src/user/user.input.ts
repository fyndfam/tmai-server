import * as jf from "joiful";

export class UpdateUsernameInput {
  @jf
    .string()
    .regex(/^[a-z][0-9a-zA-Z]{2,19}$/)
    .required()
  username: string;
}
