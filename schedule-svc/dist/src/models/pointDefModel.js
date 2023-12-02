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
exports.PointDefModel = void 0;
const admin = __importStar(require("firebase-admin"));
class PointDefModel {
    // CREATE POINT DEFENSE
    static async createPointDef(point) {
        const existingUser = await this.pointDoc
            .where("student.id", "==", point.student.id)
            .get();
        if (!existingUser.empty) {
            throw new Error("PointDef exists");
        }
        else {
            const docRef = await this.pointDoc.add({
                ...point,
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
    // UPDATE POINT DEFENSE
    static async updatePointDef(point) {
        var _b;
        const { id, assesses, ...docRef } = point;
        const existingPointDef = await this.pointDoc.doc(id).get();
        if (!existingPointDef.exists) {
            throw new Error("PointDef không tồn tại");
        }
        const existingAssesses = ((_b = existingPointDef.data()) === null || _b === void 0 ? void 0 : _b.assesses) || [];
        const updatedAssesses = [
            ...existingAssesses.filter((item) => item.lecturer.id !== assesses[0].lecturer.id),
            ...assesses,
        ];
        await this.pointDoc
            .doc(id)
            .update({ ...docRef, assesses: updatedAssesses });
    }
    static async createOrUpdatePointDef(point) {
        const studentId = point.student.id;
        const existingPointDef = await this.pointDoc
            .where("student.id", "==", studentId)
            .get();
        if (!existingPointDef.empty) {
            const pointDefId = existingPointDef.docs[0].id;
            await this.updatePointDef({ ...point, id: pointDefId });
            return { ...point, id: pointDefId };
        }
        else {
            return await this.createPointDef(point);
        }
    }
    // GET ONE POINT DEFENSE
    static async getPointDef(id) {
        const querySnapshot = await this.pointDoc
            .where("student.id", "==", id)
            .get();
        if (querySnapshot.empty) {
            return null;
        }
        const docRef = querySnapshot.docs[0];
        const data = docRef.data();
        return { ...data, id: docRef.id };
    }
    static async getPointDefForLecturer(studefId, lecId) {
        const querySnapshot = await this.pointDoc
            .where("student.id", "==", studefId)
            .get();
        if (querySnapshot.empty) {
            return null;
        }
        const docRef = querySnapshot.docs[0];
        const data = docRef.data();
        const objectWithLecId = data.assesses.find((obj) => obj.lecturer.id === lecId);
        if (objectWithLecId) {
            return { ...objectWithLecId };
        }
        return null;
    }
    // GET ALL POINT DEFENSE
    static async getAllPointDef(id) {
        const docRef = await this.pointDoc.get();
        return docRef.docs
            .map((doc) => ({ id: doc.id, ...doc.data() }))
            .filter((pointDef) => pointDef.assesses.some((assess) => assess.lecturer.id === id));
    }
    // DELETE POINT DEFENSE
    static async deletePointDef(id) {
        await this.pointDoc.doc(id).delete();
        const snapshot = await this.pointDoc.where("id", "==", id).get();
        const points = [];
        snapshot.forEach((doc) => {
            points.push({ id: doc.id, ...doc.data() });
        });
        return points;
    }
}
exports.PointDefModel = PointDefModel;
_a = PointDefModel;
PointDefModel.db = admin.firestore();
PointDefModel.pointDoc = _a.db.collection("point-def");
