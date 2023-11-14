import * as admin from "firebase-admin";
import { INotificationObject } from "../interface/notification";

export class NotificationModel {
  private static db = admin.firestore();
  private static notificationDoc = this.db.collection("notifications");
  // CREATE MESSAGE
  static async createNotification(
    notification: Omit<INotificationObject, "id">
  ): Promise<string> {
    const docRef = await this.notificationDoc.add({
      ...notification,
      createdAt: admin.firestore.Timestamp.now(),
    });
    return docRef.id;
  }

  // GET ALL NOTIFICATION OF ONE USER
  static async getAllNotifications(
    id: string
  ): Promise<INotificationObject[]> {
    const docRef = await this.notificationDoc
      .where("receiverAuthor.id", "==", id)
      .get();
    return docRef.docs.map(
      (doc) => ({ id: doc.id, ...doc.data() } as INotificationObject)
    );
  }
}
