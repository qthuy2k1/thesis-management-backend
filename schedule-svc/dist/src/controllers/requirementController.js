"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.RequirementController = void 0;
const requirementModel_1 = require("../models/requirementModel");
exports.RequirementController = {
    async createRequirement(req, res, next) {
        const requirement = req.body;
        try {
            const createdRequirement = await requirementModel_1.RequirementModel.createRequirement(requirement);
            if (createdRequirement) {
                res.status(200).json({
                    requirement: createdRequirement,
                    message: "Requirement has been created",
                });
            }
        }
        catch (err) {
            if (err instanceof Error) {
                if (err.message === 'Requirement already exists') {
                    res.status(400).json({ message: err.message });
                }
                else if (err.message === "You don't send more than 2 requirements") {
                    res.status(400).json({ message: err.message });
                }
                else {
                    next(err);
                }
            }
            else {
                next(err);
            }
        }
    },
    async getRequirement(req, res) {
        const id = req.params.id;
        try {
            const requirement = await requirementModel_1.RequirementModel.getRequirement(id);
            if (!requirement) {
                res.status(404);
                return;
            }
            res.status(200).json(requirement);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllRequirement(req, res) {
        try {
            const requirements = await requirementModel_1.RequirementModel.getAllRequirement();
            if (!requirements) {
                res.status(404).json("Requirement is empty");
                return;
            }
            res.status(200).json(requirements);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllRequirementClassroom(req, res) {
        try {
            const id = req.params.id;
            const requirements = await requirementModel_1.RequirementModel.getAllRequirementClassroom(id);
            if (!requirements) {
                res.status(404).json("Requirement is empty");
                return;
            }
            res.status(200).json(requirements);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateRequirement(req, res) {
        const requirement = req.body;
        const id = req.params.id;
        try {
            await requirementModel_1.RequirementModel.updateRequirement({ id, ...requirement });
            res.status(200).json({ id, ...requirement });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteRequirement(req, res) {
        const id = req.params.id;
        try {
            const requirements = await requirementModel_1.RequirementModel.deleteRequirement(id);
            if (!requirements) {
                res.status(404);
                return false;
            }
            return res.status(200).json(requirements);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
