"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SubmitController = void 0;
const submitModel_1 = require("../models/submitModel");
const google_config_1 = require("../config/google-config");
const fs = require("fs");
const path = require("path");
exports.SubmitController = {
    async createSubmit(req, res) {
        const submit = req.body;
        const uploadsPath = path.join(path.resolve(__dirname, ".."), "uploads");
        const submitPath = path.join(uploadsPath, submit.uid);
        try {
            const fileUrls = [];
            const files = fs.readdirSync(submitPath);
            for (const file of files) {
                const filePath = path.join(submitPath, file);
                const fileUrl = await (0, google_config_1.uploadAndGeneratePublicUrl)(file, filePath);
                fileUrls.push(fileUrl);
            }
            await submitModel_1.SubmitModel.createSubmit({
                status: submit.status,
                student: JSON.parse(submit.student),
                attachments: fileUrls,
                exerciseId: submit.exerciseId,
                uid: submit.uid,
            });
            res.status(200).json({ submit, message: "Submit has been created" });
            const removePath = `src/uploads/${submit.uid}`;
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
    async getSubmit(req, res) {
        const { studentId, exerciseId } = req.params;
        try {
            const submit = await submitModel_1.SubmitModel.getSubmit(exerciseId, studentId);
            if (!submit) {
                res.status(404);
                return;
            }
            res.status(200).json(submit);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllSubmit(req, res) {
        try {
            const id = req.params.id;
            const submits = await submitModel_1.SubmitModel.getAllSubmit(id);
            if (!submits) {
                res.status(404).json("Submit is empty");
                return;
            }
            res.status(200).json(submits);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateSubmit(req, res) {
        const submit = req.body;
        const id = req.params.id;
        try {
            await submitModel_1.SubmitModel.updateSubmit({ id, ...submit });
            res.status(200).json({ id, ...submit });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteSubmit(req, res) {
        const id = req.params.id;
        try {
            const submits = await submitModel_1.SubmitModel.deleteSubmit(id);
            if (!submits) {
                res.status(404);
                return false;
            }
            return res.status(200).json(submits);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
