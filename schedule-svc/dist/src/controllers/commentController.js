"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CommentController = void 0;
const commentModel_1 = require("../models/commentModel");
exports.CommentController = {
    async createComment(req, res) {
        const comment = req.body;
        try {
            await commentModel_1.CommentModel.createComment(comment);
            res.status(200).json({ comment, message: "Comment has been created" });
        }
        catch (error) {
            console.log(error);
            res.status(500);
        }
    },
    async getAllComment(req, res) {
        try {
            const id = req.params.id;
            const comments = await commentModel_1.CommentModel.getAllComments(id);
            if (!comments) {
                res.status(404).json("Comment is empty");
                return;
            }
            res.status(200).json(comments);
        }
        catch (error) { }
    },
};
