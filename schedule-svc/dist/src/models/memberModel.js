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
exports.MemberModel = void 0;
const admin = __importStar(require("firebase-admin"));
class MemberModel {
    // CREATE MEMBER
    static async createMember(member) {
        const existingUser = await this.memberDoc
            .where("member.id", "==", member.member.id)
            .get();
        if (!existingUser.empty) {
            throw new Error("Member exists");
        }
        else {
            const batch = admin.firestore().batch();
            const getRequirementUser = await this.requirementDoc
                .where("member.id", "==", member.member.id)
                .get();
            getRequirementUser.forEach((doc) => {
                batch.delete(doc.ref);
            });
            const docRef = this.memberDoc.doc();
            batch.set(docRef, {
                ...member,
                createdAt: admin.firestore.Timestamp.now(),
            });
            await batch.commit();
            const createdStageSnapshot = await docRef.get();
            const createdStage = {
                id: createdStageSnapshot.id,
                ...createdStageSnapshot.data(),
            };
            return createdStage;
        }
    }
    // GET ONE MEMBER
    static async getMember(id) {
        const querySnapshot = await this.memberDoc
            .where("member.id", "==", id)
            .get();
        if (querySnapshot.empty) {
            return null;
        }
        const docRef = querySnapshot.docs[0];
        const data = docRef.data();
        return { ...data, id: docRef.id };
    }
    // GET ALL MEMBERS
    static async getAllMember() {
        const docRef = await this.memberDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // GET ALL MEMBERS IN CLASS
    static async getAllMemberClassroom(id) {
        const docRef = await this.memberDoc
            .where("classroom.id", "==", id)
            // .orderBy("createdAt", "desc")
            .get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE MEMBER
    static async updateMember(auth) {
        const { id, ...docRef } = auth;
        await this.memberDoc.doc(id).update(docRef);
    }
    // DELETE MEMBER
    static async deleteMember(id) {
        await this.memberDoc.doc(id).delete();
        const snapshot = await this.memberDoc.where("id", "==", id).get();
        const auths = [];
        snapshot.forEach((doc) => {
            auths.push({ id: doc.id, ...doc.data() });
        });
        return auths;
    }
}
exports.MemberModel = MemberModel;
_a = MemberModel;
MemberModel.db = admin.firestore();
MemberModel.memberDoc = _a.db.collection("members");
MemberModel.requirementDoc = _a.db.collection("requirements");
