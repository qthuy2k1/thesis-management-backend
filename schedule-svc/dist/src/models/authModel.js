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
exports.AuthModel = void 0;
const admin = __importStar(require("firebase-admin"));
class AuthModel {
    // CREATE AUTH
    static async createAuth(user) {
        const existingUser = await this.authDoc.doc(user.id).get();
        if (existingUser.exists) {
            throw new Error("Auth exists");
        }
        else {
            const newUser = {
                ...user,
                createdAt: admin.firestore.Timestamp.now(),
            };
            await this.authDoc.doc(user.id).set(newUser);
            return newUser;
        }
    }
    // GET ONE AUTH
    static async getAuth(id) {
        const docRef = await this.authDoc.doc(id).get();
        if (!docRef.exists) {
            return null;
        }
        const data = docRef.data();
        return { ...data };
    }
    // GET ALL AUTHS
    static async getAllAuth() {
        const docRef = await this.authDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    static async getAllLecturer() {
        const docRef = await this.authDoc.where("role", "==", "lecturer").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE AUTH
    static async updateAuth(auth) {
        const { id, ...rest } = auth;
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
            if (!docRef) {
                await this.authDoc.doc(id).update(rest);
            }
        }
        catch (error) {
            console.error(error);
            throw new Error("Error updating sync auth");
        }
    }
    // DELETE AUTH
    static async deleteAuth(id) {
        await this.authDoc.doc(id).delete();
        const snapshot = await this.authDoc.where("id", "==", id).get();
        const auths = [];
        snapshot.forEach((doc) => {
            auths.push({ id: doc.id, ...doc.data() });
        });
        return auths;
    }
    // CHECK STATUS SUBCRIBE
    static async checkStatusSubscribe(id) {
        const memberRef = await this.memberDoc.where("member.id", "==", id).get();
        const requirementRef = await this.requirementDoc
            .where("member.id", "==", id)
            .get();
        if (!memberRef.empty) {
            const members = [];
            memberRef.forEach((doc) => {
                members.push({
                    id: doc.id,
                    ...doc.data(),
                    status: "SUBSCRIBED",
                });
            });
            return members;
        }
        else if (!requirementRef.empty) {
            const requirements = [];
            requirementRef.forEach((doc) => {
                requirements.push({
                    id: doc.id,
                    ...doc.data(),
                    status: "WAITING",
                });
            });
            return requirements;
        }
        return { status: "NO_SUBSCRIBE" };
    }
    static async checkStateClassroom(id) {
        const classroomRef = await this.classroomDoc
            .where("lecturer.id", "==", id)
            .get();
        if (!classroomRef.empty) {
            const classroomData = classroomRef.docs[0].data();
            return {
                ...classroomData,
                id: classroomRef.docs[0].id,
                status: (classroomData.status = "UN_LOCK"),
            };
        }
        else {
            return null;
        }
    }
    // HANDLE AUTHORIZATION CLASSROOM STATE
    static async checkAuthRoleForClassroomState(auth) {
        const requirementRef = await this.requirementDoc
            .where("member.id", "==", auth.id)
            .get();
        if (auth.role === "student") {
            return await this.checkStatusSubscribe(auth.id);
        }
        else {
            return [];
        }
    }
    // HANDLE UNSUBSCRIBE
    static async unsubscribeState(id) {
        const memberRef = await this.memberDoc.where("member.id", "==", id).get();
        if (!memberRef.empty) {
            memberRef.forEach((doc) => {
                doc.ref.delete();
            });
        }
        const requirementRef = await this.requirementDoc
            .where("member.id", "==", id)
            .get();
        if (!requirementRef.empty) {
            requirementRef.forEach((doc) => {
                doc.ref.delete();
            });
        }
    }
}
exports.AuthModel = AuthModel;
_a = AuthModel;
AuthModel.db = admin.firestore();
AuthModel.authDoc = _a.db.collection("auths");
AuthModel.memberDoc = _a.db.collection("members");
AuthModel.requirementDoc = _a.db.collection("requirements");
AuthModel.classroomDoc = _a.db.collection("classrooms");
