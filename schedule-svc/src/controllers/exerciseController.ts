import { Request, Response } from "express";
import { ExerciseModel } from "../models/exerciseModel";
import { uploadAndGeneratePublicUrl } from "../config/google-config";
import { IGeneralLinkAttachment } from "../interface/submit";
const fs = require("fs");
const path = require("path");

export const ExerciseController = {
  async createAttachment(req: any): Promise<IGeneralLinkAttachment[] | undefined> {
    const exercise = req.request;
    const uploadsPath = path.join(path.resolve(__dirname, ".."), "uploads");
    const exercisePath = path.join(uploadsPath, exercise.uid);
    try {
      const fileUrls: IGeneralLinkAttachment[] = [];
      const files: File[] = fs.readdirSync(exercisePath);
      for (const file of files) {
        const filePath = path.join(exercisePath, file);
        const fileUrl: IGeneralLinkAttachment =
          await uploadAndGeneratePublicUrl(file, filePath);
        fileUrls.push(fileUrl);
      }
      const removePath = `src/uploads/${exercise.uid}`;
      fs.rm(removePath, { recursive: true }, (error: any) => {
        if (error) {
          console.error("Error removing upload directory:", error);
        }
      });

      return fileUrls
    } catch (err) {
      console.log(err);
      // res.status(500).json({ message: "fail" });
    }
  },
  async getAttachment(req: any): Promise<File[] | IGeneralLinkAttachment[] | undefined> {
    const id = req.request.id;
    try {
      const exercise = await ExerciseModel.getExercise(id);
      if (!exercise) {
        console.log("Exercise not found")
        // res.status(404);
        return;
      }
    //   res.status(200).json(exercise);
      return exercise.attachments
    } catch (err) {
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
