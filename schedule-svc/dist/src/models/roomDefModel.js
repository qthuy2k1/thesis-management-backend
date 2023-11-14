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
exports.RoomDefModel = void 0;
const admin = __importStar(require("firebase-admin"));
class RoomDefModel {
    // CREATE ROOM DEFENSE
    static async createRoomDef(room) {
        const existingUser = await this.roomDoc
            .where("name", "==", room.name)
            .get();
        if (!existingUser.empty) {
            throw new Error("RoomDef already exists");
        }
        else {
            const docRef = await this.roomDoc.add({
                ...room,
                createdAt: admin.firestore.Timestamp.now(),
            });
            const createdRoomSnapshot = await docRef.get();
            const createdRoom = {
                id: createdRoomSnapshot.id,
                ...createdRoomSnapshot.data(),
            };
            return createdRoom;
        }
    }
    // GET ONE ROOM DEFENSE
    static async getRoomDef(id) {
        const querySnapshot = await this.roomDoc.get();
        if (querySnapshot.empty) {
            return null;
        }
        const docRef = querySnapshot.docs[0];
        const data = docRef.data();
        return { ...data, id: docRef.id };
    }
    // GET ALL ROOM DEFENSES
    static async getAllRoomDef() {
        const docRef = await this.roomDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE ROOM DEFENSE
    static async updateRoomDef(room) {
        const { id, ...docRef } = room;
        await this.roomDoc.doc(id).update(docRef);
    }
    // DELETE ROOM DEFENSE
    static async deleteRoomDef(id) {
        await this.roomDoc.doc(id).delete();
        const snapshot = await this.roomDoc.where("id", "==", id).get();
        const rooms = [];
        snapshot.forEach((doc) => {
            rooms.push({ id: doc.id, ...doc.data() });
        });
        return rooms;
    }
}
exports.RoomDefModel = RoomDefModel;
_a = RoomDefModel;
RoomDefModel.db = admin.firestore();
RoomDefModel.roomDoc = _a.db.collection("room-defs");
