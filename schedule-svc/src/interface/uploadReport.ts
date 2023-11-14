import { IGeneralLinkAttachment } from "./submit";
import { IAuthObject } from "./auth";

export interface IUploadReportObject {
  id: string;
  uid: string;
  attachments: File[] | IGeneralLinkAttachment[];
  student: IAuthObject;
  status: string;
}
