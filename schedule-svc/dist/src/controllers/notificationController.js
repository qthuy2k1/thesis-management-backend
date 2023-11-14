"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.notificationController = void 0;
const notificationModel_1 = require("../models/notificationModel");
exports.notificationController = {
    async createNotification(req) {
        const notification = req.request.noti;
        try {
            await notificationModel_1.NotificationModel.createNotification(notification);
            const uid = notification.receiverAuthor.uid;
            const notifications = await notificationModel_1.NotificationModel.getAllNotifications(uid);
            // res.status(200).json({
            let res = {
                notification,
                message: "Notification has been created",
                notifications: notifications,
            };
            // });
            return res;
        }
        catch (error) {
            console.log(error);
            // res.sendStatus(500);
        }
    },
    async getAllNotification(req, res) {
        try {
            const uid = req.params.id;
            const notifications = await notificationModel_1.NotificationModel.getAllNotifications(uid);
            if (!notifications) {
                res.sendStatus(404).json("Notification is empty");
                return;
            }
            res.status(200).json(notifications);
        }
        catch (error) { }
    },
};
