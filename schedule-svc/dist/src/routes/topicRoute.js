"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const topicController_1 = require("../controllers/topicController");
const Router = express_1.default.Router();
Router.route("/")
    .get(topicController_1.TopicController.getAllTopic)
    .post(topicController_1.TopicController.createTopic);
Router.route("/:id")
    .get(topicController_1.TopicController.getTopic)
    .put(topicController_1.TopicController.updateTopic)
    .delete(topicController_1.TopicController.deleteTopic);
exports.default = Router;
