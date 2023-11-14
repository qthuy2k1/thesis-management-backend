"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const councilController_1 = require("../controllers/councilController");
const Router = express_1.default.Router();
Router.route("/")
    .get(councilController_1.CouncilController.getAllCouncil)
    .post(councilController_1.CouncilController.createCouncil)
    .delete(councilController_1.CouncilController.deleteAllCouncil);
Router.route("/:id")
    .get(councilController_1.CouncilController.getCouncil)
    .put(councilController_1.CouncilController.updateCouncil)
    .delete(councilController_1.CouncilController.deleteCouncil);
exports.default = Router;
