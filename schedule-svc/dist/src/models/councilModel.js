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
exports.CouncilModel = void 0;
const admin = __importStar(require("firebase-admin"));
class CouncilModel {
    // CREATE COUNCIL
    static async createCouncil(user) {
        const existingUser = await this.councilDoc.doc(user.id).get();
        if (existingUser.exists) {
            throw new Error("Council exists");
        }
        else {
            const newUser = {
                ...user,
                createdAt: admin.firestore.Timestamp.now(),
            };
            await this.councilDoc.doc(user.id).set(newUser);
            return newUser;
        }
    }
    // GET ONE COUNCIL
    static async getCouncil(id) {
        const docRef = await this.councilDoc.doc(id).get();
        if (!docRef.exists) {
            return null;
        }
        const data = docRef.data();
        return { ...data };
    }
    // GET ALL COUNCILS
    static async getAllCouncil() {
        const docRef = await this.councilDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => doc.data());
    }
    // UPDATE COUNCIL
    static async updateCouncil(council) {
        const { id, ...rest } = council;
        let docRef = null;
        try {
            const collections = await admin.firestore().listCollections();
            const batch = admin.firestore().batch();
            if (rest.role === "lecturer") {
                for (const collection of collections) {
                    const querySnapshot = await collection
                        .where("lecturer.id", "==", id)
                        .get();
                    querySnapshot.forEach((doc) => {
                        const docRef = collection.doc(doc.id);
                        batch.update(docRef, { lecturer: { id, ...rest } });
                    });
                }
            }
            else {
                for (const collection of collections) {
                    const querySnapshot = await collection
                        .where("member.id", "==", id)
                        .get();
                    querySnapshot.forEach((doc) => {
                        const docRef = collection.doc(doc.id);
                        batch.update(docRef, { member: { id, ...rest } });
                    });
                }
            }
            await batch.commit();
            if (docRef) {
                await this.councilDoc.doc(id).update(docRef);
            }
        }
        catch (error) {
            console.error(error);
            throw new Error("Error to update sync council");
        }
    }
    // DELETE COUNCIL
    static async deleteCouncil(id) {
        await this.councilDoc.doc(id).delete();
        const snapshot = await this.councilDoc.where("id", "==", id).get();
        const councils = [];
        snapshot.forEach((doc) => {
            councils.push({ id: doc.id, ...doc.data() });
        });
        return councils;
    }
    static async deleteAllDocuments() {
        const querySnapshot = await this.councilDoc.get();
        const batch = admin.firestore().batch();
        querySnapshot.forEach((doc) => {
            batch.delete(doc.ref);
        });
        await batch.commit();
    }
}
exports.CouncilModel = CouncilModel;
_a = CouncilModel;
CouncilModel.db = admin.firestore();
CouncilModel.councilDoc = _a.db.collection("council-defs");
