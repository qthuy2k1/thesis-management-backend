import { IExerciseObject } from "./exercise";
import { IAuthObject } from "./auth";

export interface IGeneralLinkAttachment {
  id: string;
  name: string;
  src: string;
}

export interface ISubmitObject {
  id: string;
  uid: string;
  student: IAuthObject;
  exerciseId: IExerciseObject;
  attachments: IGeneralLinkAttachment[];
  status: string;
}
