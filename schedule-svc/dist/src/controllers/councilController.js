"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CouncilController = void 0;
const councilModel_1 = require("../models/councilModel");
exports.CouncilController = {
    async createCouncil(req, res) {
        const council = req.body;
        try {
            await councilModel_1.CouncilModel.createCouncil(council);
            res.status(200).json({ council, message: "Council has been created" });
        }
        catch (err) {
            res.status(400).json({ message: err });
        }
    },
    async getCouncil(req, res) {
        const id = req.params.id;
        try {
            const council = await councilModel_1.CouncilModel.getCouncil(id);
            if (!council) {
                res.status(404);
                return;
            }
            res.status(200).json(council);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllCouncil(req, res) {
        try {
            const councils = await councilModel_1.CouncilModel.getAllCouncil();
            if (!councils) {
                res.status(404).json("Council is empty");
                return;
            }
            res.status(200).json(councils);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateCouncil(req, res) {
        const council = req.body;
        const id = req.params.id;
        try {
            await councilModel_1.CouncilModel.updateCouncil({ id, ...council });
            res.status(200).json({ id, ...council });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteCouncil(req, res) {
        const id = req.params.id;
        try {
            const councils = await councilModel_1.CouncilModel.deleteCouncil(id);
            if (!councils) {
                res.status(404);
                return false;
            }
            return res.status(200).json(councils);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteAllCouncil(req, res) {
        try {
            await councilModel_1.CouncilModel.deleteAllDocuments();
            return res.status(200).json("Clear all collections in council");
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
