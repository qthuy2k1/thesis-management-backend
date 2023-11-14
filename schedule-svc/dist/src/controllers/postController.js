"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PostController = void 0;
const postModel_1 = require("../models/postModel");
const google_config_1 = require("../config/google-config");
const fs = require("fs");
const path = require("path");
exports.PostController = {
    async createPost(req, res) {
        const post = req.body;
        const uploadsPath = path.join(path.resolve(__dirname, ".."), "uploads");
        const postPath = path.join(uploadsPath, post.uid);
        try {
            const fileUrls = [];
            const files = fs.readdirSync(postPath);
            for (const file of files) {
                const filePath = path.join(postPath, file);
                const fileUrl = await (0, google_config_1.uploadAndGeneratePublicUrl)(file, filePath);
                fileUrls.push(fileUrl);
            }
            await postModel_1.PostModel.createPost({
                title: post.title,
                description: post.description,
                lecturer: JSON.parse(post.lecturer),
                category: JSON.parse(post.category),
                classroom: JSON.parse(post.classroom),
                uid: post.uid,
                type: post.type,
                attachments: fileUrls,
            });
            res.status(200).json({ post, message: "Post has been created" });
            const removePath = `src/uploads/${post.uid}`;
            fs.rm(removePath, { recursive: true }, (error) => {
                if (error) {
                    console.error("Error removing upload directory:", error);
                }
            });
        }
        catch (err) {
            console.log(err);
            res.status(500).json({ message: "fail" });
        }
    },
    async getPost(req, res) {
        const id = req.params.id;
        try {
            const post = await postModel_1.PostModel.getPost(id);
            if (!post) {
                res.status(404);
                return;
            }
            res.status(200).json(post);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllPost(req, res) {
        try {
            const posts = await postModel_1.PostModel.getAllPost();
            if (!posts) {
                res.status(404).json("Post is empty");
                return;
            }
            res.status(200).json(posts);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllPostInClass(req, res) {
        try {
            const id = req.params.id;
            const posts = await postModel_1.PostModel.getAllPostInClass(id);
            if (!posts) {
                res.status(404).json("Post is empty");
                return;
            }
            res.status(200).json(posts);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllPostInReportStage(req, res) {
        try {
            const { classroomId, categoryId } = req.params;
            const posts = await postModel_1.PostModel.getAllPostInReportStage(classroomId, categoryId);
            if (!posts) {
                res.status(404).json("Post is empty");
                return;
            }
            res.status(200).json(posts);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updatePost(req, res) {
        const post = req.body;
        const id = req.params.id;
        try {
            await postModel_1.PostModel.updatePost({ id, ...post });
            res.status(200).json({ id, ...post });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deletePost(req, res) {
        const id = req.params.id;
        try {
            const posts = await postModel_1.PostModel.deletePost(id);
            if (!posts) {
                res.status(404);
                return false;
            }
            return res.status(200).json(posts);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
