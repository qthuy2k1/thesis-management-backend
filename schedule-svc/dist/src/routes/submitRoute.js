"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const submitController_1 = require("../controllers/submitController");
const uploadMiddleware_1 = __importDefault(require("../middlewares/uploadMiddleware"));
const Router = express_1.default.Router();
Router.route("/").post(uploadMiddleware_1.default.array("attachment"), submitController_1.SubmitController.createSubmit);
Router.route("/:id")
    .put(submitController_1.SubmitController.updateSubmit)
    .delete(submitController_1.SubmitController.deleteSubmit);
Router.route("/ex/:id").get(submitController_1.SubmitController.getAllSubmit);
Router.route("/:exerciseId&:studentId").get(submitController_1.SubmitController.getSubmit);
exports.default = Router;
