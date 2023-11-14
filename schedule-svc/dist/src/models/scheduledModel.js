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
exports.ScheduleDefModel = void 0;
const admin = __importStar(require("firebase-admin"));
class ScheduleDefModel {
    // SAVE SCHEDULE TO DB
    static async saveScheduleDef(thesis) {
        // Delete all documents in the table
        await this.deleteAllDocuments();
        // Add a new document
        const docRef = await this.scheduleDoc.add({
            ...thesis,
            createdAt: admin.firestore.Timestamp.now(),
        });
        const createdScheduleSnapshot = await docRef.get();
        const createdSchedule = {
            id: createdScheduleSnapshot.id,
            ...createdScheduleSnapshot.data(),
        };
        return createdSchedule;
    }
    static async deleteAllDocuments() {
        const querySnapshot = await this.scheduleDoc.get();
        const batch = admin.firestore().batch();
        querySnapshot.forEach((doc) => {
            batch.delete(doc.ref);
        });
        await batch.commit();
    }
    // GET SCHEDULE
    static async getSchedule() {
        const docRef = await this.scheduleDoc.get();
        const snapshot = docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
        return snapshot[0];
    }
    // GET ONE SCHEDULE
    static async getCouncilInSchedule(id) {
        const docRef = await this.getSchedule();
        const thesis = docRef.thesis.find((item) => item.id === id);
        if (!thesis) {
            return null;
        }
        return { ...thesis };
    }
    // GET SCHEDULE FOR STUDENT
    static async getScheduleForStudent(id) {
        const docRef = await this.getSchedule();
        const thesis = docRef.thesis.find((item) => item.schedule.timeSlots.find((timeSlot) => timeSlot.student.infor.id === id));
        const studentSchedule = thesis === null || thesis === void 0 ? void 0 : thesis.schedule.timeSlots.filter((item) => item.student.infor.id === id);
        if (!thesis) {
            return null;
        }
        return {
            ...thesis,
            schedule: {
                room: thesis.schedule.room,
                timeSlots: studentSchedule || [],
            },
        };
    }
    static async getScheduleForLecturer(id) {
        const docRef = await this.getSchedule();
        const thesis = docRef.thesis.filter((item) => item.council.find((item) => item.id === id));
        if (!thesis) {
            return null;
        }
        return thesis;
    }
}
exports.ScheduleDefModel = ScheduleDefModel;
_a = ScheduleDefModel;
ScheduleDefModel.db = admin.firestore();
ScheduleDefModel.scheduleDoc = _a.db.collection("schedule-defs");
