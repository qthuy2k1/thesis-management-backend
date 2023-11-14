import { IAuthObject } from "./auth";

export interface INotificationObject {
  id: string;
  senderUser: IAuthObject;
  receiverAuthor: IAuthObject;
  type: string;
}
