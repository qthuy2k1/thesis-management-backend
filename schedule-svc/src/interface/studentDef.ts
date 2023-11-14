import { IAuthObject } from "./auth";

export interface IStudentDef {
  id: string;
  infor: IAuthObject;
  instructor: IAuthObject;
}
