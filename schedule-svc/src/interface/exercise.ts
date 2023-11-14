import { IClassroomObject } from "./classroom";
import { IReportStageObject } from "./report-stage";
import { IGeneralLinkAttachment } from "./submit";
import { IAuthObject } from "./auth";

export interface IExerciseObject {
  id: string;
  uid: string;
  title: string;
  classroom: IClassroomObject;
  category: IReportStageObject;
  lecturer: IAuthObject;
  description: string;
  deadline: Date;
  type: string;
  attachments?: File[] | IGeneralLinkAttachment[];
}
