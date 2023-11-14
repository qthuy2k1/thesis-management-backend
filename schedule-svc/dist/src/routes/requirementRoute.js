"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const requirementController_1 = require("../controllers/requirementController");
const Router = express_1.default.Router();
Router.route("/")
    .get(requirementController_1.RequirementController.getAllRequirement)
    .post(requirementController_1.RequirementController.createRequirement);
Router.route("/:id")
    .get(requirementController_1.RequirementController.getRequirement)
    .put(requirementController_1.RequirementController.updateRequirement)
    .delete(requirementController_1.RequirementController.deleteRequirement);
Router.route("/class/:id").get(requirementController_1.RequirementController.getAllRequirementClassroom // id (id của student) -> lấy ra những requirement có chứa classroom trong đó có chứa id của lecturer
);
exports.default = Router;
