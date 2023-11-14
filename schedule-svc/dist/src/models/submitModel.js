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
exports.SubmitModel = void 0;
const admin = __importStar(require("firebase-admin"));
class SubmitModel {
    // CREATE SUBMIT
    static async createSubmit(submit) {
        const existingUser = await this.submitDoc
            .where("uid", "==", submit.uid)
            .get();
        if (!existingUser.empty) {
            throw new Error("Submit exists");
        }
        else {
            const docRef = await this.submitDoc.add({
                ...submit,
                createdAt: admin.firestore.Timestamp.now(),
            });
            const createdStageSnapshot = await docRef.get();
            const createdStage = {
                id: createdStageSnapshot.id,
                ...createdStageSnapshot.data(),
            };
            return createdStage;
        }
    }
    // GET ONE SUBMIT
    static async getSubmit(exerciseId, studentId) {
        const querySnapshot = await this.submitDoc
            .where("exerciseId", "==", exerciseId)
            .where("student.id", "==", studentId)
            .get();
        if (querySnapshot.empty) {
            return null;
        }
        const docSnapshot = querySnapshot.docs[0];
        const data = docSnapshot.data();
        return { ...data };
    }
    // GET ALL SUBMIT
    static async getAllSubmit(id) {
        const docRef = await this.submitDoc.where("exerciseId", "==", id).get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE SUBMIT
    static async updateSubmit(post) {
        const { id, ...docRef } = post;
        await this.submitDoc.doc(id).update(docRef);
    }
    // DELETE SUBMIT
    static async deleteSubmit(id) {
        await this.submitDoc.doc(id).delete();
        const snapshot = await this.submitDoc.where("id", "==", id).get();
        const posts = [];
        snapshot.forEach((doc) => {
            posts.push({ id: doc.id, ...doc.data() });
        });
        return posts;
    }
}
exports.SubmitModel = SubmitModel;
_a = SubmitModel;
SubmitModel.db = admin.firestore();
SubmitModel.submitDoc = _a.db.collection("submits");
