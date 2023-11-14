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
exports.RequirementModel = void 0;
const admin = __importStar(require("firebase-admin"));
class RequirementModel {
    // CREATE REQUIREMENT
    static async createRequirement(requirement) {
        const existingUser = await this.requirementDoc
            .where("member.id", "==", requirement.member.id)
            .get();
        const existingRequirement = await this.requirementDoc
            .where("classroom.id", "==", requirement.classroom.id)
            .get();
        if (!existingUser.empty && !existingRequirement.empty) {
            throw new Error("Requirement already exists");
        }
        else {
            if (existingUser.size >= 2) {
                throw new Error("You don't send more than 2 requirements");
            }
            else {
                const docRef = await this.requirementDoc.add({
                    ...requirement,
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
    }
    // GET ONE REQUIREMENT
    static async getRequirement(id) {
        const docRef = await this.requirementDoc.doc(id).get();
        if (!docRef.exists) {
            return null;
        }
        const data = docRef.data();
        return { ...data };
    }
    // GET ALL REQUIREMENTS
    static async getAllRequirement() {
        const docRef = await this.requirementDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    static async getAllRequirementClassroom(id) {
        const docRef = await this.requirementDoc
            .where("classroom.lecturer.id", "==", id)
            .get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE REQUIREMENT
    static async updateRequirement(auth) {
        const { id, ...docRef } = auth;
        await this.requirementDoc.doc(id).update(docRef);
    }
    // DELETE REQUIREMENT
    static async deleteRequirement(id) {
        await this.requirementDoc.doc(id).delete();
        const snapshot = await this.requirementDoc.where("id", "==", id).get();
        const auths = [];
        snapshot.forEach((doc) => {
            auths.push({ id: doc.id, ...doc.data() });
        });
        return auths;
    }
}
exports.RequirementModel = RequirementModel;
_a = RequirementModel;
RequirementModel.db = admin.firestore();
RequirementModel.requirementDoc = _a.db.collection("requirements");
