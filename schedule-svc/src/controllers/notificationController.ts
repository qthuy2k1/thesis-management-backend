import { Request, Response } from "express";
import { NotificationModel } from "../models/notificationModel";

export const notificationController = {
  async createNotification(req: any): Promise<any> {
    const notification = req.request.noti;
    try {
      await NotificationModel.createNotification(notification);

      const uid = notification.receiverAuthor.uid;
      const notifications = await NotificationModel.getAllNotifications(uid);

      // res.status(200).json({
      let res = {
        notification,
        message: "Notification has been created",
        notifications: notifications,
      }
      // });

      return res;
    } catch (error) {
      console.log(error);
      // res.sendStatus(500);
    }
  },

  async getAllNotification(req: Request, res: Response) {
    try {
      const uid = req.params.id;
      const notifications = await NotificationModel.getAllNotifications(uid);
      if (!notifications) {
        res.sendStatus(404).json("Notification is empty");
        return;
      }
      res.status(200).json(notifications);
    } catch (error) {}
  },
};