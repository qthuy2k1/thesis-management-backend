"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ClassroomController = void 0;
const classroomModel_1 = require("../models/classroomModel");
exports.ClassroomController = {
    async createClassroom(req, res) {
        const classroom = req.body;
        try {
            await classroomModel_1.ClassroomModel.createClassroom(classroom);
            res
                .status(200)
                .json({ classroom, message: "Classroom has been created" });
        }
        catch (err) {
            res.status(400).json({ message: err });
        }
    },
    async getClassroom(req, res) {
        const id = req.params.id;
        try {
            const classroom = await classroomModel_1.ClassroomModel.getClassroom(id);
            if (!classroom) {
                res.status(404);
                return;
            }
            res.status(200).json(classroom);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllClassroom(req, res) {
        try {
            const classrooms = await classroomModel_1.ClassroomModel.getAllClassroom();
            if (!classrooms) {
                res.status(404).json("Classroom is empty");
                return;
            }
            res.status(200).json(classrooms);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateClassroom(req, res) {
        const classroom = req.body;
        const id = req.params.id;
        try {
            await classroomModel_1.ClassroomModel.updateClassroom({ id, ...classroom });
            res.status(200).json({ id, ...classroom });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteClassroom(req, res) {
        const id = req.params.id;
        try {
            const classrooms = await classroomModel_1.ClassroomModel.deleteClassroom(id);
            if (!classrooms) {
                res.status(404);
                return false;
            }
            return res.status(200).json(classrooms);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
