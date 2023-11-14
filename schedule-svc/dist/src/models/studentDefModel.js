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
exports.StudentDefModel = void 0;
const admin = __importStar(require("firebase-admin"));
class StudentDefModel {
    // CREATE STUDENT DEFENSE
    static async createStudentDef(auth) {
        const existingUser = await this.authDoc
            .where("id", "==", auth.infor.id)
            .get();
        if (!existingUser.empty) {
            throw new Error("auth exists");
        }
        else {
            const docRef = await this.authDoc.add({
                ...auth,
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
    // GET ONE STUDENT DEFENSE
    static async getStudentDef(id) {
        const docRef = await this.authDoc.doc(id).get();
        if (!docRef.exists) {
            return null;
        }
        const data = docRef.data();
        return { ...data };
    }
    // GET ALL STUDENT DEFENSES
    static async getAllStudentDef() {
        const docRef = await this.authDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    static async getAllStudentDefPag(page, limit) {
        let query = this.authDoc.limit(limit);
        if (page > 1) {
            const skip = (page - 1) * limit;
            query = query.offset(skip);
        }
        const docRef = await query.get();
        const studefs = [];
        docRef.forEach((doc) => {
            studefs.push({ id: doc.id, ...doc.data() });
        });
        return studefs;
    }
    // UPDATE STUDENT DEFENSE
    static async updateStudentDef(auth) {
        const { id, ...docRef } = auth;
        await this.authDoc.doc(id).update(docRef);
    }
    // DELETE STUDENT DEFENSE
    static async deleteStudentDef(id) {
        await this.authDoc.doc(id).delete();
        const snapshot = await this.authDoc.where("id", "==", id).get();
        const auths = [];
        snapshot.forEach((doc) => {
            auths.push({ id: doc.id, ...doc.data() });
        });
        return auths;
    }
    static async deleteAllDocuments() {
        const querySnapshot = await this.authDoc.get();
        const batch = admin.firestore().batch();
        querySnapshot.forEach((doc) => {
            batch.delete(doc.ref);
        });
        await batch.commit();
    }
}
exports.StudentDefModel = StudentDefModel;
_a = StudentDefModel;
StudentDefModel.db = admin.firestore();
StudentDefModel.authDoc = _a.db.collection("student-defs");
