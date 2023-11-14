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
exports.UploadReportModel = void 0;
const admin = __importStar(require("firebase-admin"));
class UploadReportModel {
    // CREATE UPLOAD REPORT
    static async createUploadReport(upload) {
        const existingUploadReport = await this.uploadDoc
            .where("uid", "==", upload.uid)
            .get();
        if (!existingUploadReport.empty) {
            throw new Error("UploadReport exists");
        }
        else {
            const docRef = await this.uploadDoc.add({
                ...upload,
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
    // GET ONE UPLOAD REPORT
    static async getUploadReport(id) {
        const querySnapshot = await this.uploadDoc
            .where("student.id", "==", id)
            .get();
        if (querySnapshot.empty) {
            return null;
        }
        const docRef = querySnapshot.docs[0];
        const data = docRef.data();
        return { ...data, id: docRef.id };
    }
    // GET ALL UPLOAD REPORT
    static async getAllUploadReport() {
        const docRef = await this.uploadDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE UPLOAD REPORT
    static async updateUploadReport(upload) {
        const { id, ...docRef } = upload;
        await this.uploadDoc.doc(id).update(docRef);
    }
    // DELETE UPLOAD REPORT
    static async deleteUploadReport(id) {
        await this.uploadDoc.doc(id).delete();
        const snapshot = await this.uploadDoc.where("id", "==", id).get();
        const uploads = [];
        snapshot.forEach((doc) => {
            uploads.push({ id: doc.id, ...doc.data() });
        });
        return uploads;
    }
}
exports.UploadReportModel = UploadReportModel;
_a = UploadReportModel;
UploadReportModel.db = admin.firestore();
UploadReportModel.uploadDoc = _a.db.collection("uploads");
