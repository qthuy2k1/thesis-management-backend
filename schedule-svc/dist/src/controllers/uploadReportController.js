"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UploadReportController = void 0;
const google_config_1 = require("../config/google-config");
const uploadReportModel_1 = require("../models/uploadReportModel");
const fs = require("fs");
const path = require("path");
exports.UploadReportController = {
    async createUploadReport(req, res) {
        const upload = req.body;
        const uploadsPath = path.join(path.resolve(__dirname, ".."), "uploads");
        const uploadPath = path.join(uploadsPath, upload.uid);
        try {
            const fileUrls = [];
            const files = fs.readdirSync(uploadPath);
            for (const file of files) {
                const filePath = path.join(uploadPath, file);
                const fileUrl = await (0, google_config_1.uploadAndGeneratePublicUrl)(file, filePath);
                fileUrls.push(fileUrl);
            }
            await uploadReportModel_1.UploadReportModel.createUploadReport({
                student: JSON.parse(upload.student),
                uid: upload.uid,
                status: upload.status,
                attachments: fileUrls,
            });
            res
                .status(200)
                .json({ upload, message: "UploadReport has been created" });
            const removePath = `src/uploads/${upload.uid}`;
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
    async getUploadReport(req, res) {
        const id = req.params.id;
        try {
            const upload = await uploadReportModel_1.UploadReportModel.getUploadReport(id);
            if (!upload) {
                res.status(404);
                return;
            }
            res.status(200).json(upload);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllUploadReport(req, res) {
        try {
            const uploads = await uploadReportModel_1.UploadReportModel.getAllUploadReport();
            if (!uploads) {
                res.status(404).json("UploadReport is empty");
                return;
            }
            res.status(200).json(uploads);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateUploadReport(req, res) {
        const upload = req.body;
        const id = req.params.id;
        try {
            await uploadReportModel_1.UploadReportModel.updateUploadReport({ id, ...upload });
            res.status(200).json({ id, ...upload });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteUploadReport(req, res) {
        const id = req.params.id;
        try {
            const uploads = await uploadReportModel_1.UploadReportModel.deleteUploadReport(id);
            if (!uploads) {
                res.status(404);
                return false;
            }
            return res.status(200).json(uploads);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
