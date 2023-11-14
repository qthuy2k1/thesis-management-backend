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
exports.TopicModel = void 0;
const admin = __importStar(require("firebase-admin"));
class TopicModel {
    // CREATE TOPIC
    static async createTopic(topic) {
        const existingUser = await this.topicDoc.doc(topic.student.id).get();
        if (existingUser.exists) {
            throw new Error("Topic exists");
        }
        else {
            const newTopic = {
                ...topic,
                createdAt: admin.firestore.Timestamp.now(),
            };
            await this.topicDoc.doc(topic.student.id).set(newTopic);
            return newTopic;
        }
    }
    // GET ONE TOPIC
    static async getTopic(studentId) {
        const docRef = await this.topicDoc
            .where("student.id", "==", studentId)
            .get();
        if (docRef.empty) {
            return null;
        }
        const data = docRef.docs[0].data();
        return { ...data, id: studentId };
    }
    // GET ALL TOPIC
    static async getAllTopic() {
        const docRef = await this.topicDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE TOPIC
    static async updateTopic(topic) {
        const { id, ...docRef } = topic;
        await this.topicDoc.doc(id).update(docRef);
    }
    // DELETE TOPIC
    static async deleteTopic(id) {
        await this.topicDoc.doc(id).delete();
        const snapshot = await this.topicDoc.where("id", "==", id).get();
        const topics = [];
        snapshot.forEach((doc) => {
            topics.push({ id: doc.id, ...doc.data() });
        });
        return topics;
    }
}
exports.TopicModel = TopicModel;
_a = TopicModel;
TopicModel.db = admin.firestore();
TopicModel.topicDoc = _a.db.collection("topics");
