"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const postController_1 = require("../controllers/postController");
const uploadMiddleware_1 = __importDefault(require("../middlewares/uploadMiddleware"));
const Router = express_1.default.Router();
Router.route("/")
    .get(postController_1.PostController.getAllPost)
    .post(uploadMiddleware_1.default.array("attachment"), postController_1.PostController.createPost);
Router.route("/:id")
    .get(postController_1.PostController.getPost)
    .put(postController_1.PostController.updatePost)
    .delete(postController_1.PostController.deletePost);
Router.route("/class/:id").get(postController_1.PostController.getAllPostInClass);
Router.route("/stage/:classroomId&:categoryId").get(postController_1.PostController.getAllPostInReportStage);
exports.default = Router;
