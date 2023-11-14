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
exports.ExerciseModel = void 0;
const admin = __importStar(require("firebase-admin"));
class ExerciseModel {
    // CREATE EXERCISE
    static async createExercise(exercise) {
        const existingExercise = await this.exerciseDoc
            .where("uid", "==", exercise.uid)
            .get();
        if (!existingExercise.empty) {
            throw new Error("Exercise exists");
        }
        else {
            const docRef = await this.exerciseDoc.add({
                ...exercise,
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
    // GET ONE EXERCISE
    static async getExercise(id) {
        const docRef = await this.exerciseDoc.doc(id).get();
        if (!docRef.exists) {
            return null;
        }
        const data = docRef.data();
        return { ...data };
    }
    // GET ALL EXERCISE
    static async getAllExercise() {
        const docRef = await this.exerciseDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // GET ALL EXERCISE IN CLASS
    static async getAllExerciseInClass(id) {
        const docRef = await this.exerciseDoc.where("classroom.id", "==", id).get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // GET ALL EXERCISE IN REPORT STAGE
    static async getAllExerciseInReportStage(classroomId, categoryId) {
        const docRef = await this.exerciseDoc
            .where("classroom.id", "==", classroomId)
            .where("category.id", "==", categoryId)
            .get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE EXERCISE
    static async updateExercise(exercise) {
        const { id, ...docRef } = exercise;
        await this.exerciseDoc.doc(id).update(docRef);
    }
    // DELETE EXERCISE
    static async deleteExercise(id) {
        await this.exerciseDoc.doc(id).delete();
        const snapshot = await this.exerciseDoc.where("id", "==", id).get();
        const exercises = [];
        snapshot.forEach((doc) => {
            exercises.push({ id: doc.id, ...doc.data() });
        });
        return exercises;
    }
}
exports.ExerciseModel = ExerciseModel;
_a = ExerciseModel;
ExerciseModel.db = admin.firestore();
ExerciseModel.exerciseDoc = _a.db.collection("exercises");
