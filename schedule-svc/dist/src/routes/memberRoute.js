"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const memberController_1 = require("../controllers/memberController");
const Router = express_1.default.Router();
Router.route("/")
    .get(memberController_1.MemberController.getAllMember)
    .post(memberController_1.MemberController.createMember);
Router.route("/:id")
    .get(memberController_1.MemberController.getMember) // id (id của student) -> lấy ra memberObject có chứa memberField mà trong memberField có chứa id
    .put(memberController_1.MemberController.updateMember)
    .delete(memberController_1.MemberController.deleteMember);
Router.route("/class/:id").get(memberController_1.MemberController.getAllMemberClassroom); // id (id của classroom) -> lấy ra những member có chứa classroom id
exports.default = Router;
