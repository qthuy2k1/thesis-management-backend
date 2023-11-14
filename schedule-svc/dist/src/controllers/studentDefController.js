"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.StudentDefController = void 0;
const studentDefModel_1 = require("../models/studentDefModel");
exports.StudentDefController = {
    async createStudentDef(req, res, next) {
        const auth = req.body;
        try {
            const createdStudentDef = await studentDefModel_1.StudentDefModel.createStudentDef(auth);
            if (createdStudentDef) {
                res.status(200).json({
                    auth: createdStudentDef,
                    message: "StudentDef has been created",
                });
            }
        }
        catch (err) {
            if (err) {
                console.error(err);
                res.status(500);
            }
        }
    },
    async getStudentDef(req, res) {
        const id = req.params.id;
        try {
            const auth = await studentDefModel_1.StudentDefModel.getStudentDef(id);
            if (!auth) {
                res.status(404);
                return;
            }
            res.status(200).json(auth);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllStudentDef(req, res) {
        try {
            const auths = await studentDefModel_1.StudentDefModel.getAllStudentDef();
            if (!auths) {
                res.status(404).json("StudentDef is empty");
                return;
            }
            res.status(200).json(auths);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllStudentDefPag(req, res) {
        try {
            const id = req.params.id;
            const limit = Number(req.query.limit);
            const page = Number(req.query.page);
            const auths = await studentDefModel_1.StudentDefModel.getAllStudentDefPag(page, limit);
            if (!auths) {
                res.status(404).json("StudentDef is empty");
                return;
            }
            res.status(200).json(auths);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateStudentDef(req, res) {
        const auth = req.body;
        const id = req.params.id;
        try {
            await studentDefModel_1.StudentDefModel.updateStudentDef({ id, ...auth });
            res.status(200).json({ id, ...auth });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteStudentDef(req, res) {
        const id = req.params.id;
        try {
            const auths = await studentDefModel_1.StudentDefModel.deleteStudentDef(id);
            if (!auths) {
                res.status(404);
                return false;
            }
            return res.status(200).json(auths);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteAllStudentDef(req, res) {
        try {
            await studentDefModel_1.StudentDefModel.deleteAllDocuments();
            return res.status(200).json("Clear all collections in auth");
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
