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
exports.ClassroomModel = void 0;
const admin = __importStar(require("firebase-admin"));
class ClassroomModel {
    // CREATE CLASSROOM
    static async createClassroom(classroom) {
        const existingUser = await this.classroomDoc
            .where("lecturer.id", "==", classroom.lecturer.id)
            .get();
        if (!existingUser.empty) {
            throw new Error("Classroom exists");
        }
        else {
            const docRef = await this.classroomDoc.add({
                ...classroom,
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
    static async getClassroom(id) {
        const querySnapshot = await this.classroomDoc
            .where("lecturer.id", "==", id)
            .get();
        if (querySnapshot.empty) {
            return null;
        }
        const docRef = querySnapshot.docs[0];
        const data = docRef.data();
        return { ...data, id: docRef.id };
    }
    // GET ALL CLASSROOM
    static async getAllClassroom() {
        const docRef = await this.classroomDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE CLASSROOM
    static async updateClassroom(classroom) {
        const { id, ...rest } = classroom;
        try {
            const batch = admin.firestore().batch();
            const querySnapshot = await this.memberDoc
                .where("classroom.id", "==", id)
                .get();
            querySnapshot.forEach((doc) => {
                const docRef = this.memberDoc.doc(doc.id);
                batch.update(docRef, { classroom: { id, ...rest } });
            });
            await batch.commit();
            await this.classroomDoc.doc(id).update({ id, ...rest });
        }
        catch (error) {
            console.error(error);
            throw new Error("Error to update sync classroom");
        }
    }
    // DELETE CLASSROOM
    static async deleteClassroom(id) {
        await this.classroomDoc.doc(id).delete();
        const snapshot = await this.classroomDoc.where("id", "==", id).get();
        const classrooms = [];
        snapshot.forEach((doc) => {
            classrooms.push({ id: doc.id, ...doc.data() });
        });
        return classrooms;
    }
}
exports.ClassroomModel = ClassroomModel;
_a = ClassroomModel;
ClassroomModel.db = admin.firestore();
ClassroomModel.classroomDoc = _a.db.collection("classrooms");
ClassroomModel.memberDoc = _a.db.collection("members");
