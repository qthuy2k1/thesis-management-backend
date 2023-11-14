"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReportStageController = void 0;
const reportStageModel_1 = require("../models/reportStageModel");
exports.ReportStageController = {
    async createReportStage(req, res) {
        const stage = req.body;
        try {
            await reportStageModel_1.ReportStageModel.createReportStage(stage);
            res.status(200).json({ stage, message: "ReportStage has been created" });
        }
        catch (err) {
            res.status(400).json({ message: err });
        }
    },
    async getReportStage(req, res) {
        const id = req.params.id;
        try {
            const stage = await reportStageModel_1.ReportStageModel.getReportStage(id);
            if (!stage) {
                res.status(404);
                return;
            }
            res.status(200).json(stage);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllReportStage(req, res) {
        try {
            const stages = await reportStageModel_1.ReportStageModel.getAllReportStage();
            if (!stages) {
                res.status(404).json("ReportStage is empty");
                return;
            }
            res.status(200).json(stages);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateReportStage(req, res) {
        const stage = req.body;
        const id = req.params.id;
        try {
            await reportStageModel_1.ReportStageModel.updateReportStage({ id, ...stage });
            res.status(200).json({ id, ...stage });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteReportStage(req, res) {
        const id = req.params.id;
        try {
            const stages = await reportStageModel_1.ReportStageModel.deleteReportStage(id);
            if (!stages) {
                res.status(404);
                return false;
            }
            return res.status(200).json(stages);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
