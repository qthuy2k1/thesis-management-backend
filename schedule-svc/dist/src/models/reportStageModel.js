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
exports.ReportStageModel = void 0;
const admin = __importStar(require("firebase-admin"));
class ReportStageModel {
    // CREATE CREATE
    static async createReportStage(stage) {
        const existingUser = await this.reportStageDoc
            .where("value", "==", stage.value)
            .get();
        if (!existingUser.empty) {
            throw new Error("Report Stage exists");
        }
        else {
            const docRef = await this.reportStageDoc.add({
                ...stage,
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
    // GET ONE CREATE
    static async getReportStage(id) {
        const docRef = await this.reportStageDoc.doc(id).get();
        if (!docRef.exists) {
            return null;
        }
        const data = docRef.data();
        return { ...data };
    }
    // GET ALL CREATES
    static async getAllReportStage() {
        const docRef = await this.reportStageDoc.orderBy("createdAt", "asc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE CREATE
    static async updateReportStage(auth) {
        const { id, ...docRef } = auth;
        await this.reportStageDoc.doc(id).update(docRef);
    }
    // DELETE CREATE
    static async deleteReportStage(id) {
        await this.reportStageDoc.doc(id).delete();
        const snapshot = await this.reportStageDoc.where("id", "==", id).get();
        const auths = [];
        snapshot.forEach((doc) => {
            auths.push({ id: doc.id, ...doc.data() });
        });
        return auths;
    }
}
exports.ReportStageModel = ReportStageModel;
_a = ReportStageModel;
ReportStageModel.db = admin.firestore();
ReportStageModel.reportStageDoc = _a.db.collection("report-stage");
