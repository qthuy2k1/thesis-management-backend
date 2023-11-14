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
exports.PostModel = void 0;
const admin = __importStar(require("firebase-admin"));
class PostModel {
    // CREATE POST
    static async createPost(post) {
        const existingPost = await this.postDoc.where("uid", "==", post.uid).get();
        if (!existingPost.empty) {
            throw new Error("Post exists");
        }
        else {
            const docRef = await this.postDoc.add({
                ...post,
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
    // GET ONE POST
    static async getPost(id) {
        const docRef = await this.postDoc.doc(id).get();
        if (!docRef.exists) {
            return null;
        }
        const data = docRef.data();
        return { ...data };
    }
    // GET ALL POST
    static async getAllPost() {
        const docRef = await this.postDoc.orderBy("createdAt", "desc").get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // GET ALL POST IN CLASS
    static async getAllPostInClass(id) {
        const docRef = await this.postDoc.where("classroom.id", "==", id).get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // GET ALL POST IN REPORT STAGE
    static async getAllPostInReportStage(classroomId, categoryId) {
        const docRef = await this.postDoc
            .where("classroom.id", "==", classroomId)
            .where("category.id", "==", categoryId)
            .get();
        return docRef.docs.map((doc) => ({ id: doc.id, ...doc.data() }));
    }
    // UPDATE POST
    static async updatePost(post) {
        const { id, ...docRef } = post;
        await this.postDoc.doc(id).update(docRef);
    }
    // DELETE POST
    static async deletePost(id) {
        await this.postDoc.doc(id).delete();
        const snapshot = await this.postDoc.where("id", "==", id).get();
        const posts = [];
        snapshot.forEach((doc) => {
            posts.push({ id: doc.id, ...doc.data() });
        });
        return posts;
    }
}
exports.PostModel = PostModel;
_a = PostModel;
PostModel.db = admin.firestore();
PostModel.postDoc = _a.db.collection("posts");
