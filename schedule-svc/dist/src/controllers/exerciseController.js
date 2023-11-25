"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ExerciseController = void 0;
const exerciseModel_1 = require("../models/exerciseModel");
const google_config_1 = require("../config/google-config");
const fs = require("fs");
const path = require("path");
exports.ExerciseController = {
    async createAttachment(req) {
        const exercise = req.request;
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
            const removePath = `src/uploads/${exercise.uid}`;
            fs.rm(removePath, { recursive: true }, (error) => {
                if (error) {
                    console.error("Error removing upload directory:", error);
                }
            });
            return fileUrls;
        }
        catch (err) {
            console.log(err);
            // res.status(500).json({ message: "fail" });
        }
    },
    async getAttachment(req) {
        const id = req.request.id;
        try {
            const exercise = await exerciseModel_1.ExerciseModel.getExercise(id);
            if (!exercise) {
                console.log("Exercise not found");
                // res.status(404);
                return;
            }
            //   res.status(200).json(exercise);
            return exercise.attachments;
        }
        catch (err) {
            console.error(err);
            //   res.status(500);
        }
    },
    //   async getAllExercise(req: Request, res: Response) {
    //     try {
    //       const exercises = await ExerciseModel.getAllExercise();
    //       if (!exercises) {
    //         res.status(404).json("Exercise is empty");
    //         return;
    //       }
    //       res.status(200).json(exercises);
    //     } catch (err) {
    //       console.error(err);
    //       res.status(500);
    //     }
    //   },
};
