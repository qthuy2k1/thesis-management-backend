"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const exerciseController_1 = require("../controllers/exerciseController");
const uploadMiddleware_1 = __importDefault(require("../middlewares/uploadMiddleware"));
const Router = express_1.default.Router();
Router.route("/")
    .get(exerciseController_1.ExerciseController.getAllExercise)
    .post(uploadMiddleware_1.default.array("attachment"), exerciseController_1.ExerciseController.createExercise);
Router.route("/:id")
    .get(exerciseController_1.ExerciseController.getExercise)
    .put(exerciseController_1.ExerciseController.updateExercise)
    .delete(exerciseController_1.ExerciseController.deleteExercise);
Router.route("/class/:id").get(exerciseController_1.ExerciseController.getAllExerciseInClass);
Router.route("/stage/:classroomId&:categoryId").get(exerciseController_1.ExerciseController.getAllExerciseInReportStage);
exports.default = Router;
