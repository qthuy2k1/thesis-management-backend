"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const commentController_1 = require("../controllers/commentController");
const Router = express_1.default.Router();
Router.route("/").post(commentController_1.CommentController.createComment);
Router.route("/:id").get(commentController_1.CommentController.getAllComment);
exports.default = Router;
