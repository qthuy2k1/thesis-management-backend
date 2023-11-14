"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ExerciseController = void 0;
const exerciseModel_1 = require("../models/exerciseModel");
const google_config_1 = require("../config/google-config");
const fs = require("fs");
const path = require("path");
exports.ExerciseController = {
    async createExercise(req, res) {
        const exercise = req.body;
        const uploadsPath = path.join(path.resolve(__dirname, ".."), "uploads");
        const exercisePath = path.join(uploadsPath, exercise.uid);
        try {
            const fileUrls = [];
            const files = fs.readdirSync(exercisePath);
            for (const file of files) {
                const filePath = path.join(exercisePath, file);
                const fileUrl = await (0, google_config_1.uploadAndGeneratePublicUrl)(file, filePath);
                fileUrls.push(fileUrl);
            }
            await exerciseModel_1.ExerciseModel.createExercise({
                title: exercise.title,
                description: exercise.description,
                lecturer: JSON.parse(exercise.lecturer),
                category: JSON.parse(exercise.category),
                classroom: JSON.parse(exercise.classroom),
                uid: exercise.uid,
                type: exercise.type,
                attachments: fileUrls,
                deadline: exercise.deadline,
            });
            res.status(200).json({ exercise, message: "Exercise has been created" });
            const removePath = `src/uploads/${exercise.uid}`;
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
    async getExercise(req, res) {
        const id = req.params.id;
        try {
            const exercise = await exerciseModel_1.ExerciseModel.getExercise(id);
            if (!exercise) {
                res.status(404);
                return;
            }
            res.status(200).json(exercise);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllExercise(req, res) {
        try {
            const exercises = await exerciseModel_1.ExerciseModel.getAllExercise();
            if (!exercises) {
                res.status(404).json("Exercise is empty");
                return;
            }
            res.status(200).json(exercises);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllExerciseInClass(req, res) {
        try {
            const id = req.params.id;
            const exercises = await exerciseModel_1.ExerciseModel.getAllExerciseInClass(id);
            if (!exercises) {
                res.status(404).json("Exercise is empty");
                return;
            }
            res.status(200).json(exercises);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllExerciseInReportStage(req, res) {
        try {
            const { classroomId, categoryId } = req.params;
            const exercises = await exerciseModel_1.ExerciseModel.getAllExerciseInReportStage(classroomId, categoryId);
            if (!exercises) {
                res.status(404).json("Exercise is empty");
                return;
            }
            res.status(200).json(exercises);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateExercise(req, res) {
        const exercise = req.body;
        const id = req.params.id;
        try {
            await exerciseModel_1.ExerciseModel.updateExercise({ id, ...exercise });
            res.status(200).json({ id, ...exercise });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteExercise(req, res) {
        const id = req.params.id;
        try {
            const exercises = await exerciseModel_1.ExerciseModel.deleteExercise(id);
            if (!exercises) {
                res.status(404);
                return false;
            }
            return res.status(200).json(exercises);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
