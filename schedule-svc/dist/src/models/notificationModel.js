"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.NotificationModel = void 0;
const admin = __importStar(require("firebase-admin"));
class NotificationModel {
    // CREATE MESSAGE
    static async createNotification(notification) {
        const docRef = await this.notificationDoc.add({
            ...notification,
            createdAt: admin.firestore.Timestamp.now(),
        });
        return docRef.id;
    }
    // GET ALL NOTIFICATION OF ONE USER
    static async getAllNotifications(id) {
        const docRef = await this.notificationDoc
            .where("receiverAuthor.id", "==", id)
            .get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
}
exports.NotificationModel = NotificationModel;
_a = NotificationModel;
NotificationModel.db = admin.firestore();
NotificationModel.notificationDoc = _a.db.collection("notifications");
