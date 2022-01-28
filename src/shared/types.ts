import * as jf from "joiful";

export const objectId = (): any =>
  jf
    .string()
    .regex(/^[0-9a-fA-F]{24}$/)
    .required();
